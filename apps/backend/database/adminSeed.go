package database

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"time"

	"github.com/Limitless-Hoops/limitless-hoops/models"
	"github.com/Limitless-Hoops/limitless-hoops/utils"
)

func AdminSeed() error {
	admins := []struct {
		FirstName   string
		LastName    string
		PhoneNumber string
		Email       string
		Password    string
		Role        string
		DateOfBirth time.Time
	}{
		{"Anthony", "Premo", "+1514814204", "superadmin@limitlesshoops.dev", "superadmin", "superadmin", time.Date(1999, 7, 12, 0, 0, 0, 0, time.UTC)},
		{"Madison", "Premo", "+13157057136", "admin@limitlesshoops.dev", "admin", "admin", time.Date(1996, 11, 3, 0, 0, 0, 0, time.UTC)},
		{"Colonel", "Sanders", "+18002255532", "staff@limitlesshoops.dev", "staff", "staff", time.Date(1950, 3, 15, 0, 0, 0, 0, time.UTC)},
	}

	for _, a := range admins {
		var existing models.Admin
		if err := DB.Where("email = ? OR phone_number = ?", a.Email, a.PhoneNumber).First(&existing).Error; err == nil {
			log.Printf("✅ Admin %s already exists. Skipping...\n", a.Email)
			continue
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		hash, err := utils.HashPassword(a.Password)
		if err != nil {
			return err
		}

		admin := models.Admin{
			FirstName:    a.FirstName,
			LastName:     a.LastName,
			PhoneNumber:  a.PhoneNumber,
			Email:        a.Email,
			PasswordHash: hash,
			Role:         a.Role,
			DateOfBirth:  &a.DateOfBirth,
		}

		if err := DB.Create(&admin).Error; err != nil {
			return err
		}

		log.Printf("✅ Admin %s seeded successfully.\n", a.Email)
	}

	return nil
}
