package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type QRCodeConfig struct {
	Bucket string
}

type Config struct {
	Port   string
	DB     DBConfig
	URL    Url
	QRCode QRCodeConfig
}

type Url struct {
	BaseUrl            string
	FileServiceBaseUrl string
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
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "127.0.0.1"),
			Port:     getEnv("DB_PORT", "3306"),
			User:     getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", "shortlink"),
		},
		URL: Url{
			BaseUrl:            getEnv("BASE_URL", "127.0.0.1"),
			FileServiceBaseUrl: getEnv("FILE_SERVICE_BASE_URL", "127.0.0.1"),
		},
		QRCode: QRCodeConfig{
			Bucket: getEnv("QRCODE_BUCKET", "shortlink-qrcodes"),
		},
	}

	return AppConfig
}

func getEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
