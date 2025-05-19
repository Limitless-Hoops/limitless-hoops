package services

import (
	"github.com/Limitless-Hoops/limitless-hoops/database"
	"github.com/Limitless-Hoops/limitless-hoops/dto"
	"github.com/Limitless-Hoops/limitless-hoops/models"
	"gorm.io/gorm"
)

// GetAllAdminsWithDependentCount returns all admins and their dependent counts
func GetAllAdminsWithDependentCount() ([]dto.AdminWithCountDTO, error) {
	var admins []models.Admin
	err := database.DB.Preload("Dependents", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "admin_id")
	}).Find(&admins).Error
	if err != nil {
		return nil, err
	}

	var result []dto.AdminWithCountDTO
	for _, a := range admins {
		result = append(result, dto.AdminWithCountDTO{
			ID:             a.ID,
			FirstName:      a.FirstName,
			LastName:       a.LastName,
			Email:          a.Email,
			PhoneNumber:    a.PhoneNumber,
			Role:           a.Role,
			DependentCount: len(a.Dependents),
		})
	}
	return result, nil
}

// GetAdminByID returns a basic admin profile
func GetAdminByID(adminID uint) (*dto.AdminBasicDTO, error) {
	var admin models.Admin
	if err := database.DB.First(&admin, adminID).Error; err != nil {
		return nil, err
	}
	return &dto.AdminBasicDTO{
		ID:          admin.ID,
		FirstName:   admin.FirstName,
		LastName:    admin.LastName,
		Email:       admin.Email,
		PhoneNumber: admin.PhoneNumber,
		Role:        admin.Role,
	}, nil
}

// CreateAdmin inserts a new admin
func CreateAdmin(admin *models.Admin) error {
	return database.DB.Create(admin).Error
}

// UpdateAdmin modifies an existing admin's allowed fields
func UpdateAdmin(adminID uint, updates map[string]interface{}) error {
	tx := database.DB.Model(&models.Admin{}).Where("id = ?", adminID).Updates(updates)
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return tx.Error
}

// FindAdminByEmail returns an admin by email
func FindAdminByEmail(email string) (*models.Admin, error) {
	var admin models.Admin
	if err := database.DB.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

// FindAdminByID returns an admin by ID
func FindAdminByID(id uint) (*models.Admin, error) {
	var admin models.Admin
	if err := database.DB.First(&admin, id).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

// GetAdminBasicInfo returns a lightweight admin profile for /me endpoint
func GetAdminBasicInfo(id uint) (*dto.AdminBasicDTO, error) {
	var admin models.Admin
	if err := database.DB.First(&admin, id).Error; err != nil {
		return nil, err
	}

	return &dto.AdminBasicDTO{
		ID:          admin.ID,
		FirstName:   admin.FirstName,
		LastName:    admin.LastName,
		Email:       admin.Email,
		PhoneNumber: admin.PhoneNumber,
		Role:        admin.Role,
	}, nil
}
