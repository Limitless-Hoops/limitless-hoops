package controllers

import (
	"strconv"
	"time"

	"github.com/Limitless-Hoops/limitless-hoops/models"
	"github.com/Limitless-Hoops/limitless-hoops/services"
	"github.com/Limitless-Hoops/limitless-hoops/utils"
	"github.com/gofiber/fiber/v2"
)

// UserInput DTO for incoming user data
type UserInput struct {
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	PhoneNumber    string    `json:"phone_number"`
	Password       string    `json:"password"` // Accept plain text password here
	MembershipTier string    `json:"membership_tier"`
	DateOfBirth    time.Time `json:"date_of_birth"`
}

// UserUpdateInput DTO for partial update (PATCH)
type UserUpdateInput struct {
	FirstName      *string    `json:"first_name,omitempty"`
	LastName       *string    `json:"last_name,omitempty"`
	Email          *string    `json:"email,omitempty"`
	PhoneNumber    *string    `json:"phone_number,omitempty"`
	Password       *string    `json:"password,omitempty"` // Accept plain text password here
	MembershipTier *string    `json:"membership_tier,omitempty"`
	DateOfBirth    *time.Time `json:"date_of_birth,omitempty"`
}

// GetUsers GET /users
func GetUsers(c *fiber.Ctx) error {
	users, err := services.GetAllUsersWithDependentCount()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

// GetUserByID GET /users/:id
func GetUserByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	user, err := services.GetUserByIDWithDependentsAndContacts(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}

// GetUserDependents GET /users/:id/dependents
func GetUserDependents(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	dependents, err := services.GetDependentsForUser(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(dependents)
}

// CreateUser POST /users
func CreateUser(c *fiber.Ctx) error {
	var input UserInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
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

	return c.Status(fiber.StatusCreated).JSON(user)
}

// UpdateUser PATCH /users/:id
func UpdateUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var input UserUpdateInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
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
	if input.Password != nil && *input.Password != "" {
		hash, err := utils.HashPassword(*input.Password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
		}
		updates["password_hash"] = hash
	}

	if err := services.UpdateUser(uint(id), updates); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
