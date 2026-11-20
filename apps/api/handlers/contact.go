package handlers

import (
	"regexp"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/limitless-hoops/api/database"
	"github.com/limitless-hoops/api/models"
	"gorm.io/gorm"
)

// resolveUserID looks up the authenticated user's ID directly from the Supabase session JWT.
func resolveUserID(c *fiber.Ctx) (string, error) {
	sub, ok := c.Locals("user_id").(string)
	if !ok || sub == "" {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Missing user ID in token")
	}
	return sub, nil
}

// ─── Emails ─────────────────────────────────────────────────────────────────

func GetMyEmails(c *fiber.Ctx) error {
	userID, err := resolveUserID(c)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	var emails []models.UserEmail
	if err := database.DB.Where("user_id = ?", userID).Find(&emails).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch emails"})
	}

	return c.JSON(emails)
}

type AddEmailRequest struct {
	Email string `json:"email"`
	Label string `json:"label"`
}

func AddEmail(c *fiber.Ctx) error {

	userID, err := resolveUserID(c)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	var req AddEmailRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	if req.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "email is required"})
	}

	var count int64
	database.DB.Model(&models.UserEmail{}).Where("user_id = ?", userID).Count(&count)
	if count >= 2 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Maximum of 2 emails allowed per account"})
	}

	var existing models.UserEmail
	if err := database.DB.Where("email = ? AND user_id = ?", req.Email, userID).First(&existing).Error; err == nil {
		if existing.IsVerified {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "This email is already added and verified on your account."})
		}
		// Email already exists but unverified, just trigger resend
		otp, err := createOTP(existing.ID, "email")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate OTP"})
		}
		if sendErr := sendEmailOTP(existing.Email, otp.Code); sendErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send email. Check API logs or Resend configuration."})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Verification email resent.",
			"email":   existing,
		})
	}

	newEmail := models.UserEmail{
		ID:             uuid.New().String(),
		UserID:         userID,
		Email:          req.Email,
		IsPrimary:      false,
		IsVerified:     false, // Verification flow handled separately
		SourceProvider: "manual",
		Label:          req.Label,
		CreatedAt:      time.Now(),
	}

	if err := database.DB.Create(&newEmail).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to add email",
			"details": err.Error(),
		})
	}

	// Trigger email verification OTP via Resend
	otp, err := createOTP(newEmail.ID, "email")
	if err != nil {
		database.DB.Delete(&newEmail)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate OTP"})
	}
	if sendErr := sendEmailOTP(newEmail.Email, otp.Code); sendErr != nil {
		// Rollback the creation if we can't deliver the OTP
		if delErr := database.DB.Delete(&newEmail).Error; delErr != nil {
			log.Printf("Failed to rollback unverified email %s: %v", newEmail.ID, delErr)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send email. Check API logs or Resend configuration."})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Email added. Verification required before it can be used for login.",
		"email":   newEmail,
	})
}

func DeleteEmail(c *fiber.Ctx) error {

	userID, err := resolveUserID(c)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	id := c.Params("id")
	var target models.UserEmail
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&target).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Email record not found"})
	}

	if target.IsPrimary {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot delete the primary email"})
	}
	if target.SourceProvider != "manual" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Provider-sourced emails can only be removed by disconnecting the provider",
		})
	}

	// Guard: must not be the only email
	var count int64
	database.DB.Model(&models.UserEmail{}).Where("user_id = ?", userID).Count(&count)
	if count <= 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot delete the only email on the account"})
	}

	if err := database.DB.Delete(&target).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete email"})
	}

	return c.JSON(fiber.Map{"message": "Email deleted"})
}

func SetPrimaryEmail(c *fiber.Ctx) error {

	userID, err := resolveUserID(c)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	id := c.Params("id")

	err = database.DB.Transaction(func(tx *gorm.DB) error {
		// Verify target belongs to user
		var target models.UserEmail
		if err := tx.Where("id = ? AND user_id = ?", id, userID).First(&target).Error; err != nil {
			return fiber.NewError(fiber.StatusNotFound, "Email record not found")
		}

		// Demote all other emails for this user
		if err := tx.Model(&models.UserEmail{}).
			Where("user_id = ? AND id != ?", userID, id).
			Update("is_primary", false).Error; err != nil {
			return err
		}

		// Promote target
		if err := tx.Model(&target).Update("is_primary", true).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		if fe, ok := err.(*fiber.Error); ok {
			return c.Status(fe.Code).JSON(fiber.Map{"error": fe.Message})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update primary email"})
	}

	return c.JSON(fiber.Map{"message": "Primary email updated"})
}

// ─── Phones ─────────────────────────────────────────────────────────────────

func GetMyPhones(c *fiber.Ctx) error {

	userID, err := resolveUserID(c)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	var phones []models.UserPhone
	if err := database.DB.Where("user_id = ?", userID).Find(&phones).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch phones"})
	}

	return c.JSON(phones)
}

type AddPhoneRequest struct {
	Phone string `json:"phone"`
	Label string `json:"label"`
}

func AddPhone(c *fiber.Ctx) error {

	userID, err := resolveUserID(c)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	var req AddPhoneRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	if req.Phone == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "phone is required"})
	}

	var count int64
	database.DB.Model(&models.UserPhone{}).Where("user_id = ?", userID).Count(&count)
	if count >= 2 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Maximum of 2 phone numbers allowed per account"})
	}

	var existing models.UserPhone
	if err := database.DB.Where("phone = ? AND user_id = ?", req.Phone, userID).First(&existing).Error; err == nil {
		if existing.IsVerified {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "This phone number is already added and verified on your account."})
		}
		// Phone already exists but unverified, just trigger resend
		otp, err := createOTP(existing.ID, "phone")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate OTP"})
		}
		if sendErr := sendPhoneOTP(existing.Phone, otp.Code); sendErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send SMS."})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Verification code resent.",
			"phone":   existing,
		})
	}

	newPhone := models.UserPhone{
		ID:         uuid.New().String(),
		UserID:     userID,
		Phone:      req.Phone,
		IsPrimary:  false,
		IsVerified: false, // Verification flow handled separately
		Label:      req.Label,
		CreatedAt:  time.Now(),
	}

	if err := database.DB.Create(&newPhone).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to add phone",
			"details": err.Error(),
		})
	}

	// Trigger SMS verification via Twilio
	otp, err := createOTP(newPhone.ID, "phone")
	if err != nil {
		database.DB.Delete(&newPhone)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate OTP"})
	}
	if sendErr := sendPhoneOTP(newPhone.Phone, otp.Code); sendErr != nil {
		// Rollback the creation if we can't deliver the OTP
		if delErr := database.DB.Delete(&newPhone).Error; delErr != nil {
			log.Printf("Failed to rollback unverified phone %s: %v", newPhone.ID, delErr)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send SMS."})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Phone added. Verification required before it can be used for login.",
		"phone":   newPhone,
	})
}

func DeletePhone(c *fiber.Ctx) error {

	userID, err := resolveUserID(c)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	id := c.Params("id")
	var target models.UserPhone
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&target).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Phone record not found"})
	}

	if target.IsPrimary {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot delete the primary phone"})
	}

	// Guard: must not be the only phone
	var count int64
	database.DB.Model(&models.UserPhone{}).Where("user_id = ?", userID).Count(&count)
	if count <= 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot delete the only phone on the account"})
	}

	if err := database.DB.Delete(&target).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete phone"})
	}

	return c.JSON(fiber.Map{"message": "Phone deleted"})
}

func SetPrimaryPhone(c *fiber.Ctx) error {

	userID, err := resolveUserID(c)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	id := c.Params("id")

	err = database.DB.Transaction(func(tx *gorm.DB) error {
		// Verify target belongs to user
		var target models.UserPhone
		if err := tx.Where("id = ? AND user_id = ?", id, userID).First(&target).Error; err != nil {
			return fiber.NewError(fiber.StatusNotFound, "Phone record not found")
		}

		// Demote all other phones for this user
		if err := tx.Model(&models.UserPhone{}).
			Where("user_id = ? AND id != ?", userID, id).
			Update("is_primary", false).Error; err != nil {
			return err
		}

		// Promote target
		if err := tx.Model(&target).Update("is_primary", true).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		if fe, ok := err.(*fiber.Error); ok {
			return c.Status(fe.Code).JSON(fiber.Map{"error": fe.Message})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update primary phone"})
	}

	return c.JSON(fiber.Map{"message": "Primary phone updated"})
}

type VerifyCodeRequest struct {
	Code string `json:"code"`
}

func VerifyEmail(c *fiber.Ctx) error {

	userID, err := resolveUserID(c)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	id := c.Params("id")
	var req VerifyCodeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if !validateOTP(id, "email", req.Code) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid or expired verification code."})
	}

	if err := database.DB.Model(&models.UserEmail{}).Where("id = ? AND user_id = ?", id, userID).Update("is_verified", true).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to verify email"})
	}
	return c.JSON(fiber.Map{"message": "Email verified successfully"})
}

func ResendEmailVerification(c *fiber.Ctx) error {

	userID, err := resolveUserID(c)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	id := c.Params("id")
	var target models.UserEmail
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&target).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Email not found"})
	}

	// Trigger Resend OTP
	otp, err := createOTP(target.ID, "email")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate OTP"})
	}
	if sendErr := sendEmailOTP(target.Email, otp.Code); sendErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send email. Check API logs or Resend configuration."})
	}

	return c.JSON(fiber.Map{"message": "Verification code resent"})
}

func VerifyPhone(c *fiber.Ctx) error {

	userID, err := resolveUserID(c)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	id := c.Params("id")
	var req VerifyCodeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if !validateOTP(id, "phone", req.Code) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid or expired verification code."})
	}

	if err := database.DB.Model(&models.UserPhone{}).Where("id = ? AND user_id = ?", id, userID).Update("is_verified", true).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to verify phone"})
	}
	return c.JSON(fiber.Map{"message": "Phone verified successfully"})
}

func ResendPhoneVerification(c *fiber.Ctx) error {

	userID, err := resolveUserID(c)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	id := c.Params("id")
	var target models.UserPhone
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&target).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Phone not found"})
	}

	// Trigger Twilio SMS OTP
	otp, err := createOTP(target.ID, "phone")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate OTP"})
	}
	if sendErr := sendPhoneOTP(target.Phone, otp.Code); sendErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send SMS."})
	}

	return c.JSON(fiber.Map{"message": "Verification code resent"})
}

// --- UPDATE HANDLERS ---

type UpdateContactRequest struct {
	Value string `json:"value"`
}

func UpdateEmail(c *fiber.Ctx) error {
	id := c.Params("id")
	sub, _ := c.Locals("user_id").(string)

	var req UpdateContactRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	var email models.UserEmail
	if err := database.DB.First(&email, "id = ? AND user_id = ?", id, sub).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Email not found"})
	}

	if email.Email != req.Value {
		// Verify if email is already taken
		var existing models.UserEmail
		if err := database.DB.First(&existing, "email = ?", req.Value).Error; err == nil && existing.ID != email.ID {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Email is already in use"})
		}

		email.Email = req.Value
		email.IsVerified = false
		if err := database.DB.Save(&email).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update email"})
		}
	}

	return c.JSON(email)
}

func UpdatePhone(c *fiber.Ctx) error {
	id := c.Params("id")
	sub, _ := c.Locals("user_id").(string)

	var req UpdateContactRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	var phone models.UserPhone
	if err := database.DB.First(&phone, "id = ? AND user_id = ?", id, sub).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Phone not found"})
	}

	// Normalize phone for comparison and storage
	normalizedPhone := req.Value
	digits := regexp.MustCompile(`\D`).ReplaceAllString(normalizedPhone, "")
	if len(digits) == 10 {
		normalizedPhone = "+1" + digits
	} else if len(digits) == 11 && digits[0] == '1' {
		normalizedPhone = "+" + digits
	}

	if phone.Phone != normalizedPhone {
		var existing models.UserPhone
		if err := database.DB.First(&existing, "phone = ?", normalizedPhone).Error; err == nil && existing.ID != phone.ID {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Phone is already in use"})
		}

		phone.Phone = normalizedPhone
		phone.IsVerified = false
		if err := database.DB.Save(&phone).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update phone"})
		}
	}

	return c.JSON(phone)
}
