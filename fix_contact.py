import re

with open('apps/api/handlers/contact.go', 'r') as f:
    content = f.read()

# Block 1
old1 = """		// Email already exists but unverified, just trigger resend
		otp, err := createOTP(existing.ID, "email")
		if err == nil {
			if sendErr := sendEmailOTP(existing.Email, otp.Code); sendErr != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send email. Check API logs or Resend configuration."})
			}
		}"""
new1 = """		// Email already exists but unverified, just trigger resend
		otp, err := createOTP(existing.ID, "email")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate OTP"})
		}
		if sendErr := sendEmailOTP(existing.Email, otp.Code); sendErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send email. Check API logs or Resend configuration."})
		}"""
content = content.replace(old1, new1)

# Block 2
old2 = """	// Trigger email verification OTP via Resend
	otp, err := createOTP(newEmail.ID, "email")
	if err == nil {
		if sendErr := sendEmailOTP(newEmail.Email, otp.Code); sendErr != nil {
			// Rollback the creation if we can't deliver the OTP
			database.DB.Delete(&newEmail)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send email. Check API logs or Resend configuration."})
		}
	}"""
new2 = """	// Trigger email verification OTP via Resend
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
	}"""
content = content.replace(old2, new2)

# Block 3
old3 = """		// Phone already exists but unverified, just trigger resend
		otp, err := createOTP(existing.ID, "phone")
		if err == nil {
			if sendErr := sendPhoneOTP(existing.Phone, otp.Code); sendErr != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send SMS."})
			}
		}"""
new3 = """		// Phone already exists but unverified, just trigger resend
		otp, err := createOTP(existing.ID, "phone")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate OTP"})
		}
		if sendErr := sendPhoneOTP(existing.Phone, otp.Code); sendErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send SMS."})
		}"""
content = content.replace(old3, new3)

# Block 4
old4 = """	// Trigger SMS verification via Twilio
	otp, err := createOTP(newPhone.ID, "phone")
	if err == nil {
		if sendErr := sendPhoneOTP(newPhone.Phone, otp.Code); sendErr != nil {
			// Rollback the creation if we can't deliver the OTP
			database.DB.Delete(&newPhone)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send SMS."})
		}
	}"""
new4 = """	// Trigger SMS verification via Twilio
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
	}"""
content = content.replace(old4, new4)

# We need to make sure 'log' is imported. Let's check imports.
if 'import (' in content and '"log"' not in content:
    content = content.replace('import (', 'import (\n\t"log"', 1)

with open('apps/api/handlers/contact.go', 'w') as f:
    f.write(content)

