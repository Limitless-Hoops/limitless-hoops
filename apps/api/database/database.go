package database

import (
	"fmt"
	"log"
	"os"

	"github.com/limitless-hoops/api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("Connected to PostgreSQL Database!")
	DB = db

	log.Println("Running AutoMigrate...")
	err = DB.AutoMigrate(
		&models.User{},
		&models.UserEmail{},
		&models.UserPhone{},
		&models.ContactVerification{},
		&models.NotificationPreference{},
		&models.Household{},
		&models.HouseholdPlayer{},
		&models.PlayerContact{},
		&models.Player{},
		&models.Team{},
		&models.Application{},
		&models.Membership{},
		&models.SeatInventory{},
		&models.PaymentLedger{},
		&models.StripeEvent{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database. \n", err)
	}
	log.Println("AutoMigrate completed successfully!")
}
