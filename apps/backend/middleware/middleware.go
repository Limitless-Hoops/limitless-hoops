package middleware

import (
	"github.com/Limitless-Hoops/limitless-hoops/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Setup TODO: Set Correct Origin and Allow Credentials to True when Frontend is built
func Setup(app *fiber.App, conf *config.Config) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     conf.FrontendUrl,
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders:     "Content-Type, Authorization",
		AllowCredentials: true,
	}))
}
