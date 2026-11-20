package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/limitless-hoops/api/database"
	"github.com/limitless-hoops/api/models"
)

// DisconnectProvider handles the backend portion of disconnecting a social provider.
// It converts any emails associated with the provider into manual OTP emails.
func DisconnectProvider(c *fiber.Ctx) error {

	userID, err := resolveUserID(c)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	provider := c.Params("provider")
	if provider == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Provider is required"})
	}

	// H11: Verify they have at least one verified login method before disconnecting
	var verifiedEmails int64
	var verifiedPhones int64
	database.DB.Model(&models.UserEmail{}).Where("user_id = ? AND is_verified = ?", userID, true).Count(&verifiedEmails)
	database.DB.Model(&models.UserPhone{}).Where("user_id = ? AND is_verified = ?", userID, true).Count(&verifiedPhones)
	
	if verifiedEmails == 0 && verifiedPhones == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "You must have at least one verified email or phone number to disconnect this provider.",
		})
	}

	// Instead of deleting the email, we convert it to a manual email so they can still use it for OTP.
	if err := database.DB.Model(&models.UserEmail{}).
		Where("user_id = ? AND source_provider = ?", userID, provider).
		Update("source_provider", "manual").Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update provider email"})
	}

	return c.JSON(fiber.Map{
		"message": "Provider disconnected successfully from Limitless Hoops. Proceed with client-side provider removal.",
	})
}
