package controllers

import (
	"strconv"

	"github.com/Limitless-Hoops/limitless-hoops/dto"
	"github.com/Limitless-Hoops/limitless-hoops/models"
	"github.com/Limitless-Hoops/limitless-hoops/services"
	"github.com/Limitless-Hoops/limitless-hoops/utils"
	"github.com/Limitless-Hoops/limitless-hoops/validator"
	"github.com/gofiber/fiber/v2"
)

// GetAdmins GET /admins → list all admins with dependent counts
func GetAdmins(c *fiber.Ctx) error {
	admins, err := services.GetAllAdminsWithDependentCount()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(admins)
}

// GetAdminByID GET /admins/:id → return basic admin info
func GetAdminByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid admin ID"})
	}

	admin, err := services.GetAdminByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Admin not found"})
	}

	return c.JSON(admin)
}

// CreateAdmin POST /admins
func CreateAdmin(c *fiber.Ctx) error {
	var input dto.CreateAdminDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := validator.Validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"validation_error": err.Error()})
	}

	hash, err := utils.HashPassword(input.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	admin := models.Admin{
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Email:        input.Email,
		PhoneNumber:  input.PhoneNumber,
		PasswordHash: hash,
		Role:         input.Role,
		DateOfBirth:  &input.DateOfBirth,
	}

	if err := services.CreateAdmin(&admin); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":            admin.ID,
		"first_name":    admin.FirstName,
		"last_name":     admin.LastName,
		"email":         admin.Email,
		"phone_number":  admin.PhoneNumber,
		"role":          admin.Role,
		"date_of_birth": admin.DateOfBirth,
	})
}

// UpdateAdmin PATCH /admins/:id
func UpdateAdmin(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid admin ID"})
	}

	var input dto.UpdateAdminDTO
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
	if input.Role != nil {
		updates["role"] = *input.Role
	}
	if input.DateOfBirth != nil {
		updates["date_of_birth"] = *input.DateOfBirth
	}

	if err := services.UpdateAdmin(uint(id), updates); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
