package controllers

import val "github.com/Limitless-Hoops/limitless-hoops/validator"
import (
	"strconv"

	"github.com/Limitless-Hoops/limitless-hoops/dto"
	"github.com/Limitless-Hoops/limitless-hoops/models"
	"github.com/Limitless-Hoops/limitless-hoops/services"
	"github.com/Limitless-Hoops/limitless-hoops/utils"
	"github.com/gofiber/fiber/v2"
)

// GetUsers returns a lightweight list of users with dependent counts
func GetUsers(c *fiber.Ctx) error {
	users, err := services.GetAllUsersWithDependentCount()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

// GetUserByID returns a full user profile with their dependents and emergency contacts
func GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	user, err := services.GetUserByIDWithDependentsAndContacts(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}

// GetUserDependents returns just the dependents of a given user
func GetUserDependents(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	dependents, err := services.GetDependentsForUser(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(dependents)
}

// CreateUser creates a new user with a hashed password and validated input
func CreateUser(c *fiber.Ctx) error {
	var input dto.CreateUserDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := val.Validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"validation_error": err.Error()})
	}
	hash, err := utils.HashPassword(input.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}
	user := models.User{
		FirstName:      input.FirstName,
		LastName:       input.LastName,
		Email:          input.Email,
		PhoneNumber:    input.PhoneNumber,
		PasswordHash:   hash,
		MembershipTier: input.MembershipTier,
		DateOfBirth:    &input.DateOfBirth,
	}
	if err := services.CreateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":              user.ID,
		"first_name":      user.FirstName,
		"last_name":       user.LastName,
		"email":           user.Email,
		"phone_number":    user.PhoneNumber,
		"membership_tier": user.MembershipTier,
		"date_of_birth":   user.DateOfBirth,
	})
}

// UpdateUser applies patch-style updates to allowed fields on a user (excluding password)
func UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	var input dto.UpdateUserDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := val.Validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"validation_error": err.Error()})
	}

	updates := make(map[string]interface{})
	if input.FirstName != nil {
		updates["first_name"] = *input.FirstName
	}
	if input.LastName != nil {
		updates["last_name"] = *input.LastName
	}
	if input.Email != nil {
		updates["email"] = *input.Email
	}
	if input.PhoneNumber != nil {
		updates["phone_number"] = *input.PhoneNumber
	}
	if input.MembershipTier != nil && *input.MembershipTier != "" {
		updates["membership_tier"] = *input.MembershipTier
	}
	if input.DateOfBirth != nil {
		updates["date_of_birth"] = *input.DateOfBirth
	}

	if err := services.UpdateUser(uint(id), updates); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
