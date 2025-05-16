package services

import (
	"github.com/Limitless-Hoops/limitless-hoops/database"
	"github.com/Limitless-Hoops/limitless-hoops/dto"
	"github.com/Limitless-Hoops/limitless-hoops/models"
	"gorm.io/gorm"
)

// GetAllUsersWithDependentCount returns a lightweight user list
func GetAllUsersWithDependentCount() ([]dto.UserWithCountDTO, error) {
	var users []models.User
	err := database.DB.
		Preload("GuardianLinks", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "user_id")
		}).
		Find(&users).Error
	if err != nil {
		return nil, err
	}

	var result []dto.UserWithCountDTO
	for _, u := range users {
		result = append(result, dto.UserWithCountDTO{
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

// GetUserByIDWithDependentsAndContacts returns a full user profile
func GetUserByIDWithDependentsAndContacts(userID uint) (dto.UserWithDependentsDTO, error) {
	var user models.User
	err := database.DB.
		Preload("GuardianLinks.Dependent.EmergencyContacts").
		First(&user, userID).Error
	if err != nil {
		return dto.UserWithDependentsDTO{}, err
	}

	var dependents []dto.DependentDTO
	for _, link := range user.GuardianLinks {
		d := link.Dependent
		var contacts []dto.EmergencyContactDTO
		for _, ec := range d.EmergencyContacts {
			contacts = append(contacts, dto.EmergencyContactDTO{
				ID:          ec.ID,
				FirstName:   ec.FirstName,
				LastName:    ec.LastName,
				PhoneNumber: ec.PhoneNumber,
				DateOfBirth: ec.DateOfBirth,
				Relation:    ec.Relation,
			})
		}
		dependents = append(dependents, dto.DependentDTO{
			ID:                d.ID,
			FirstName:         d.FirstName,
			LastName:          d.LastName,
			Email:             d.Email,
			PhoneNumber:       d.PhoneNumber,
			DateOfBirth:       d.DateOfBirth,
			MembershipTier:    d.MembershipTier,
			EmergencyContacts: contacts,
		})
	}

	return dto.UserWithDependentsDTO{
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

// GetDependentsForUser returns dependents for a user
func GetDependentsForUser(userID uint) ([]dto.DependentDTO, error) {
	var links []models.GuardianLink
	err := database.DB.
		Preload("Dependent.EmergencyContacts").
		Where("user_id = ?", userID).
		Find(&links).Error
	if err != nil {
		return nil, err
	}

	var dependents []dto.DependentDTO
	for _, link := range links {
		d := link.Dependent
		var contacts []dto.EmergencyContactDTO
		for _, ec := range d.EmergencyContacts {
			contacts = append(contacts, dto.EmergencyContactDTO{
				ID:          ec.ID,
				FirstName:   ec.FirstName,
				LastName:    ec.LastName,
				PhoneNumber: ec.PhoneNumber,
				DateOfBirth: ec.DateOfBirth,
				Relation:    ec.Relation,
			})
		}
		dependents = append(dependents, dto.DependentDTO{
			ID:                d.ID,
			FirstName:         d.FirstName,
			LastName:          d.LastName,
			Email:             d.Email,
			PhoneNumber:       d.PhoneNumber,
			DateOfBirth:       d.DateOfBirth,
			MembershipTier:    d.MembershipTier,
			EmergencyContacts: contacts,
		})
	}

	return dependents, nil
}

// CreateUser inserts a new user into the database
func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

// UpdateUser modifies an existing user
func UpdateUser(userID uint, updates map[string]interface{}) error {
	return database.DB.
		Model(&models.User{}).
		Where("id = ?", userID).
		Updates(updates).Error
}

// FindUserByEmail returns a user by email
func FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindUserByID returns a user by ID
func FindUserByID(userID uint) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserBasicInfo returns basic DTO info for self-profile endpoint
func GetUserBasicInfo(userID uint) (*dto.UserWithCountDTO, error) {
	var user models.User
	err := database.DB.
		Preload("GuardianLinks", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "user_id")
		}).
		First(&user, userID).Error
	if err != nil {
		return nil, err
	}

	userWithCount := &dto.UserWithCountDTO{
		ID:             user.ID,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		PhoneNumber:    user.PhoneNumber,
		MembershipTier: user.MembershipTier,
		DependentCount: len(user.GuardianLinks),
	}

	return userWithCount, nil
}
