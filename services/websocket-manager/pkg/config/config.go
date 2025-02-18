package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds application settings
type Config struct {
	Port string
}

// LoadConfig loads environment variables
func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using defaults")
	}

	return Config{
		Port: getEnv("PORT", "8080"),
	}
}

// getEnv retrieves env variable or default
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
