package services

import (
	"time"

	"github.com/Limitless-Hoops/limitless-hoops/database"
	"github.com/Limitless-Hoops/limitless-hoops/models"
	"gorm.io/gorm"
)

type UserWithCount struct {
	ID             uint   `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
	MembershipTier string `json:"membership_tier"`
	DependentCount int    `json:"dependent_count"`
}

func GetAllUsersWithDependentCount() ([]UserWithCount, error) {
	var users []models.User
	err := database.DB.
		Preload("GuardianLinks", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "user_id")
		}).
		Find(&users).Error
	if err != nil {
		return nil, err
	}

	var result []UserWithCount
	for _, u := range users {
		result = append(result, UserWithCount{
			ID:             u.ID,
			FirstName:      u.FirstName,
			LastName:       u.LastName,
			Email:          u.Email,
			PhoneNumber:    u.PhoneNumber,
			MembershipTier: u.MembershipTier,
			DependentCount: len(u.GuardianLinks),
		})
	}
	return result, nil
}

type EmergencyContactResponse struct {
	ID          uint       `json:"id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	PhoneNumber string     `json:"phone_number"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	Relation    string     `json:"relation"`
}

type DependentWithContacts struct {
	ID                uint                       `json:"id"`
	FirstName         string                     `json:"first_name"`
	LastName          string                     `json:"last_name"`
	Email             *string                    `json:"email"`
	PhoneNumber       *string                    `json:"phone_number"`
	DateOfBirth       *time.Time                 `json:"date_of_birth"`
	MembershipTier    string                     `json:"membership_tier"`
	EmergencyContacts []EmergencyContactResponse `json:"emergency_contacts"`
}

type UserWithDependents struct {
	ID             uint                    `json:"id"`
	FirstName      string                  `json:"first_name"`
	LastName       string                  `json:"last_name"`
	Email          string                  `json:"email"`
	PhoneNumber    string                  `json:"phone_number"`
	MembershipTier string                  `json:"membership_tier"`
	DateOfBirth    *time.Time              `json:"date_of_birth"`
	Dependents     []DependentWithContacts `json:"dependents"`
}

func GetUserByIDWithDependentsAndContacts(userID uint) (*UserWithDependents, error) {
	var user models.User
	err := database.DB.
		Preload("GuardianLinks.Dependent.EmergencyContacts").
		First(&user, userID).Error
	if err != nil {
		return nil, err
	}

	var dependents []DependentWithContacts
	for _, link := range user.GuardianLinks {
		d := link.Dependent

		var cleanContacts []EmergencyContactResponse
		for _, ec := range d.EmergencyContacts {
			cleanContacts = append(cleanContacts, EmergencyContactResponse{
				ID:          ec.ID,
				FirstName:   ec.FirstName,
				LastName:    ec.LastName,
				PhoneNumber: ec.PhoneNumber,
				DateOfBirth: ec.DateOfBirth,
				Relation:    ec.Relation,
			})
		}

		dependents = append(dependents, DependentWithContacts{
			ID:                d.ID,
			FirstName:         d.FirstName,
			LastName:          d.LastName,
			Email:             d.Email,
			PhoneNumber:       d.PhoneNumber,
			DateOfBirth:       d.DateOfBirth,
			MembershipTier:    d.MembershipTier,
			EmergencyContacts: cleanContacts,
		})
	}

	return &UserWithDependents{
		ID:             user.ID,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		PhoneNumber:    user.PhoneNumber,
		MembershipTier: user.MembershipTier,
		DateOfBirth:    user.DateOfBirth,
		Dependents:     dependents,
	}, nil
}

func GetDependentsForUser(userID uint) ([]DependentWithContacts, error) {
	var links []models.GuardianLink
	err := database.DB.
		Preload("Dependent.EmergencyContacts").
		Where("user_id = ?", userID).
		Find(&links).Error
	if err != nil {
		return nil, err
	}

	var dependents []DependentWithContacts
	for _, link := range links {
		d := link.Dependent

		var cleanContacts []EmergencyContactResponse
		for _, ec := range d.EmergencyContacts {
			cleanContacts = append(cleanContacts, EmergencyContactResponse{
				ID:          ec.ID,
				FirstName:   ec.FirstName,
				LastName:    ec.LastName,
				PhoneNumber: ec.PhoneNumber,
				DateOfBirth: ec.DateOfBirth,
				Relation:    ec.Relation,
			})
		}

		dependents = append(dependents, DependentWithContacts{
			ID:                d.ID,
			FirstName:         d.FirstName,
			LastName:          d.LastName,
			Email:             d.Email,
			PhoneNumber:       d.PhoneNumber,
			DateOfBirth:       d.DateOfBirth,
			MembershipTier:    d.MembershipTier,
			EmergencyContacts: cleanContacts,
		})
	}

	return dependents, nil
}

func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func UpdateUser(userID uint, updates map[string]interface{}) error {
	return database.DB.
		Model(&models.User{}).
		Where("id = ?", userID).
		Updates(updates).Error
}
