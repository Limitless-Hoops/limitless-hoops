package routes

import (
	"github.com/Limitless-Hoops/limitless-hoops/controllers"
	"github.com/Limitless-Hoops/limitless-hoops/middleware"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(router fiber.Router) {
	group := router.Group("/auth")

	group.Post("/login", controllers.Login)                                         // POST /auth/login → log in and get JWT
	group.Get("/me", middleware.JWTProtected(), controllers.Me)                     // GET /auth/me → current user info (auth required)
	group.Patch("/password", middleware.JWTProtected(), controllers.UpdatePassword) // PATCH /auth/password → update password (auth required)
}
