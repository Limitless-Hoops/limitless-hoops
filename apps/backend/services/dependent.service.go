package services

import (
	"github.com/Limitless-Hoops/limitless-hoops/database"
	"github.com/Limitless-Hoops/limitless-hoops/dto"
	"github.com/Limitless-Hoops/limitless-hoops/models"
	"gorm.io/gorm"
)

// GetAllDependentsWithGuardianCount returns all dependents with the number of guardians
func GetAllDependentsWithGuardianCount() ([]dto.DependentWithCountDTO, error) {
	var dependents []models.Dependent
	err := database.DB.
		Preload("GuardianLinks", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "dependent_id")
		}).
		Find(&dependents).Error
	if err != nil {
		return nil, err
	}

	var result []dto.DependentWithCountDTO
	for _, d := range dependents {
		result = append(result, dto.DependentWithCountDTO{
			ID:             d.ID,
			FirstName:      d.FirstName,
			LastName:       d.LastName,
			Email:          d.Email,
			PhoneNumber:    d.PhoneNumber,
			MembershipTier: d.MembershipTier,
			GuardianCount:  len(d.GuardianLinks),
		})
	}

	return result, nil
}

// GetDependentByID returns basic profile info for a dependent
func GetDependentByID(dependentID uint) (*dto.DependentBasicDTO, error) {
	var dep models.Dependent
	if err := database.DB.First(&dep, dependentID).Error; err != nil {
		return nil, err
	}

	return &dto.DependentBasicDTO{
		ID:          dep.ID,
		FirstName:   dep.FirstName,
		LastName:    dep.LastName,
		Email:       dep.Email,
		PhoneNumber: dep.PhoneNumber,
	}, nil
}

// CreateDependent inserts a new dependent
func CreateDependent(dep *models.Dependent) error {
	return database.DB.Create(dep).Error
}

// UpdateDependent updates allowed fields for a dependent
func UpdateDependent(dependentID uint, updates map[string]interface{}) error {
	return database.DB.Model(&models.Dependent{}).
		Where("id = ?", dependentID).
		Updates(updates).Error
}

// FindDependentByEmail returns a dependent by email
func FindDependentByEmail(email string) (*models.Dependent, error) {
	var dep models.Dependent
	if err := database.DB.Where("email = ?", email).First(&dep).Error; err != nil {
		return nil, err
	}
	return &dep, nil
}

// FindDependentByID returns a dependent by ID
func FindDependentByID(id uint) (*models.Dependent, error) {
	var dep models.Dependent
	if err := database.DB.First(&dep, id).Error; err != nil {
		return nil, err
	}
	return &dep, nil
}

// GetDependentBasicInfo returns basic profile DTO for a dependent
func GetDependentBasicInfo(id uint) (*dto.DependentBasicDTO, error) {
	var dependent models.Dependent
	if err := database.DB.First(&dependent, id).Error; err != nil {
		return nil, err
	}

	return &dto.DependentBasicDTO{
		ID:          dependent.ID,
		FirstName:   dependent.FirstName,
		LastName:    dependent.LastName,
		Email:       dependent.Email,
		PhoneNumber: dependent.PhoneNumber,
	}, nil
}
