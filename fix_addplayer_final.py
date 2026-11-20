import re

with open('apps/api/handlers/account.go', 'r') as f:
    content = f.read()

old_block = """	if err := database.DB.Create(&player).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to add player",
			"details": err.Error(),
		})
	}"""

new_block = """	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&player).Error; err != nil {
			return err
		}
		
		hp := models.HouseholdPlayer{
			HouseholdID:  householdID,
			PlayerID:     player.ID,
			Relationship: req.Relationship,
		}
		if err := tx.Create(&hp).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create player",
			"details": err.Error(),
		})
	}"""

content = content.replace(old_block, new_block)

with open('apps/api/handlers/account.go', 'w') as f:
    f.write(content)

