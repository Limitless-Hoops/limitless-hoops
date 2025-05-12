package routes

import (
	"github.com/Limitless-Hoops/limitless-hoops/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router) {
	group := router.Group("/users")

	group.Get("/", controllers.GetUsers)                        // GET /users
	group.Get("/:id", controllers.GetUserByID)                  // GET /users/:id
	group.Get("/:id/dependents", controllers.GetUserDependents) // GET /users/:id/dependents
	group.Post("/", controllers.CreateUser)                     // POST /users
	group.Patch("/:id", controllers.UpdateUser)                 // PATCH /users/:id
}
