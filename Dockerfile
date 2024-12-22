# Build stage
FROM golang:1.23-alpine AS builder

# Install necessary build tools
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/email-sender ./cmd/email-sender

# Final stage
FROM alpine:3.19

# Add ca-certificates for SMTP TLS connections
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/email-sender .

# Create a non-root user
RUN adduser -D appuser
USER appuser

# Expose the port the service runs on
EXPOSE 8080

# Run the application
CMD ["./email-sender"]