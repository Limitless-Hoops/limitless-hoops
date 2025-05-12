package database

import (
	"github.com/Limitless-Hoops/limitless-hoops/models"
	"log"
)

func Migrate() {
	log.Println("🚀 Starting database migrations...")

	modelsToMigrate := []interface{}{
		&models.Admin{},
		&models.User{},
		&models.Dependent{},
		&models.GuardianLink{},
		&models.EmergencyContact{},
	}

	if err := DB.AutoMigrate(modelsToMigrate...); err != nil {
		log.Fatalf("❌ Failed to migrate database models: %v", err)
	}

	log.Println("✅ Database models migrated successfully!")
}
