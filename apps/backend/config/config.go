package config

import "os"

type Config struct {
	// DB Vars
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	// FRONTEND and BACKEND Vars
	FrontendUrl string
	BackendPort string

	// AUTH
	JwtKey string
}

func LoadConfig() *Config {
	return &Config{
		// Database Vars
		DbHost:     getEnv("DB_HOST", "localhost"),
		DbPort:     getEnv("DB_PORT", "5432"),
		DbUser:     getEnv("DB_USER", "postgres"),
		DbPassword: getEnv("DB_PASSWORD", "<PASSWORD>"),
		DbName:     getEnv("DB_NAME", "postgres"),

		// Frontend and Backend Vars
		FrontendUrl: getEnv("FRONTEND_URL", "https://limitlesshoops.dev"),
		BackendPort: getEnv("VITE_BACKEND_PORT", "80"),

		// Auth
		JwtKey: getEnv("JWT_SECRET", ""),
	}
}

func getEnv(key, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	return val
}
