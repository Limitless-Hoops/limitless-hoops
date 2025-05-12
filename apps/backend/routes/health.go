package routes

import "github.com/gofiber/fiber/v2"

func Health(router fiber.Router) {
	router.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "Success",
			"message": "Backend is healthy",
		})
	})
}
