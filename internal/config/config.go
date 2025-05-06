package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

var AppConfig *Config

func Load() *Config {
	// Load .env file if exists
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, relying on environment variables")
	}

	AppConfig = &Config{
		Port: getEnv("PORT", "8080"),
	}

	return AppConfig
}

func getEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
