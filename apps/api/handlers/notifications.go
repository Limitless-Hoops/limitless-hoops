package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/limitless-hoops/api/database"
	"github.com/limitless-hoops/api/models"
)

// NotificationPreferenceEnriched adds contact label and value for display in the UI.
type NotificationPreferenceEnriched struct {
	models.NotificationPreference
	ContactLabel string `json:"contact_label"`
	ContactValue string `json:"contact_value"` // email address or phone number
	IsPrimary    bool   `json:"is_primary"`
}

func GetMyNotifications(c *fiber.Ctx) error {

	userID, err := resolveUserID(c)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	var prefs []models.NotificationPreference
	if err := database.DB.Where("user_id = ?", userID).Find(&prefs).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch preferences"})
	}

	// Build enriched response by joining contact details
	enriched := make([]NotificationPreferenceEnriched, 0, len(prefs))
	for _, pref := range prefs {
		row := NotificationPreferenceEnriched{NotificationPreference: pref}
		switch pref.TargetType {
		case "email":
			var ue models.UserEmail
			if err := database.DB.First(&ue, "id = ?", pref.TargetID).Error; err == nil {
				row.ContactLabel = ue.Label
				row.ContactValue = ue.Email
				row.IsPrimary = ue.IsPrimary
			}
		case "phone":
			var up models.UserPhone
			if err := database.DB.First(&up, "id = ?", pref.TargetID).Error; err == nil {
				row.ContactLabel = up.Label
				row.ContactValue = up.Phone
				row.IsPrimary = up.IsPrimary
			}
		}
		enriched = append(enriched, row)
	}

	return c.JSON(enriched)
}

type UpdateNotificationRequest struct {
	Enabled bool `json:"enabled"`
}

func UpdateNotificationPreference(c *fiber.Ctx) error {

	userID, err := resolveUserID(c)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	id := c.Params("id")
	var pref models.NotificationPreference
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&pref).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Preference not found"})
	}

	var req UpdateNotificationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Business rule: security notifications cannot be disabled on the primary contact
	if pref.Category == "security" && !req.Enabled {
		// Check if the target contact is primary
		isPrimary := false
		switch pref.TargetType {
		case "email":
			var ue models.UserEmail
			if err := database.DB.First(&ue, "id = ?", pref.TargetID).Error; err == nil {
				isPrimary = ue.IsPrimary
			}
		case "phone":
			var up models.UserPhone
			if err := database.DB.First(&up, "id = ?", pref.TargetID).Error; err == nil {
				isPrimary = up.IsPrimary
			}
		}
		if isPrimary {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Security notifications cannot be disabled on your primary contact",
			})
		}
	}

	if err := database.DB.Model(&pref).Update("enabled", req.Enabled).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update preference"})
	}

	return c.JSON(fiber.Map{"message": "Preference updated", "enabled": req.Enabled})
}
