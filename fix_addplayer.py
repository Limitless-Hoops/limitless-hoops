import re

with open('apps/api/handlers/account.go', 'r') as f:
    content = f.read()

old_add_player = """	playerID := uuid.New().String()
	player := models.Player{
		ID:             playerID,
		HouseholdID:    householdID,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		DOB:            req.DOB,
		GraduationYear: req.GraduationYear,
		Email:          req.Email,
		Phone:          req.Phone,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := database.DB.Create(&player).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create player"})
	}

	return c.Status(fiber.StatusCreated).JSON(player)"""

new_add_player = """	playerID := uuid.New().String()
	player := models.Player{
		ID:             playerID,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		DOB:            req.DOB,
		GraduationYear: req.GraduationYear,
		Email:          req.Email,
		Phone:          req.Phone,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err := database.DB.Transaction(func(tx *gorm.DB) error {
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
	}

	return c.Status(fiber.StatusCreated).JSON(player)"""

content = content.replace(old_add_player, new_add_player)

with open('apps/api/handlers/account.go', 'w') as f:
    f.write(content)

