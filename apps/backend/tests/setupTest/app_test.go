package testsetup

import (
	"github.com/Limitless-Hoops/limitless-hoops/database"
	"github.com/Limitless-Hoops/limitless-hoops/routes"
	"github.com/gofiber/fiber/v2"
)

// NewTestApp sets up a Fiber app with the test DB and registered routes.
// This function is used in integration tests to simulate real app behavior.
func NewTestApp() *fiber.App {
	app := fiber.New()

	database.DB = ConnectTestDB()

	routes.Setup(app)

	return app
}
