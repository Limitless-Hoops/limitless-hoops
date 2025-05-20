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
	PasswordHash string     `gorm:"not null" json:"-"`
	Role         string     `gorm:"default:staff" json:"role"` // e.g. staff, admin, superadmin
	DateOfBirth  *time.Time `gorm:"not null" json:"date_of_birth"`

	LastLogin    *time.Time `json:"last_login,omitempty"`
	LastActiveAt *time.Time `json:"last_active_at,omitempty"`

	Dependents []Dependent `gorm:"foreignKey:AdminID" json:"dependents"`
}
