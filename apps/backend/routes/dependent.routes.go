package routes

import (
	"github.com/Limitless-Hoops/limitless-hoops/controllers"
	"github.com/gofiber/fiber/v2"
)

func DependentRoutes(router fiber.Router) {
	group := router.Group("/dependents")

	group.Get("/", controllers.GetDependents)        // GET /dependents
	group.Get("/:id", controllers.GetDependentByID)  // GET /dependents/:id
	group.Post("/", controllers.CreateDependent)     // POST /dependents
	group.Patch("/:id", controllers.UpdateDependent) // PATCH /dependents/:id
}
