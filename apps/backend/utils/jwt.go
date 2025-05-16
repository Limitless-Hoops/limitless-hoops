package utils

import (
	"github.com/Limitless-Hoops/limitless-hoops/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"` // "user", "dependent", "admin"
	jwt.RegisteredClaims
}

func GenerateJWT(userID uint, role string) (string, error) {
	expiration := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.JwtKey))
}
