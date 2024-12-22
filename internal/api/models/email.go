package models

type EmailRequest struct {
	// Recipients
	To  []Recipient `json:"to"`            // Multiple recipients
	Cc  []Recipient `json:"cc,omitempty"`  // Optional CC recipients
	Bcc []Recipient `json:"bcc,omitempty"` // Optional BCC recipients

	// Content
	Subject  string `json:"subject"`
	Body     string `json:"body"`               // Plain text content
	HtmlBody string `json:"htmlBody,omitempty"` // Optional HTML content

	// Recipient configuration
	From *Recipient `json:"from"`

	// Attachments (optional)
	Attachments []Attachment `json:"attachments,omitempty"`
}

type Recipient struct {
	Email string `json:"email"`          // Required email address
	Name  string `json:"name,omitempty"` // Optional display name
}

type Attachment struct {
	Filename    string `json:"filename"`
	Content     string `json:"content"` // Base64 encoded content
	ContentType string `json:"contentType"`
}
