package main

import (
	"github.com/Limitless-Hoops/limitless-hoops/database"
	"github.com/Limitless-Hoops/limitless-hoops/middleware"
	"github.com/Limitless-Hoops/limitless-hoops/routes"
	"github.com/Limitless-Hoops/limitless-hoops/server"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	middleware.Setup(app)

	database.Connect()
	defer database.Close()
	database.PopulateDB()

	routes.Setup(app)

	server.Start(app)
}
