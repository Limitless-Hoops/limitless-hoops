package database

import (
	"log"
	"time"

	"github.com/Limitless-Hoops/limitless-hoops/models"
)

func EmergencyContactSeed() error {
	var dependents []models.Dependent
	if err := DB.Find(&dependents).Error; err != nil {
		return err
	}

	type contactSeed struct {
		FirstName   string
		LastName    string
		PhoneNumber string
		Relation    string
		DateOfBirth time.Time
	}

	emergencySets := map[string][]contactSeed{
		"Jimmy": {
			{"Linda", "Smith", "+15556660001", "grandmother", time.Date(1960, 5, 12, 0, 0, 0, 0, time.UTC)},
			{"Coach", "Tom", "+15556660002", "coach", time.Date(1978, 4, 9, 0, 0, 0, 0, time.UTC)},
		},
		"Ella": {
			{"Nina", "Clark", "+15556660003", "aunt", time.Date(1982, 8, 20, 0, 0, 0, 0, time.UTC)},
			{"Uncle", "Jack", "+15556660004", "uncle", time.Date(1980, 2, 14, 0, 0, 0, 0, time.UTC)},
			{"Family", "Doctor", "+15556660005", "doctor", time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)},
		},
		"Leo": {
			{"Friend", "Alex", "+15556660006", "family friend", time.Date(1991, 10, 30, 0, 0, 0, 0, time.UTC)},
		},
	}

	for _, dep := range dependents {
		contacts, ok := emergencySets[dep.FirstName]
		if !ok {
			log.Printf("⚠️ No emergency contacts defined for dependent %s %s\n", dep.FirstName, dep.LastName)
			continue
		}

		for _, c := range contacts {
			contact := models.EmergencyContact{
				FirstName:   c.FirstName,
				LastName:    c.LastName,
				PhoneNumber: c.PhoneNumber,
				Relation:    c.Relation,
				DateOfBirth: &c.DateOfBirth,
			}

			if err := DB.Create(&contact).Error; err != nil {
				return err
			}

			if err := DB.Model(&dep).Association("EmergencyContacts").Append(&contact); err != nil {
				return err
			}
		}

		log.Printf("✅ Assigned %d emergency contacts to %s %s\n", len(contacts), dep.FirstName, dep.LastName)
	}

	return nil
}
