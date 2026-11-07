import re

with open('apps/api/handlers/contact.go', 'r') as f:
    content = f.read()

new_handlers = """
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
	digits := regexp.MustCompile(`\\D`).ReplaceAllString(normalizedPhone, "")
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
"""

if 'UpdateEmail' not in content:
    content += new_handlers

    # Add regexp import if needed
    if '"regexp"' not in content:
        content = content.replace('import (', 'import (\n\t"regexp"')

with open('apps/api/handlers/contact.go', 'w') as f:
    f.write(content)

