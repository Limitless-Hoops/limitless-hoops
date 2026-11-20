package models

import "time"

type ContactVerification struct {
	ID         string    `gorm:"primaryKey;type:varchar(255)" json:"id"`
	TargetID   string    `gorm:"type:varchar(255);index" json:"target_id"` // UserEmail.ID or UserPhone.ID
	TargetType string    `gorm:"type:varchar(50)" json:"target_type"`      // "email" or "phone"
	Code       string    `gorm:"type:varchar(10)" json:"code"`
	ExpiresAt  time.Time `json:"expires_at"`
	CreatedAt  time.Time `json:"created_at"`
}
