package services

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"mime/multipart"
	"net/smtp"
	"strings"

	"github.com/huahuayu/email-sender/internal/api/models"
	"github.com/huahuayu/email-sender/internal/config"
	"github.com/sirupsen/logrus"
)

type EmailService struct {
	config *config.Config
}

func NewEmailService(cfg *config.Config) *EmailService {
	return &EmailService{
		config: cfg,
	}
}

func (s *EmailService) SendEmail(req models.EmailRequest) error {
	logrus.Debugf("Preparing to send email to %d recipients", len(req.To))

	auth := smtp.PlainAuth("", s.config.FromEmail, s.config.AppPassword, s.config.SMTPHost)

	// Determine sender
	from := s.config.FromEmail
	if req.From != nil {
		if req.From.Name != "" {
			from = fmt.Sprintf("%s <%s>", req.From.Name, req.From.Email)
		} else {
			from = req.From.Email
		}
	}
	logrus.Debugf("Sending from: %s", from)

	// Create email message
	var buf bytes.Buffer

	// Add headers
	buf.WriteString(fmt.Sprintf("From: %s\r\n", from))

	// Format recipients with names if provided
	toList := formatRecipientList(req.To)
	buf.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(toList, ", ")))

	if len(req.Cc) > 0 {
		ccList := formatRecipientList(req.Cc)
		buf.WriteString(fmt.Sprintf("Cc: %s\r\n", strings.Join(ccList, ", ")))
	}
	// Note: BCC recipients should not appear in the headers
	buf.WriteString(fmt.Sprintf("Subject: %s\r\n", req.Subject))

	// Handle multipart message if there's HTML content or attachments
	if req.HtmlBody != "" || len(req.Attachments) > 0 {
		writer := multipart.NewWriter(&buf)
		buf.WriteString("MIME-Version: 1.0\r\n")
		buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\r\n\r\n", writer.Boundary()))

		// Add text part
		if req.Body != "" {
			textPart, _ := writer.CreatePart(map[string][]string{
				"Content-Type": {"text/plain; charset=UTF-8"},
			})
			textPart.Write([]byte(req.Body))
		}

		// Add HTML part
		if req.HtmlBody != "" {
			htmlPart, _ := writer.CreatePart(map[string][]string{
				"Content-Type": {"text/html; charset=UTF-8"},
			})
			htmlPart.Write([]byte(req.HtmlBody))
		}

		// Add attachments
		for _, att := range req.Attachments {
			attPart, _ := writer.CreatePart(map[string][]string{
				"Content-Type":              {att.ContentType},
				"Content-Transfer-Encoding": {"base64"},
				"Content-Disposition":       {fmt.Sprintf(`attachment; filename="%s"`, att.Filename)},
			})
			decoded, _ := base64.StdEncoding.DecodeString(att.Content)
			attPart.Write(decoded)
		}

		writer.Close()
	} else {
		// Simple plain text email
		buf.WriteString("Content-Type: text/plain; charset=UTF-8\r\n\r\n")
		buf.WriteString(req.Body)
	}

	// Collect all recipient email addresses for SMTP envelope
	var allRecipients []string
	allRecipients = append(allRecipients, extractEmails(req.To)...)  // Primary recipients
	allRecipients = append(allRecipients, extractEmails(req.Cc)...)  // Carbon Copy recipients
	allRecipients = append(allRecipients, extractEmails(req.Bcc)...) // Blind Carbon Copy recipients

	// Before sending email
	logrus.Debugf("Attempting to connect to SMTP server %s:%s", s.config.SMTPHost, s.config.SMTPPort)

	// Send email
	err := smtp.SendMail(
		s.config.SMTPHost+":"+s.config.SMTPPort,
		auth,
		s.config.FromEmail,
		allRecipients,
		buf.Bytes(),
	)

	if err != nil {
		logrus.Errorf("SMTP error: %v", err)
		return err
	}

	return nil
}

// Helper function to format recipient list with names
func formatRecipientList(recipients []models.Recipient) []string {
	formatted := make([]string, len(recipients))
	for i, r := range recipients {
		if r.Name != "" {
			formatted[i] = fmt.Sprintf("%s <%s>", r.Name, r.Email)
		} else {
			formatted[i] = r.Email
		}
	}
	return formatted
}

// Helper function to extract email addresses from recipients
func extractEmails(recipients []models.Recipient) []string {
	emails := make([]string, len(recipients))
	for i, r := range recipients {
		emails[i] = r.Email
	}
	return emails
}
