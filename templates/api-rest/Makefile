# {{.ProjectName}} Makefile

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=gofmt

# Build parameters
BINARY_NAME={{.PackageName}}
BINARY_PATH=./bin/$(BINARY_NAME)
MAIN_PATH=./cmd/main.go

# Docker parameters
DOCKER_IMAGE={{.PackageName}}
DOCKER_TAG=latest

.PHONY: all build clean test coverage deps fmt lint run docker-build docker-run help

# Default target
all: deps fmt lint test build

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	$(GOBUILD) -o $(BINARY_PATH) $(MAIN_PATH)

# Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -f $(BINARY_PATH)

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Run tests with coverage
coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

# Format code
fmt:
	@echo "Formatting code..."
	$(GOFMT) -s -w .

# Lint code
lint:
	@echo "Linting code..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Install it with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

# Run the application
run:
	@echo "Running $(BINARY_NAME)..."
	$(GOCMD) run $(MAIN_PATH)

# Run with live reload
dev:
	@echo "Running with live reload..."
	@if command -v air >/dev/null 2>&1; then \
		air; \
	else \
		echo "Air not installed. Install it with: go install github.com/cosmtrek/air@latest"; \
		$(GOCMD) run $(MAIN_PATH); \
	fi

{{- if .HasDocker}}
# Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

# Run Docker container
docker-run: docker-build
	@echo "Running Docker container..."
	docker run --rm -p 8080:8080 --env-file .env $(DOCKER_IMAGE):$(DOCKER_TAG)

# Run with docker-compose
docker-compose-up:
	@echo "Starting services with docker-compose..."
	docker-compose up --build

# Stop docker-compose services
docker-compose-down:
	@echo "Stopping services..."
	docker-compose down
{{- end}}

# Install development tools
install-tools:
	@echo "Installing development tools..."
	$(GOCMD) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$(GOCMD) install github.com/cosmtrek/air@latest
	{{- if hasFeature "API Documentation"}}
	$(GOCMD) install github.com/swaggo/swag/cmd/swag@latest
	{{- end}}

{{- if hasFeature "API Documentation"}}
# Generate Swagger documentation
swagger:
	@echo "Generating Swagger documentation..."
	swag init -g cmd/main.go -o docs/
{{- end}}

# Database migrations (if using a database)
{{- if .DatabaseType}}
migrate-up:
	@echo "Running database migrations..."
	# Add your migration command here

migrate-down:
	@echo "Rolling back database migrations..."
	# Add your rollback command here
{{- end}}

# Show help
help:
	@echo "Available targets:"
	@echo "  build          - Build the binary"
	@echo "  clean          - Clean build artifacts"
	@echo "  test           - Run tests"
	@echo "  coverage       - Run tests with coverage"
	@echo "  deps           - Download dependencies"
	@echo "  fmt            - Format code"
	@echo "  lint           - Lint code"
	@echo "  run            - Run the application"
	@echo "  dev            - Run with live reload"
	{{- if .HasDocker}}
	@echo "  docker-build   - Build Docker image"
	@echo "  docker-run     - Run Docker container"
	@echo "  docker-compose-up   - Start with docker-compose"
	@echo "  docker-compose-down - Stop docker-compose services"
	{{- end}}
	@echo "  install-tools  - Install development tools"
	{{- if hasFeature "API Documentation"}}
	@echo "  swagger        - Generate Swagger documentation"
	{{- end}}
	{{- if .DatabaseType}}
	@echo "  migrate-up     - Run database migrations"
	@echo "  migrate-down   - Rollback database migrations"
	{{- end}}
	@echo "  help           - Show this help message"
