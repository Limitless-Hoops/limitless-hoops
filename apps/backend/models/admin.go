package models

import (
	"gorm.io/gorm"
	"time"
)

type Admin struct {
	gorm.Model
	FirstName    string     `gorm:"not null" json:"first_name"`
	LastName     string     `gorm:"not null" json:"last_name"`
	PhoneNumber  string     `gorm:"unique;not null" json:"phone_number"`
	Email        string     `gorm:"unique;not null" json:"email"`
	PasswordHash string     `gorm:"not null" json:"-" json:"-"`
	Role         string     `gorm:"default:staff" json:"role"` // e.g. staff, admin, superadmin
	DateOfBirth  *time.Time `gorm:"not null" json:"date_of_birth"`

	Dependents []Dependent `gorm:"foreignKey:AdminID" json:"dependents"`
}
