package controllers

import (
	"github.com/Limitless-Hoops/limitless-hoops/dto"
	"github.com/Limitless-Hoops/limitless-hoops/services"
	val "github.com/Limitless-Hoops/limitless-hoops/validator"
	"github.com/gofiber/fiber/v2"
	"log"
)

// Login POST /auth/login → authenticate and return JWT
func Login(c *fiber.Ctx) error {
	var input dto.LoginDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := val.Validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"validation_error": err.Error()})
	}

	token, err := services.Login(input)
	if err != nil {
		log.Println("❌ Login failed:", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}

	return c.JSON(fiber.Map{"token": token})
}

// Me GET /auth/me → return current role-based profile info
func Me(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	role := c.Locals("role").(string)
	log.Println("🔍 Role in token:", role)

	profile, err := services.GetProfileByRole(userID, role)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(profile)
}

// UpdatePassword PATCH /auth/password → change password
func UpdatePassword(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	role := c.Locals("role").(string)

	var input dto.UpdatePasswordDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := val.Validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"validation_error": err.Error()})
	}

	if err := services.UpdatePasswordByRole(userID, role, input.OldPassword, input.NewPassword); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
