package server

import (
	"context"
	"github.com/Limitless-Hoops/limitless-hoops/config"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start(app *fiber.App, conf *config.Config) {
	go func() {
		if err := app.Listen(":" + conf.BackendPort); err != nil {
			log.Panicf("Failed to start server: %v", err)
		}
	}()
	log.Println("Server running on port " + conf.BackendPort)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.Shutdown(); err != nil {
		log.Panicf("Server shutdown failed: %v", err)
	}

	log.Println("Server gracefully stopped")
}
