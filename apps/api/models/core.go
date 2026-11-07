package models

import (
	"time"
)

type User struct {
	ID        string      `gorm:"primaryKey;type:varchar(255)" json:"id"`
	FirstName string      `gorm:"type:varchar(255)" json:"first_name"`
	LastName  string      `gorm:"type:varchar(255)" json:"last_name"`
	DOB       string      `gorm:"type:varchar(255)" json:"dob"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Emails    []UserEmail `gorm:"foreignKey:UserID" json:"emails,omitempty"`
	Phones    []UserPhone `gorm:"foreignKey:UserID" json:"phones,omitempty"`
}

type UserEmail struct {
	ID             string    `gorm:"primaryKey;type:varchar(255)" json:"id"`
	UserID         string    `gorm:"type:varchar(255);uniqueIndex:idx_user_email_unique;index" json:"user_id"`
	Email          string    `gorm:"type:varchar(255);uniqueIndex:idx_user_email_unique" json:"email"`
	IsPrimary      bool      `gorm:"default:false" json:"is_primary"`
	IsVerified     bool      `gorm:"default:false" json:"is_verified"`
	SourceProvider string    `gorm:"type:varchar(255);default:'manual'" json:"source_provider"` // 'manual', 'google', 'facebook'
	Label          string    `gorm:"type:varchar(255)" json:"label"`
	CreatedAt      time.Time `json:"created_at"`
}

type UserPhone struct {
	ID         string    `gorm:"primaryKey;type:varchar(255)" json:"id"`
	UserID     string    `gorm:"type:varchar(255);uniqueIndex:idx_user_phone_unique;index" json:"user_id"`
	Phone      string    `gorm:"type:varchar(255);uniqueIndex:idx_user_phone_unique" json:"phone"`
	IsPrimary  bool      `gorm:"default:false" json:"is_primary"`
	IsVerified bool      `gorm:"default:false" json:"is_verified"`
	Label      string    `gorm:"type:varchar(255)" json:"label"`
	CreatedAt  time.Time `json:"created_at"`
}

type NotificationPreference struct {
	ID         string `gorm:"primaryKey;type:varchar(255)" json:"id"`
	UserID     string `gorm:"type:varchar(255);index" json:"user_id"`
	TargetType string `gorm:"type:varchar(255)" json:"target_type"` // 'email' or 'phone'
	TargetID   string `gorm:"type:varchar(255)" json:"target_id"`
	Category   string `gorm:"type:varchar(255)" json:"category"` // 'billing', 'events', 'announcements', 'security'
	Enabled    bool   `gorm:"default:true" json:"enabled"`
}

type Household struct {
	ID        string    `gorm:"primaryKey;type:varchar(255)" json:"id"`
	OwnerID   string    `gorm:"type:varchar(255);index" json:"owner_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
	// Relations
	Owner   User     `gorm:"foreignKey:OwnerID" json:"owner"`
	Players []Player `gorm:"many2many:household_players;" json:"players"`
}

type HouseholdPlayer struct {
	HouseholdID  string `gorm:"primaryKey;type:varchar(255)" json:"household_id"`
	PlayerID     string `gorm:"primaryKey;type:varchar(255)" json:"player_id"`
	Relationship string `gorm:"type:varchar(255)" json:"relationship"` // e.g., "Mother", "Father"
}

type Player struct {
	ID                  string    `gorm:"primaryKey;type:varchar(255)" json:"id"` // Maps to legacy PersonID
	FirstName           string    `gorm:"type:varchar(255)" json:"first_name"`
	LastName            string    `gorm:"type:varchar(255)" json:"last_name"`
	DOB                 string    `gorm:"type:varchar(255)" json:"dob"`
	Email               string    `gorm:"type:varchar(255)" json:"email"` // Optional player email
	Phone               string    `gorm:"type:varchar(255)" json:"phone"` // Optional player phone
	Grade2027_28        string    `gorm:"type:varchar(255)" json:"grade_2027_28"`
	ReferenceSchoolYear string    `gorm:"type:varchar(255)" json:"reference_school_year"`
	GraduationYear      string    `gorm:"type:varchar(255)" json:"graduation_year"`
	PublicDisplayName   string    `gorm:"type:varchar(255)" json:"public_display_name"`
	ReferralCode        string    `gorm:"type:varchar(255)" json:"referral_code"`
	Venmo               string    `gorm:"type:varchar(255)" json:"venmo"`
	Status              string    `gorm:"type:varchar(255)" json:"status"`
	Notes               string    `gorm:"type:text" json:"notes"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	
	// Relations
	Contacts []PlayerContact `gorm:"foreignKey:PlayerID" json:"contacts"`
}

type PlayerContact struct {
	ID           string    `gorm:"primaryKey;type:varchar(255)" json:"id"`
	PlayerID     string    `gorm:"type:varchar(255);index" json:"player_id"`
	FirstName    string    `gorm:"type:varchar(255)" json:"first_name"`
	LastName     string    `gorm:"type:varchar(255)" json:"last_name"`
	Phone        string    `gorm:"type:varchar(255)" json:"phone"`
	Email        string    `gorm:"type:varchar(255)" json:"email"`
	Relationship string    `gorm:"type:varchar(255)" json:"relationship"` // e.g., "Grandmother"
	ReceiveSMS   bool      `gorm:"default:true" json:"receive_sms"`
	ReceiveEmail bool      `gorm:"default:true" json:"receive_email"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
