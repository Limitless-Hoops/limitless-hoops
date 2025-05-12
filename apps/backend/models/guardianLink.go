package models

import "gorm.io/gorm"

type GuardianLink struct {
	gorm.Model
	UserID      uint `gorm:"not null" json:"user_id"`
	DependentID uint `gorm:"not null" json:"dependent_id"`

	User      User      `gorm:"foreignKey:UserID" json:"user"`
	Dependent Dependent `gorm:"foreignKey:DependentID" json:"dependent"`

	Relation    string `gorm:"not null" json:"relation"` // e.g., mother, father, guardian
	AccessLevel string `gorm:"default:'full'" json:"access_level"`
	IsPrimary   bool   `gorm:"default:false" json:"is_primary"`
}
