import re

with open('apps/api/handlers/account.go', 'r') as f:
    content = f.read()

# Remove HouseholdID: householdID, from AddPlayer
content = re.sub(r'ID:\s*playerID,\s*HouseholdID:\s*householdID,', r'ID:             playerID,', content)

# Change database.DB.Create(&player) to Transaction for AddPlayer
# Let's just find the function and rewrite it carefully.
match = re.search(r'func AddPlayer\(c \*fiber.Ctx\) error \{.*?\n\}', content, re.DOTALL)
if match:
    func_body = match.group(0)
    # If the transaction is not there, replace it
    if 'tx.Create(&hp)' not in func_body:
        new_func_body = func_body.replace("""	if err := database.DB.Create(&player).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create player"})
	}""", """	err := database.DB.Transaction(func(tx *gorm.DB) error {
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create player"})
	}""")
        content = content.replace(func_body, new_func_body)

with open('apps/api/handlers/account.go', 'w') as f:
    f.write(content)

