package middleware

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

func Protected() fiber.Handler {
	supabaseURL := os.Getenv("SUPABASE_URL")
	if supabaseURL == "" {
		supabaseURL = "https://rhhqnmwnckgzhdyqsspb.supabase.co"
	}

	jwksURL := supabaseURL + "/auth/v1/.well-known/jwks.json"

	// Fetch Supabase's public key set once at startup
	keySet, err := jwk.Fetch(context.Background(), jwksURL)
	if err != nil {
		panic(fmt.Sprintf("Failed to fetch Supabase JWKS: %v", err))
	}

	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing or invalid authorization header",
			})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse and verify the JWT using Supabase's public keys (supports ES256 + HS256)
		token, err := jwt.Parse([]byte(tokenString), jwt.WithKeySet(keySet))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "Invalid token",
				"details": err.Error(),
			})
		}

		// Attach the Supabase User ID (sub) to the request context
		if sub, ok := token.Subject(); ok && sub != "" {
			c.Locals("user_id", sub)
		}

		// Also extract email and phone if present
		var email string
		if err := token.Get("email", &email); err == nil && email != "" {
			c.Locals("email", email)
		}

		var phone string
		if err := token.Get("phone", &phone); err == nil && phone != "" {
			c.Locals("phone", phone)
		}

		// Extract auth provider from app_metadata
		var appMetadata map[string]interface{}
		if err := token.Get("app_metadata", &appMetadata); err == nil {
			if provider, ok := appMetadata["provider"].(string); ok && provider != "" {
				c.Locals("auth_method", provider) // e.g. 'email', 'phone', 'google', 'facebook'
				if provider == "email" || provider == "phone" {
					c.Locals("provider", "manual")
				} else {
					c.Locals("provider", provider)
				}
			}
		}

		return c.Next()
	}
}
