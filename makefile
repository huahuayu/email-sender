.PHONY: build run test clean docker-up docker-down docker-logs help

# Go parameters
BINARY_NAME=email-sender
MAIN_PATH=cmd/email-sender/main.go

# Build parameters
BUILD_DIR=build
GOARCH=amd64

help:
	@echo "Available commands:"
	@echo "  make build       	- Build the application"
	@echo "  make run        	- Run the application"
	@echo "  make test       	- Run tests"
	@echo "  make clean      	- Clean build artifacts"
	@echo "  make docker-up   	- Start services with Docker Compose"
	@echo "  make docker-down 	- Stop Docker Compose services"
	@echo "  make docker-logs 	- View Docker Compose logs"

build:
	@echo "Building..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

run:
	@go run $(MAIN_PATH)

test:
	@echo "Running tests..."
	@go test -v ./...

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@go clean
	@echo "Clean complete"

docker-up:
	@echo "Starting services..."
	@docker compose up --build -d

docker-down:
	@echo "Stopping services..."
	@docker compose down

docker-logs:
	@docker compose logs -f

# Default target
.DEFAULT_GOAL := help