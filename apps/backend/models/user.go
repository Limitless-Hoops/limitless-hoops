package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	FirstName      string     `gorm:"not null" json:"first_name"`
	LastName       string     `gorm:"not null" json:"last_name"`
	Email          string     `gorm:"unique;not null" json:"email"`
	PhoneNumber    string     `gorm:"unique;not null" json:"phone_number"`
	PasswordHash   string     `gorm:"not null" json:"-"`
	MembershipTier string     `gorm:"type:varchar(20);default:'free'" json:"membership_tier"`
	DateOfBirth    *time.Time `gorm:"not null" json:"date_of_birth"`

	GuardianLinks []GuardianLink `gorm:"constraint:OnDelete:CASCADE;" json:"guardian_links"`
}
