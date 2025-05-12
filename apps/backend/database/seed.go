package database

import "log"

func Seed() {
	log.Println("🌱 Starting database seeding...")

	if err := AdminSeed(); err != nil {
		log.Fatalf("❌ Failed to seed Admins: %v", err)
	}
	if err := UserSeed(); err != nil {
		log.Fatalf("❌ Failed to seed Users: %v", err)
	}
	if err := DependentSeed(); err != nil {
		log.Fatalf("❌ Failed to seed Dependents: %v", err)
	}
	if err := EmergencyContactSeed(); err != nil {
		log.Fatalf("❌ Failed to seed Emergency Contacts: %v", err)
	}
}
