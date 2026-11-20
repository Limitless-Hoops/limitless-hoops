package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v76"

	"github.com/limitless-hoops/api/database"
	"github.com/limitless-hoops/api/handlers"
	"github.com/limitless-hoops/api/middleware"
)

func main() {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	database.ConnectDB()

	// Initialize Stripe
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	app := fiber.New(fiber.Config{
		AppName: "Limitless Hoops API",
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Limitless Hoops API is healthy!")
	})

	// Stripe Webhooks (Unprotected, uses signature verification)
	app.Post("/api/webhooks/stripe", handlers.HandleStripeWebhook)

	// API Routes (Accounts/Households)
	api := app.Group("/api", middleware.Protected())
	api.Post("/households", handlers.CreateHousehold)
	api.Get("/households/me", handlers.GetMyHousehold)
	api.Get("/households/:id", handlers.GetHousehold)
	api.Post("/households/:id/players", handlers.AddPlayer)
	api.Post("/players/:id/contacts", handlers.AddPlayerContact)

	// Contact management — emails
	api.Put("/users/me", handlers.UpdateMyPersonalInfo)
	api.Get("/users/me/emails", handlers.GetMyEmails)
	api.Post("/users/me/emails", handlers.AddEmail)
	api.Delete("/users/me/emails/:id", handlers.DeleteEmail)
	api.Put("/users/me/emails/:id", handlers.UpdateEmail)
	api.Put("/users/me/emails/:id/set-primary", handlers.SetPrimaryEmail)
	api.Post("/users/me/emails/:id/verify", handlers.VerifyEmail)
	api.Post("/users/me/emails/:id/resend", handlers.ResendEmailVerification)

	// Contact management — phones
	api.Get("/users/me/phones", handlers.GetMyPhones)
	api.Post("/users/me/phones", handlers.AddPhone)
	api.Delete("/users/me/phones/:id", handlers.DeletePhone)
	api.Put("/users/me/phones/:id", handlers.UpdatePhone)
	api.Put("/users/me/phones/:id/set-primary", handlers.SetPrimaryPhone)
	api.Post("/users/me/phones/:id/verify", handlers.VerifyPhone)
	api.Post("/users/me/phones/:id/resend", handlers.ResendPhoneVerification)

	// Notification preferences
	api.Get("/users/me/notifications", handlers.GetMyNotifications)
	api.Put("/users/me/notifications/:id", handlers.UpdateNotificationPreference)

	// Provider management
	api.Post("/users/me/providers/:provider/disconnect", handlers.DisconnectProvider)

	// Stripe Checkout
	api.Post("/checkout", handlers.CreateCheckoutSession)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Starting server on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
