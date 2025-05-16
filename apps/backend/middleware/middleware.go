package middleware

import (
	"github.com/Limitless-Hoops/limitless-hoops/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Setup(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     config.AppConfig.FrontendUrl,
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders:     "Content-Type, Authorization",
		AllowCredentials: true,
	}))
}
