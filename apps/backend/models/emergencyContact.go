package models

import (
	"gorm.io/gorm"
	"time"
)

type EmergencyContact struct {
	gorm.Model
	FirstName   string     `gorm:"not null" json:"first_name"`
	LastName    string     `gorm:"not null" json:"last_name"`
	PhoneNumber string     `gorm:"not null" json:"phone_number"`
	DateOfBirth *time.Time `gorm:"not null" json:"date_of_birth"`
	Relation    string     `gorm:"not null" json:"relation"` // e.g. "mother", "uncle"
}
