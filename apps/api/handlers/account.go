package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/limitless-hoops/api/database"
	"github.com/limitless-hoops/api/models"
	"gorm.io/gorm"
)

// DTOs for the incoming JSON
type CreateHouseholdRequest struct {
	Parent struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
		DOB       string `json:"dob"`
	} `json:"parent"`
	Players []struct {
		FirstName      string `json:"first_name"`
		LastName       string `json:"last_name"`
		GraduationYear string `json:"graduation_year"`
	} `json:"players"`
}

func CreateHousehold(c *fiber.Ctx) error {
	var req CreateHouseholdRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	sub, ok := c.Locals("user_id").(string)
	if !ok || sub == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized: missing user_id"})
	}

	email, ok := c.Locals("email").(string)
	jwtEmail := email
	if !ok || email == "" {
		email = req.Parent.Email
	}
	req.Parent.Email = email // Enforce authenticated email

	phone, ok := c.Locals("phone").(string)
	jwtPhone := phone
	if !ok || phone == "" {
		phone = req.Parent.Phone
	}
	req.Parent.Phone = phone // Enforce authenticated phone

	// Age validation (18+) — H6: use proper calendar arithmetic, not 18*365 days
	if req.Parent.DOB == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Date of birth is required"})
	}
	dobTime, err := time.Parse("2006-01-02", req.Parent.DOB)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid date format for DOB"})
	}
	now := time.Now()
	minAge := time.Date(now.Year()-18, now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	if dobTime.After(minAge) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "You must be at least 18 years old to create an account"})
	}

	// Determine source provider and auth method from JWT locals
	sourceProvider := "manual"
	if p, ok := c.Locals("provider").(string); ok && p != "" {
		sourceProvider = p
	}
	authMethod := "email"
	if m, ok := c.Locals("auth_method").(string); ok && m != "" {
		authMethod = m
	}

	notifCategories := []string{"billing", "events", "announcements", "security"}

	err = database.DB.Transaction(func(tx *gorm.DB) error {
		var user models.User

		// 1. Look up user STRICTLY by their Supabase sub ID
		if err := tx.First(&user, "id = ?", sub).Error; err == nil {
			// H8: User record exists — check if they also have a household.
			// If yes, just return (idempotent). If no, continue to create the household below.
			var existingHousehold models.Household
			if tx.Where("owner_id = ?", user.ID).First(&existingHousehold).Error == nil {
				// Already has a household — nothing to do
				return nil
			}
			// User exists but has no household — fall through to household creation below
		} else {
			// 2. User does not exist. Check if their email is already claimed by ANOTHER account.
			// C8: only check email conflict when we actually have an email
			if email != "" {
				var existingEmail models.UserEmail
				if tx.Where("email = ?", email).First(&existingEmail).Error == nil {
					// Impersonation attempt — this email is already claimed.
					// Only delete the orphaned Supabase account if the email came from their JWT
					// (i.e. they verified ownership). Don't delete on a manual form typo.
					if email == jwtEmail {
						go func(orphanSub string) {
							supabaseURL := os.Getenv("SUPABASE_URL")
							if supabaseURL == "" {
								supabaseURL = "https://rhhqnmwnckgzhdyqsspb.supabase.co"
							}
							serviceRoleKey := os.Getenv("SUPABASE_SERVICE_ROLE_KEY")
							if serviceRoleKey != "" {
								req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/auth/v1/admin/users/%s", supabaseURL, orphanSub), nil)
								req.Header.Set("Authorization", "Bearer "+serviceRoleKey)
								req.Header.Set("apikey", serviceRoleKey)
								client := &http.Client{Timeout: 5 * time.Second}
								client.Do(req)
							}
						}(sub)
					}
					return fiber.NewError(fiber.StatusConflict, "This email is already linked to another Limitless Hoops account.")
				}
			}

			// 2.5 User does not exist and email is free. Check if their phone is already claimed.
			if phone != "" {
				var existingPhone models.UserPhone
				if tx.Where("phone = ?", phone).First(&existingPhone).Error == nil {
					// Only delete the orphaned Supabase account if phone came from the JWT
					if phone == jwtPhone {
						go func(orphanSub string) {
							supabaseURL := os.Getenv("SUPABASE_URL")
							if supabaseURL == "" {
								supabaseURL = "https://rhhqnmwnckgzhdyqsspb.supabase.co"
							}
							serviceRoleKey := os.Getenv("SUPABASE_SERVICE_ROLE_KEY")
							if serviceRoleKey != "" {
								req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/auth/v1/admin/users/%s", supabaseURL, orphanSub), nil)
								req.Header.Set("Authorization", "Bearer "+serviceRoleKey)
								req.Header.Set("apikey", serviceRoleKey)
								client := &http.Client{Timeout: 5 * time.Second}
								client.Do(req)
							}
						}(sub)
					}
					return fiber.NewError(fiber.StatusConflict, "This phone number is already linked to another Limitless Hoops account.")
				}
			}

			// 3. User does not exist and email/phone are free — create new user with ID = sub
			user = models.User{
				ID:        sub,
				FirstName: req.Parent.FirstName,
				LastName:  req.Parent.LastName,
				DOB:       req.Parent.DOB,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			if err := tx.Create(&user).Error; err != nil {
				return err
			}

			// C9: Create primary UserEmail ONLY when we actually have an email address
			if email != "" {
				userEmail := models.UserEmail{
					ID:             uuid.New().String(),
					UserID:         user.ID,
					Email:          email,
					IsPrimary:      true,
					IsVerified:     (authMethod == "email" || sourceProvider == "google"),
					SourceProvider: sourceProvider,
					CreatedAt:      time.Now(),
				}
				if err := tx.Create(&userEmail).Error; err != nil {
					return err
				}

				// Seed notification preferences for email
				for _, cat := range notifCategories {
					pref := models.NotificationPreference{
						ID:         uuid.New().String(),
						UserID:     user.ID,
						TargetType: "email",
						TargetID:   userEmail.ID,
						Category:   cat,
						Enabled:    true,
					}
					if err := tx.Create(&pref).Error; err != nil {
						return err
					}
				}
			}

			// Create primary UserPhone (if phone provided)
			if phone != "" {
				userPhone := models.UserPhone{
					ID:         uuid.New().String(),
					UserID:     user.ID,
					Phone:      phone,
					IsPrimary:  true,
					IsVerified: (authMethod == "phone"),
					CreatedAt:  time.Now(),
				}
				if err := tx.Create(&userPhone).Error; err != nil {
					return err
				}

				// Seed notification preferences for phone
				for _, cat := range notifCategories {
					pref := models.NotificationPreference{
						ID:         uuid.New().String(),
						UserID:     user.ID,
						TargetType: "phone",
						TargetID:   userPhone.ID,
						Category:   cat,
						Enabled:    true,
					}
					if err := tx.Create(&pref).Error; err != nil {
						return err
					}
				}
			}
		}

		// 1. Create Household linked directly to User
		householdID := uuid.New().String()
		household := models.Household{
			ID:        householdID,
			OwnerID:   user.ID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := tx.Create(&household).Error; err != nil {
			return err
		}

		// 3. Create Players
		for _, pReq := range req.Players {
			playerID := uuid.New().String()
			player := models.Player{
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
			}
		}

		return nil
	})

	if err != nil {
		if e, ok := err.(*fiber.Error); ok {
			return c.Status(e.Code).JSON(fiber.Map{
				"error": e.Message,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create household",
			"details": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Household created successfully",
	})
}

func GetMyHousehold(c *fiber.Ctx) error {
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
}

func GetHousehold(c *fiber.Ctx) error {
	id := c.Params("id")

	var household models.Household
	// Preload the related Users and Players
	if err := database.DB.Preload("Users").Preload("Players").First(&household, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Household not found",
		})
	}

	return c.JSON(household)
}

type AddPlayerRequest struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	DOB            string `json:"dob"`
	GraduationYear string `json:"graduation_year"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Relationship   string `json:"relationship"`
}

func AddPlayer(c *fiber.Ctx) error {
	householdID := c.Params("id")
	var req AddPlayerRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if req.DOB != "" {
		dobTime, err := time.Parse("2006-01-02", req.DOB)
		if err == nil && dobTime.After(time.Now()) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Date of birth cannot be in the future"})
		}
	}

	playerID := uuid.New().String()
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create player",
			"details": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(player)
}

type AddPlayerContactRequest struct {
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

type UpdateMyPersonalInfoRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	DOB       string `json:"dob"`
}

func UpdateMyPersonalInfo(c *fiber.Ctx) error {
	sub, _ := c.Locals("user_id").(string)
	
	var req UpdateMyPersonalInfoRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	updates := map[string]interface{}{}
	if req.FirstName != "" { updates["first_name"] = req.FirstName }
	if req.LastName != "" { updates["last_name"] = req.LastName }
	if req.DOB != "" { updates["dob"] = req.DOB }

	if len(updates) > 0 {
		if err := database.DB.Model(&models.User{}).Where("id = ?", sub).Updates(updates).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update profile"})
		}
	}

	return c.JSON(fiber.Map{"message": "Profile updated"})
}
