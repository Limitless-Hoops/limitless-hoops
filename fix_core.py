import re

with open('apps/api/models/core.go', 'r') as f:
    content = f.read()

# Replace Household to the end
old_block = content[content.find('type Household struct {'):]

new_block = """type Household struct {
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
"""

content = content.replace(old_block, new_block)

with open('apps/api/models/core.go', 'w') as f:
    f.write(content)

