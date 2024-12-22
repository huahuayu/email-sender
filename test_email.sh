#!/bin/bash

# Test 1: Simple plain text email
echo "Test 1: Sending plain text email..."
curl -X POST http://localhost:8080/send \
    -H "Content-Type: application/json" \
    -d '{
    "from": {
        "email": "sender@example.com",
        "name": "John Doe"
    },
    "to": [{"email": "recipient@example.com", "name": "Jane Doe"}],
    "subject": "Test Plain Text Email",
    "body": "This is a test email sent via the email-sender service."
}'

echo -e "\n\n"

# Test 2: HTML email with multiple recipients and CC
echo "Test 2: Sending HTML email with CC..."
curl -X POST http://localhost:8080/send \
    -H "Content-Type: application/json" \
    -d '{
    "to": [
        {"email": "recipient1@example.com", "name": "John Smith"},
        {"email": "recipient2@example.com", "name": "Jane Smith"}
    ],
    "cc": [
        {"email": "cc-person@example.com", "name": "CC Recipient"}
    ],
    "subject": "HTML Test Email with CC",
    "body": "This is the plain text version.",
    "htmlBody": "<h1>Hello</h1><p>This is a <b>HTML</b> test email with CC recipients.</p>"
}'

echo -e "\n\n"

# Test 3: Email with custom sender, CC, BCC, and base64 attachment
echo "Test 3: Sending email with custom sender, CC, BCC, and attachment..."
curl -X POST http://localhost:8080/send \
    -H "Content-Type: application/json" \
    -d '{
    "to": [{"email": "primary@example.com", "name": "Primary Recipient"}],
    "cc": [{"email": "cc-recipient@example.com", "name": "CC Person"}],
    "bcc": [{"email": "bcc-recipient@example.com"}],
    "subject": "Test Email with Everything",
    "body": "This is a test email with all features.",
    "htmlBody": "<h1>Full Feature Test</h1><p>This email includes:</p><ul><li>Custom sender</li><li>CC recipient</li><li>BCC recipient</li><li>Attachment</li></ul>",
    "from": {
        "email": "custom@example.com",
        "name": "Custom Sender"
    },
    "attachments": [{
        "filename": "test.txt",
        "content": "SGVsbG8sIHRoaXMgaXMgYSB0ZXN0IGF0dGFjaG1lbnQu",
        "contentType": "text/plain"
    }]
}'

echo -e "\n\n"
