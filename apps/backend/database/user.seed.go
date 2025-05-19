package database

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"time"

	"github.com/Limitless-Hoops/limitless-hoops/models"
	"github.com/Limitless-Hoops/limitless-hoops/utilities"
)

func UserSeed() error {

	users := []struct {
		FirstName      string
		LastName       string
		Email          string
		PhoneNumber    string
		Password       string
		MembershipTier string
		DateOfBirth    time.Time
	}{
		{"Sarah", "Smith", "sarah@limitlesshoops.dev", "+10000000001", "password1", "basic", time.Date(1988, 6, 15, 0, 0, 0, 0, time.UTC)},
		{"Mike", "Johnson", "mike@limitlesshoops.dev", "+10000000002", "password2", "prime", time.Date(1985, 9, 27, 0, 0, 0, 0, time.UTC)},
		{"Ava", "Clark", "ava@limitlesshoops.dev", "+10000000003", "password3", "free", time.Date(1990, 3, 4, 0, 0, 0, 0, time.UTC)},
	}

	for _, u := range users {
		var existing models.User
		if err := DB.Where("email = ? OR phone_number = ?", u.Email, u.PhoneNumber).First(&existing).Error; err == nil {
			log.Printf("✅ User %s already exists. Skipping...\n", u.Email)
			continue
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		hash, err := utilities.HashPassword(u.Password)
		if err != nil {
			return err
		}
		log.Println("🔐 Storing hash for", u.Email, "→", hash)

		user := models.User{
			FirstName:      u.FirstName,
			LastName:       u.LastName,
			Email:          u.Email,
			PhoneNumber:    u.PhoneNumber,
			PasswordHash:   hash,
			MembershipTier: u.MembershipTier,
			DateOfBirth:    &u.DateOfBirth,
		}

		if err := DB.Create(&user).Error; err != nil {
			log.Fatalf("Failed to seed user %s: %v", u.Email, err)
		}

		log.Printf("✅ User %s seeded.\n", u.Email)
	}

	return nil
}
