import re

with open('apps/api/handlers/account.go', 'r') as f:
    content = f.read()

# Delete AddGuardian and RemoveGuardian completely
old_guardian_endpoints = content[content.find('type AddGuardianRequest struct {'):]
# Actually wait, let's find the exact block to replace.

def replace_between(text, start_str, end_str, replacement):
    start = text.find(start_str)
    if start == -1: return text
    # we don't have end_str because it goes to EOF, or if we do, we find it.
    return text[:start] + replacement

new_contact_endpoints = """type AddPlayerContactRequest struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Relationship string `json:"relationship"`
	ReceiveSMS   bool   `json:"receive_sms"`
	ReceiveEmail bool   `json:"receive_email"`
}

func AddPlayerContact(c *fiber.Ctx) error {
	playerID := c.Params("id")
	
	// Ensure player exists
	var player models.Player
	if err := database.DB.First(&player, "id = ?", playerID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Player not found"})
	}

	var req AddPlayerContactRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	contact := models.PlayerContact{
		ID:           uuid.New().String(),
		PlayerID:     playerID,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Phone:        req.Phone,
		Email:        req.Email,
		Relationship: req.Relationship,
		ReceiveSMS:   req.ReceiveSMS,
		ReceiveEmail: req.ReceiveEmail,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := database.DB.Create(&contact).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to add contact"})
	}

	return c.Status(fiber.StatusCreated).JSON(contact)
}

func RemovePlayerContact(c *fiber.Ctx) error {
	contactID := c.Params("contactId")
	if err := database.DB.Delete(&models.PlayerContact{}, "id = ?", contactID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to remove contact"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Contact removed"})
}
"""

content = replace_between(content, 'type AddGuardianRequest struct {', None, new_contact_endpoints)

with open('apps/api/handlers/account.go', 'w') as f:
    f.write(content)

