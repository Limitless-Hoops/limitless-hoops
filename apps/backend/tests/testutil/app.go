package testutil

import (
	"github.com/Limitless-Hoops/limitless-hoops/database"
	"github.com/Limitless-Hoops/limitless-hoops/routes"
	"github.com/gofiber/fiber/v2"
)

func NewTestApp() *fiber.App {
	app := fiber.New()
	database.DB = ConnectTestDB()
	routes.Setup(app)
	return app
}
