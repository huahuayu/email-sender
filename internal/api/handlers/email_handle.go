package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/huahuayu/email-sender/internal/api/models"
	"github.com/huahuayu/email-sender/internal/services"
	"github.com/sirupsen/logrus"
)

type EmailHandler struct {
	emailService *services.EmailService
}

func NewEmailHandler(emailService *services.EmailService) *EmailHandler {
	return &EmailHandler{
		emailService: emailService,
	}
}

func (h *EmailHandler) HandleSendEmail(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("Received email send request from %s", r.RemoteAddr)

	if r.Method != http.MethodPost {
		logrus.Warnf("Invalid method %s from %s", r.Method, r.RemoteAddr)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logrus.Errorf("Failed to decode request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	logrus.Infof("Processing email request to %d recipients, subject: %s", len(req.To), req.Subject)

	// Validate required fields
	if len(req.To) == 0 {
		logrus.Warn("Request validation failed: no recipients specified")
		http.Error(w, "At least one recipient is required", http.StatusBadRequest)
		return
	}
	if req.Subject == "" {
		logrus.Warn("Request validation failed: no subject specified")
		http.Error(w, "Subject is required", http.StatusBadRequest)
		return
	}
	if req.Body == "" && req.HtmlBody == "" {
		logrus.Warn("Request validation failed: no body content specified")
		http.Error(w, "Either body or htmlBody is required", http.StatusBadRequest)
		return
	}

	// Validate email addresses
	for _, recipient := range req.To {
		if recipient.Email == "" {
			logrus.Warn("Invalid recipient email address found in request")
			http.Error(w, "Invalid recipient email address", http.StatusBadRequest)
			return
		}
	}

	if err := h.emailService.SendEmail(req); err != nil {
		logrus.Errorf("Failed to send email: %v", err)
		http.Error(w, "Failed to send email: "+err.Error(), http.StatusInternalServerError)
		return
	}

	logrus.Info("Email sent successfully")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Email sent successfully"})
}
