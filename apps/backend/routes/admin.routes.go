package routes

import (
	"github.com/Limitless-Hoops/limitless-hoops/controllers"
	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(router fiber.Router) {
	group := router.Group("/admins")

	group.Get("/", controllers.GetAdmins)
	group.Get("/:id", controllers.GetAdminByID)
	group.Post("/", controllers.CreateAdmin)
	group.Patch("/:id", controllers.UpdateAdmin)
	
}
