package database

import (
	"errors"
	"log"
	"time"

	"github.com/Limitless-Hoops/limitless-hoops/models"
	"github.com/Limitless-Hoops/limitless-hoops/utils"
	"gorm.io/gorm"
)

func DependentSeed() error {
	// Fetch guardian users
	var sarah, mike, ava models.User
	if err := DB.First(&sarah, "email = ?", "sarah@limitlesshoops.dev").Error; err != nil {
		return err
	}
	if err := DB.First(&mike, "email = ?", "mike@limitlesshoops.dev").Error; err != nil {
		return err
	}
	if err := DB.First(&ava, "email = ?", "ava@limitlesshoops.dev").Error; err != nil {
		return err
	}

	// Fetch coaches (admins)
	var anthony, madison models.Admin
	if err := DB.First(&anthony, "email = ?", "superadmin@limitlesshoops.dev").Error; err != nil {
		return err
	}
	if err := DB.First(&madison, "email = ?", "admin@limitlesshoops.dev").Error; err != nil {
		return err
	}

	// Define dependents with their guardians and assigned coach
	dependents := []struct {
		FirstName      string
		LastName       string
		Email          string
		Phone          string
		Password       string
		MembershipTier string
		DateOfBirth    time.Time
		Guardians      []models.User
		Coach          models.Admin
	}{
		{"Jimmy", "Smith", "jimmy@limitlesshoops.dev", "+20000000001", "kidpass1", "basic", time.Date(2012, 4, 10, 0, 0, 0, 0, time.UTC), []models.User{sarah, mike}, anthony},
		{"Ella", "Clark", "ella@limitlesshoops.dev", "+20000000002", "kidpass2", "prime", time.Date(2013, 6, 3, 0, 0, 0, 0, time.UTC), []models.User{ava}, madison},
		{"Leo", "Clark", "leo@limitlesshoops.dev", "+20000000003", "kidpass3", "free", time.Date(2015, 1, 22, 0, 0, 0, 0, time.UTC), []models.User{ava}, madison},
	}

	for _, d := range dependents {
		var existing models.Dependent
		if err := DB.Where("email = ? OR phone_number = ?", d.Email, d.Phone).First(&existing).Error; err == nil {
			log.Printf("✅ Dependent %s %s already exists. Skipping...\n", d.FirstName, d.LastName)
			continue
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		hash, err := utils.HashPassword(d.Password)
		if err != nil {
			return err
		}

		dependent := models.Dependent{
			FirstName:      d.FirstName,
			LastName:       d.LastName,
			Email:          &d.Email,
			PhoneNumber:    &d.Phone,
			PasswordHash:   hash,
			DateOfBirth:    &d.DateOfBirth,
			MembershipTier: d.MembershipTier,
			AdminID:        &d.Coach.ID,
		}

		if err := DB.Create(&dependent).Error; err != nil {
			return err
		}

		for i, guardian := range d.Guardians {
			link := models.GuardianLink{
				UserID:      guardian.ID,
				DependentID: dependent.ID,
				Relation:    "parent",
				IsPrimary:   i == 0,
				AccessLevel: "full",
			}
			if err := DB.Create(&link).Error; err != nil {
				return err
			}
		}

		log.Printf("✅ Dependent %s %s seeded and linked.\n", d.FirstName, d.LastName)
	}

	return nil
}
