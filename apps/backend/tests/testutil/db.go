package testutil

import (
	"context"
	"fmt"
	"github.com/Limitless-Hoops/limitless-hoops/models"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	testDB      *gorm.DB
	dbContainer testcontainers.Container
)

// ConnectTestDB spins up a temporary Postgres container and connects GORM to it.
func ConnectTestDB() *gorm.DB {
	if testDB != nil {
		return testDB
	}

	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image: "postgres:17.4-alpine3.21",
		Env: map[string]string{
			"POSTGRES_DB":       "limitless_test",
			"POSTGRES_USER":     "test_user",
			"POSTGRES_PASSWORD": "test_password",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	var err error
	dbContainer, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("❌ Failed to start test container: %v", err)
	}

	host, _ := dbContainer.Host(ctx)
	port, _ := dbContainer.MappedPort(ctx, "5432")

	dsn := fmt.Sprintf("host=%s user=test_user password=test_password dbname=limitless_test port=%s sslmode=disable", host, port.Port())
	testDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to test DB: %v", err)
	}

	if err := RunTestMigrations(testDB); err != nil {
		log.Fatalf("❌ Migration error: %v", err)
	}

	return testDB
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

func ResetTestDB(db *gorm.DB) error {
	log.Println("🔄 Resetting test database...")

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

// TearDownTestDB stops and removes the Postgres container after all tests finish.
func TearDownTestDB() {
	if dbContainer != nil {
		_ = dbContainer.Terminate(context.Background())
	}
}
