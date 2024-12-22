package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	ServerPort  string
	SMTPHost    string
	SMTPPort    string
	FromEmail   string
	AppPassword string
	LogLevel    string
}

func LoadConfig() *Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		logrus.Warnf("Error loading .env file: %v", err)
	}

	cfg := &Config{
		ServerPort:  getEnvOrDefault("SERVER_PORT", "8080"),
		SMTPHost:    getEnvOrDefault("SMTP_HOST", "smtp.gmail.com"),
		SMTPPort:    getEnvOrDefault("SMTP_PORT", "587"),
		FromEmail:   getEnvOrDefault("FROM_EMAIL", ""),
		AppPassword: getEnvOrDefault("APP_PASSWORD", ""),
		LogLevel:    getEnvOrDefault("LOG_LEVEL", "info"),
	}

	// Configure logrus
	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	return cfg
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
