# Email Sender

Email Sender is a Go service that allows you sent email by api.

Gmail SMTP is tested, but it should work with other SMTP servers too.

## Quick Start

### Clone the repository
```sh
git clone https://github.com/huahuayu/email-sender.git
cd email-sender
```

### Config

Config through environment variables or `.env` file.

```env
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
FROM_EMAIL=your-email@gmail.com
APP_PASSWORD=your-google-app-password
SERVER_PORT=8080 # optional
LOG_LEVEL=info # optional
```

### Build & Run

Run in local 
```bash
make run
``` 

Or run in Docker container 
```bash
make docker-run
```

## Endpoint

`POST /send` - Send an email

**Request Body**
```json
{
  "from": {"email": "your-email@gmail.com", "name": "Your Name"},
  "to": [{"email": "recipient@example.com", "name": "Recipient Name"}],
  "cc": [{"email": "cc@example.com", "name": "CC Name"}],
  "bcc": [{"email": "bcc@example.com", "name": "BCC Name"}],
  "subject": "Email Subject",
  "body": "Plain text body",
  "htmlBody": "<p>HTML body</p>",
  "attachments": [
    {
      "filename": "example.txt",
      "content": "SGVsbG8gd29ybGQ=",
      "contentType": "text/plain"
    }
  ]
}
```

`from`, `cc`, `bcc`, `htmlBody`, `attachments` are optional fields.

`attachments.content` should be base64 encoded.

**Example Request**

```sh
curl -X POST http://localhost:8080/send \
    -H "Content-Type: application/json" \
    -d '{
    "to": [{"email": ""recipient@example.com", "name": "Alice"}],
    "subject": "Test Plain Text Email",
    "body": "This is a test email sent via the email-sender service."
}'
```
**Response**

```json
{
  "message": "Email sent successfully"
}
```