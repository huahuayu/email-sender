package main

import (
	"fmt"
	"net/http"

	"github.com/huahuayu/email-sender/internal/api/handlers"
	"github.com/huahuayu/email-sender/internal/config"
	"github.com/huahuayu/email-sender/internal/services"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.LoadConfig()

	if err := validateConfig(cfg); err != nil {
		logrus.Fatalf("Configuration error: %v", err)
	}

	// Initialize services
	emailService := services.NewEmailService(cfg)

	// Initialize handlers
	emailHandler := handlers.NewEmailHandler(emailService)

	// Setup routes
	http.HandleFunc("/send", emailHandler.HandleSendEmail)

	// Start server
	logrus.Infof("Server starting on :%s", cfg.ServerPort)
	logrus.Infof("Email configuration: SMTP_HOST=%s, SMTP_PORT=%s, FROM_EMAIL=%s",
		cfg.SMTPHost,
		cfg.SMTPPort,
		cfg.FromEmail)

	if err := http.ListenAndServe(":"+cfg.ServerPort, nil); err != nil {
		logrus.Fatalf("Server failed to start: %v", err)
	}
}

func validateConfig(cfg *config.Config) error {
	if cfg.FromEmail == "" {
		return fmt.Errorf("FROM_EMAIL is required")
	}
	if cfg.AppPassword == "" {
		return fmt.Errorf("APP_PASSWORD is required")
	}
	if cfg.SMTPHost == "" {
		return fmt.Errorf("SMTP_HOST is required")
	}
	if cfg.SMTPPort == "" {
		return fmt.Errorf("SMTP_PORT is required")
	}
	if cfg.ServerPort == "" {
		return fmt.Errorf("SERVER_PORT is required")
	}
	return nil
}
