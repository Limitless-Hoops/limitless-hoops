import re

with open('apps/api/handlers/account.go', 'r') as f:
    content = f.read()

# 1. Update CreateHousehold logic
# Find existingHu check and replace with Owner check
old_hu_check = """			var existingHu models.HouseholdUser
			if tx.Where("user_id = ?", user.ID).First(&existingHu).Error == nil {
				// Already has a household — nothing to do
				return nil
			}"""
new_hu_check = """			var existingHousehold models.Household
			if tx.Where("owner_id = ?", user.ID).First(&existingHousehold).Error == nil {
				// Already has a household — nothing to do
				return nil
			}"""
content = content.replace(old_hu_check, new_hu_check)

# Update Household creation inside CreateHousehold
old_household_creation = """		// 1. Create Household
		householdID := uuid.New().String()
		household := models.Household{
			ID:        householdID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := tx.Create(&household).Error; err != nil {
			return err
		}

		// 2. Link User to Household as Primary Guardian
		householdUser := models.HouseholdUser{
			HouseholdID: householdID,
			UserID:      user.ID,
			Role:        "Primary",
		}
		if err := tx.Create(&householdUser).Error; err != nil {
			return err
		}"""
new_household_creation = """		// 1. Create Household linked directly to User
		householdID := uuid.New().String()
		household := models.Household{
			ID:        householdID,
			OwnerID:   user.ID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := tx.Create(&household).Error; err != nil {
			return err
		}"""
content = content.replace(old_household_creation, new_household_creation)

# Update Player creation inside CreateHousehold
old_player_creation = """			player := models.Player{
				ID:             playerID,
				HouseholdID:    householdID,
				FirstName:      pReq.FirstName,
				LastName:       pReq.LastName,
				GraduationYear: pReq.GraduationYear,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			}
			if err := tx.Create(&player).Error; err != nil {
				return err
			}"""
new_player_creation = """			player := models.Player{
				ID:             playerID,
				FirstName:      pReq.FirstName,
				LastName:       pReq.LastName,
				GraduationYear: pReq.GraduationYear,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			}
			if err := tx.Create(&player).Error; err != nil {
				return err
			}
			
			// Link player to household
			hp := models.HouseholdPlayer{
				HouseholdID:  householdID,
				PlayerID:     player.ID,
				Relationship: "Primary",
			}
			if err := tx.Create(&hp).Error; err != nil {
				return err
			}"""
content = content.replace(old_player_creation, new_player_creation)

# 2. Update GetMyHousehold
old_get_my_household = """func GetMyHousehold(c *fiber.Ctx) error {
	sub, ok := c.Locals("user_id").(string)
	if !ok || sub == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var hu models.HouseholdUser
	if err := database.DB.Where("user_id = ?", sub).First(&hu).Error; err != nil {
		return c.JSON(nil)
	}

	var household models.Household
	if err := database.DB.
		Preload("Users").
		Preload("Users.Emails").
		Preload("Users.Phones").
		Preload("Players").
		First(&household, "id = ?", hu.HouseholdID).Error; err != nil {
		return c.JSON(nil)
	}

	return c.JSON(household)
}"""
new_get_my_household = """func GetMyHousehold(c *fiber.Ctx) error {
	sub, ok := c.Locals("user_id").(string)
	if !ok || sub == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var household models.Household
	if err := database.DB.
		Preload("Owner").
		Preload("Owner.Emails").
		Preload("Owner.Phones").
		Preload("Players").
		First(&household, "owner_id = ?", sub).Error; err != nil {
		return c.JSON(nil)
	}

	return c.JSON(household)
}"""
content = content.replace(old_get_my_household, new_get_my_household)

# 3. Update AddPlayer
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
	}"""
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
			Relationship: req.Relationship, // We will add this to AddPlayerRequest
		}
		if err := tx.Create(&hp).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create player"})
	}"""
content = content.replace(old_add_player, new_add_player)

# Also update AddPlayerRequest to include Relationship
old_add_req = """type AddPlayerRequest struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	DOB            string `json:"dob"`
	GraduationYear string `json:"graduation_year"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
}"""
new_add_req = """type AddPlayerRequest struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	DOB            string `json:"dob"`
	GraduationYear string `json:"graduation_year"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Relationship   string `json:"relationship"`
}"""
content = content.replace(old_add_req, new_add_req)

with open('apps/api/handlers/account.go', 'w') as f:
    f.write(content)

