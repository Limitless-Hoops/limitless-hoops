package middleware

import (
	"github.com/Limitless-Hoops/limitless-hoops/config"
	"github.com/Limitless-Hoops/limitless-hoops/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Get("Authorization")
		if tokenStr == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or invalid token"})
		}

		// Strip "Bearer " if present
		if len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
			tokenStr = tokenStr[7:]
		}

		log.Println("🔐 JWT secret length:", len(config.AppConfig.JwtKey))
		claims := &utils.Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.AppConfig.JwtKey), nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		c.Locals("user_id", claims.UserID)
		c.Locals("role", claims.Role)
		return c.Next()
	}
}
