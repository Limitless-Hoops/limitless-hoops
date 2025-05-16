package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	Health(v1)
	UserRoutes(v1)
	AuthRoutes(v1)
	AdminRoutes(v1)
	DependentRoutes(v1)
}
