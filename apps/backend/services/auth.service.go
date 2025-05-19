package services

import (
	"errors"
	"log"

	"github.com/Limitless-Hoops/limitless-hoops/database"
	"github.com/Limitless-Hoops/limitless-hoops/dto"
	"github.com/Limitless-Hoops/limitless-hoops/models"
	"github.com/Limitless-Hoops/limitless-hoops/utilities"
)

// Login authenticates based on role and returns JWT
func Login(dto dto.LoginDTO) (string, error) {
	var hashedPassword string
	var id uint

	log.Println("Login attempt for:", dto.Email, "as", dto.Role)

	switch dto.Role {
	case "staff", "admin", "superadmin":
		admin, err := FindAdminByEmail(dto.Email)
		if err != nil {
			log.Println("Admin not found:", err)
			return "", err
		}
		hashedPassword = admin.PasswordHash
		id = admin.ID

	case "user":
		user, err := FindUserByEmail(dto.Email)
		if err != nil {
			log.Println("User not found:", err)
			return "", err
		}
		hashedPassword = user.PasswordHash
		id = user.ID

	case "dependent":
		dep, err := FindDependentByEmail(dto.Email)
		if err != nil {
			log.Println("Dependent not found:", err)
			return "", err
		}
		hashedPassword = dep.PasswordHash
		id = dep.ID

	default:
		log.Println("Invalid role provided:", dto.Role)
		return "", errors.New("invalid role")
	}

	if err := utilities.CheckPassword(hashedPassword, dto.Password); err != nil {
		log.Println("Password check failed:", err)
		return "", errors.New("invalid credentials")
	}

	return utilities.GenerateJWT(id, dto.Role)
}

// GetProfileByRole returns DTO info for /me based on the JWT role
func GetProfileByRole(id uint, role string) (interface{}, error) {
	switch role {
	case "admin":
		return GetAdminBasicInfo(id)
	case "user":
		return GetUserBasicInfo(id)
	case "dependent":
		return GetDependentBasicInfo(id)
	default:
		return nil, errors.New("invalid role")
	}
}

// UpdatePasswordByRole hashes and updates password by role
func UpdatePasswordByRole(id uint, role, oldPassword, newPassword string) error {
	var currentHash string
	var updateFn func(uint, string) error

	switch role {
	case "admin":
		admin, err := FindAdminByID(id)
		if err != nil {
			return err
		}
		currentHash = admin.PasswordHash
		updateFn = UpdateAdminPassword
	case "user":
		user, err := FindUserByID(id)
		if err != nil {
			return err
		}
		currentHash = user.PasswordHash
		updateFn = UpdateUserPassword
	case "dependent":
		dep, err := FindDependentByID(id)
		if err != nil {
			return err
		}
		currentHash = dep.PasswordHash
		updateFn = UpdateDependentPassword
	default:
		return errors.New("invalid role")
	}

	if err := utilities.CheckPassword(currentHash, oldPassword); err != nil {
		return errors.New("old password is incorrect")
	}

	hashed, err := utilities.HashPassword(newPassword)
	if err != nil {
		return errors.New("failed to hash new password")
	}

	return updateFn(id, hashed)
}

// UpdateAdminPassword Password update helpers
func UpdateAdminPassword(id uint, hash string) error {
	return database.DB.Model(&models.Admin{}).Where("id = ?", id).Update("password_hash", hash).Error
}

func UpdateDependentPassword(id uint, hash string) error {
	return database.DB.Model(&models.Dependent{}).Where("id = ?", id).Update("password_hash", hash).Error
}

// UpdateUserPassword securely updates a user's password
func UpdateUserPassword(userID uint, hashedPassword string) error {
	return database.DB.Model(&models.User{}).Where("id = ?", userID).Update("password_hash", hashedPassword).Error
}
