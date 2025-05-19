package controllers

import (
	"strconv"

	"github.com/Limitless-Hoops/limitless-hoops/dto"
	"github.com/Limitless-Hoops/limitless-hoops/models"
	"github.com/Limitless-Hoops/limitless-hoops/services"
	"github.com/Limitless-Hoops/limitless-hoops/utilities"
	"github.com/Limitless-Hoops/limitless-hoops/validator"
	"github.com/gofiber/fiber/v2"
)

// GetDependents GET /dependents → list all dependents with guardian counts
func GetDependents(c *fiber.Ctx) error {
	dependents, err := services.GetAllDependentsWithGuardianCount()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(dependents)
}

// GetDependentByID GET /dependents/:id → basic dependent info
func GetDependentByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid dependent ID"})
	}

	dependent, err := services.GetDependentByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Dependent not found"})
	}

	return c.JSON(dependent)
}

// CreateDependent POST /dependents
func CreateDependent(c *fiber.Ctx) error {
	var input dto.CreateDependentDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := validator.Validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"validation_error": err.Error()})
	}

	hash, err := utilities.HashPassword(input.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	dependent := models.Dependent{
		FirstName:      input.FirstName,
		LastName:       input.LastName,
		Email:          &input.Email,
		PhoneNumber:    &input.PhoneNumber,
		PasswordHash:   hash,
		MembershipTier: input.MembershipTier,
		DateOfBirth:    &input.DateOfBirth,
		AdminID:        input.AdminID,
	}

	if err := services.CreateDependent(&dependent); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":              dependent.ID,
		"first_name":      dependent.FirstName,
		"last_name":       dependent.LastName,
		"email":           dependent.Email,
		"phone_number":    dependent.PhoneNumber,
		"membership_tier": dependent.MembershipTier,
		"date_of_birth":   dependent.DateOfBirth,
		"admin_id":        dependent.AdminID,
	})
}

// UpdateDependent PATCH /dependents/:id
func UpdateDependent(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid dependent ID"})
	}

	var input dto.UpdateDependentDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := validator.Validate.Struct(input); err != nil {
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
	if input.MembershipTier != nil {
		updates["membership_tier"] = *input.MembershipTier
	}
	if input.DateOfBirth != nil {
		updates["date_of_birth"] = *input.DateOfBirth
	}
	if input.AdminID != nil {
		updates["admin_id"] = *input.AdminID
	}

	if err := services.UpdateDependent(uint(id), updates); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
