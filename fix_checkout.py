import re

with open('apps/api/handlers/checkout.go', 'r') as f:
    content = f.read()

# Replace metadata block
old_meta = """		Metadata: map[string]string{
			"player_id":    player.ID,
			"household_id": player.HouseholdID,
			"plan_type":    req.PlanType,
		},"""

new_meta = """		Metadata: map[string]string{
			"player_id":    player.ID,
			"household_id": household.ID,
			"plan_type":    req.PlanType,
		},"""

content = content.replace(old_meta, new_meta)

# Get household ID from user
old_validate = """	// 1. Validate Player exists
	var player models.Player
	if err := database.DB.First(&player, "id = ?", req.PlayerID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Player not found"})
	}"""

new_validate = """	// 1. Validate Player exists
	var player models.Player
	if err := database.DB.First(&player, "id = ?", req.PlayerID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Player not found"})
	}
	
	// Get household ID of the authenticated user
	sub, _ := c.Locals("user_id").(string)
	var household models.Household
	database.DB.First(&household, "owner_id = ?", sub)"""

content = content.replace(old_validate, new_validate)

with open('apps/api/handlers/checkout.go', 'w') as f:
    f.write(content)

