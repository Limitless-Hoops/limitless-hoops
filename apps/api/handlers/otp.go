package handlers

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/resend/resend-go/v2"
	"github.com/limitless-hoops/api/database"
	"github.com/limitless-hoops/api/models"
)

// Generate a 6-digit code securely
func generateOTPCode() string {
	b := make([]byte, 3)
	rand.Read(b)
	code := fmt.Sprintf("%06d", int(b[0])<<16|int(b[1])<<8|int(b[2]))
	return code[len(code)-6:] // ensure exactly 6 digits
}

// Generate and save an OTP to the database for a target email/phone ID
func createOTP(targetID string, targetType string) (*models.ContactVerification, error) {
	// First, invalidate any existing OTPs for this target
	database.DB.Where("target_id = ? AND target_type = ?", targetID, targetType).Delete(&models.ContactVerification{})

	code := generateOTPCode()
	ver := models.ContactVerification{
		ID:         uuid.New().String(),
		TargetID:   targetID,
		TargetType: targetType,
		Code:       code,
		ExpiresAt:  time.Now().Add(10 * time.Minute), // 10 min expiration
	}

	if err := database.DB.Create(&ver).Error; err != nil {
		return nil, err
	}
	return &ver, nil
}

// Send OTP via Resend
func sendEmailOTP(email string, code string) error {
	apiKey := os.Getenv("RESEND_API_KEY")
	if apiKey == "" {
		log.Printf("⚠️  RESEND_API_KEY not set. Mocking email to %s with code %s\n", email, code)
		return nil
	}

	client := resend.NewClient(apiKey)

	fromEmail := os.Getenv("RESEND_FROM_EMAIL")
	if fromEmail == "" {
		// Resend's default testing domain. Only allows sending to your own signup email!
		fromEmail = "onboarding@resend.dev"
	}

	params := &resend.SendEmailRequest{
		From:    fromEmail,
		To:      []string{email},
		Subject: "Your Verification Code",
		Html:    fmt.Sprintf("<h2>Verify your email</h2><p>Your 6-digit verification code is: <strong>%s</strong></p><p>This code will expire in 10 minutes.</p>", code),
	}

	_, err := client.Emails.Send(params)
	if err != nil {
		log.Printf("Failed to send email to %s: %v", email, err)
		return err
	}
	return nil
}

// Send OTP via Twilio
func sendPhoneOTP(phone string, code string) error {
	// TODO: implement Twilio SDK here when credentials are added
	log.Printf("⚠️  TWILIO not configured. Mocking SMS to %s with code %s\n", phone, code)
	return nil
}

// Validate an OTP code
func validateOTP(targetID string, targetType string, code string) bool {
	var ver models.ContactVerification
	err := database.DB.Where("target_id = ? AND target_type = ? AND code = ? AND expires_at > ?", targetID, targetType, code, time.Now()).First(&ver).Error
	if err != nil {
		return false
	}

	// Delete it after successful validation to prevent reuse
	database.DB.Delete(&ver)
	return true
}
