package testsetup

import (
	"github.com/Limitless-Hoops/limitless-hoops/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectTestDB() *gorm.DB {
	dsn := "host=localhost user=test_user password=test_password dbname=limitless_test port=5433 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to test database: %v", err)
	}

	err = RunTestMigrations(db)
	if err != nil {
		log.Fatalf("❌ Failed to run test migrations: %v", err)
	}

	return db
}

func RunTestMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Admin{},
		&models.User{},
		&models.Dependent{},
		&models.GuardianLink{},
		&models.EmergencyContact{},
	)

}

// ResetTestDB truncates all test tables and resets their primary key sequences.
func ResetTestDB(db *gorm.DB) error {
	log.Println("🔄 Resetting test database...")

	// Use raw SQL to truncate all tables and reset auto-incrementing IDs
	return db.Exec(`
		TRUNCATE TABLE 
			admins, 
			users, 
			dependents, 
			guardian_links, 
			emergency_contacts 
		RESTART IDENTITY CASCADE;
	`).Error
}
