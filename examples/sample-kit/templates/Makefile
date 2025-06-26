.PHONY: help build run test clean docker-build docker-run lint

# Default target
help: ## Show this help message
	@echo "{{project_name}} - {{description}}"
	@echo ""
	@echo "Available targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Build targets
build: ## Build the application
	@echo "Building {{project_name}}..."
	go build -o bin/{{project_name}} main.go

build-linux: ## Build for Linux
	GOOS=linux GOARCH=amd64 go build -o bin/{{project_name}}-linux-amd64 main.go

build-windows: ## Build for Windows
	GOOS=windows GOARCH=amd64 go build -o bin/{{project_name}}-windows-amd64.exe main.go

build-darwin: ## Build for macOS
	GOOS=darwin GOARCH=amd64 go build -o bin/{{project_name}}-darwin-amd64 main.go

# Run targets
run: ## Run the application
	@echo "Starting {{project_name}}..."
	go run main.go

dev: ## Run in development mode with auto-reload
	@echo "Starting {{project_name}} in development mode..."
	GIN_MODE=debug go run main.go

# Test targets
test: ## Run tests
	@echo "Running tests..."
	go test -v ./...

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Quality targets
lint: ## Run linter
	@echo "Running linter..."
	golangci-lint run

fmt: ## Format code
	@echo "Formatting code..."
	go fmt ./...

vet: ## Run go vet
	@echo "Running go vet..."
	go vet ./...

# Dependency targets
deps: ## Download dependencies
	@echo "Downloading dependencies..."
	go mod download

deps-update: ## Update dependencies
	@echo "Updating dependencies..."
	go mod tidy

# Docker targets
docker-build: ## Build Docker image
	@echo "Building Docker image..."
	docker build -t {{project_name}}:latest .

docker-run: ## Run Docker container
	@echo "Running Docker container..."
	docker run -p {{port}}:{{port}} --env-file .env {{project_name}}:latest

docker-stop: ## Stop Docker container
	@echo "Stopping Docker container..."
	docker stop $$(docker ps -q --filter ancestor={{project_name}}:latest)

# Clean targets
clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	rm -f coverage.out coverage.html

clean-all: clean ## Clean all artifacts including dependencies
	@echo "Cleaning all artifacts..."
	go clean -modcache

# Development targets
install-tools: ## Install development tools
	@echo "Installing development tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/air-verse/air@latest

watch: ## Run with auto-reload using air
	@echo "Starting {{project_name}} with auto-reload..."
	air
