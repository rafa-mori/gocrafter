# GoCrafter Makefile
# Build automation for Go project scaffolding tool

# Variables
BINARY_NAME := gocrafter
VERSION := $(shell git describe --tags --always --dirty)
BUILD_TIME := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
COMMIT := $(shell git rev-parse --short HEAD)
LDFLAGS := -ldflags "-s -w -X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME) -X main.Commit=$(COMMIT)"

# Colors for output
COLOR_GREEN := \033[32m
COLOR_YELLOW := \033[33m
COLOR_RED := \033[31m
COLOR_BLUE := \033[34m
COLOR_RESET := \033[0m

# Logging functions
define log
	@printf "%b[%s]%b %s\n" "$(COLOR_BLUE)" "$(1)" "$(COLOR_RESET)" "$(2)"
endef

define log_success
	@printf "%b[SUCCESS]%b %s\n" "$(COLOR_GREEN)" "$(COLOR_RESET)" "$(1)"
endef

define log_error
	@printf "%b[ERROR]%b %s\n" "$(COLOR_RED)" "$(COLOR_RESET)" "$(1)"
endef

# Default target
.DEFAULT_GOAL := help

# Build the binary
.PHONY: build
build:
	$(call log,BUILD,Building $(BINARY_NAME))
	@go build $(LDFLAGS) -o $(BINARY_NAME) ./cmd
	$(call log_success,Binary built successfully: $(BINARY_NAME))

# Install the binary to GOPATH/bin
.PHONY: install
install:
	$(call log,INSTALL,Installing $(BINARY_NAME))
	@go install $(LDFLAGS) ./cmd
	$(call log_success,$(BINARY_NAME) installed successfully)

# Run tests
.PHONY: test
test:
	$(call log,TEST,Running tests)
	@go test -v -race -coverprofile=coverage.out ./...
	$(call log_success,Tests completed successfully)

# Run tests with coverage report
.PHONY: test-coverage
test-coverage: test
	$(call log,COVERAGE,Generating coverage report)
	@go tool cover -html=coverage.out -o coverage.html
	$(call log_success,Coverage report generated: coverage.html)

# Run linting
.PHONY: lint
lint:
	$(call log,LINT,Running linters)
	@golangci-lint run
	$(call log_success,Linting completed successfully)

# Format code
.PHONY: fmt
fmt:
	$(call log,FORMAT,Formatting code)
	@gofmt -s -w .
	@goimports -w .
	$(call log_success,Code formatted successfully)

# Clean build artifacts
.PHONY: clean
clean:
	$(call log,CLEAN,Cleaning build artifacts)
	@rm -f $(BINARY_NAME)
	@rm -f coverage.out coverage.html
	@rm -rf dist/
	$(call log_success,Cleanup completed)

# Run the application
.PHONY: run
run: build
	$(call log,RUN,Running $(BINARY_NAME))
	@./$(BINARY_NAME) $(ARGS)

# Development setup
.PHONY: setup
setup:
	$(call log,SETUP,Setting up development environment)
	@go mod download
	@go install golang.org/x/tools/cmd/goimports@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$(call log_success,Development environment setup completed)

# Build for multiple platforms
.PHONY: build-all
build-all: clean
	$(call log,BUILD,Building for multiple platforms)
	@mkdir -p dist
	@GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-linux-amd64 ./cmd
	@GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-linux-arm64 ./cmd
	@GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-amd64 ./cmd
	@GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-arm64 ./cmd
	@GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-windows-amd64.exe ./cmd
	$(call log_success,Multi-platform builds completed)

# Create release packages
.PHONY: package
package: build-all
	$(call log,PACKAGE,Creating release packages)
	@cd dist && for file in *; do \
		if [[ "$$file" == *"windows"* ]]; then \
			zip -r "$$file.zip" "$$file"; \
		else \
			tar -czf "$$file.tar.gz" "$$file"; \
		fi; \
	done
	$(call log_success,Release packages created)

# Run integration tests
.PHONY: test-integration
test-integration: build
	$(call log,INTEGRATION,Running integration tests)
	@./$(BINARY_NAME) list
	@./$(BINARY_NAME) info api-rest
	@./$(BINARY_NAME) new test-project --template api-rest --quick
	@cd test-project && go mod tidy && go build
	@rm -rf test-project
	$(call log_success,Integration tests completed)

# Check dependencies for security vulnerabilities
.PHONY: security
security:
	$(call log,SECURITY,Checking for security vulnerabilities)
	@go list -json -deps ./... | nancy sleuth
	$(call log_success,Security check completed)

# Generate documentation
.PHONY: docs
docs:
	$(call log,DOCS,Generating documentation)
	@go run cmd/main.go docs > docs/cli-reference.md
	$(call log_success,Documentation generated)

# Docker build
.PHONY: docker-build
docker-build:
	$(call log,DOCKER,Building Docker image)
	@docker build -t rafamori/gocrafter:latest .
	@docker build -t rafamori/gocrafter:$(VERSION) .
	$(call log_success,Docker image built successfully)

# Docker run
.PHONY: docker-run
docker-run:
	$(call log,DOCKER,Running Docker container)
	@docker run --rm -it rafamori/gocrafter:latest $(ARGS)

# Release preparation
.PHONY: pre-release
pre-release: clean fmt lint test test-integration build-all package
	$(call log,RELEASE,Preparing release)
	$(call log_success,Release preparation completed)

# Show help
.PHONY: help
help:
	@echo ""
	@echo "$(COLOR_BLUE)GoCrafter - Go Project Scaffolding Tool$(COLOR_RESET)"
	@echo ""
	@echo "$(COLOR_YELLOW)Development Commands:$(COLOR_RESET)"
	@echo "  setup           Setup development environment"
	@echo "  build           Build the binary"
	@echo "  install         Install the binary to GOPATH/bin"
	@echo "  run             Build and run the application (use ARGS='...' for arguments)"
	@echo "  clean           Clean build artifacts"
	@echo ""
	@echo "$(COLOR_YELLOW)Code Quality:$(COLOR_RESET)"
	@echo "  fmt             Format code"
	@echo "  lint            Run linters"
	@echo "  test            Run tests"
	@echo "  test-coverage   Run tests with coverage report"
	@echo "  test-integration Run integration tests"
	@echo "  security        Check for security vulnerabilities"
	@echo ""
	@echo "$(COLOR_YELLOW)Release:$(COLOR_RESET)"
	@echo "  build-all       Build for multiple platforms"
	@echo "  package         Create release packages"
	@echo "  pre-release     Complete release preparation"
	@echo ""
	@echo "$(COLOR_YELLOW)Docker:$(COLOR_RESET)"
	@echo "  docker-build    Build Docker image"
	@echo "  docker-run      Run Docker container (use ARGS='...' for arguments)"
	@echo ""
	@echo "$(COLOR_YELLOW)Documentation:$(COLOR_RESET)"
	@echo "  docs            Generate documentation"
	@echo ""
	@echo "$(COLOR_YELLOW)Examples:$(COLOR_RESET)"
	@echo "  make run ARGS='list'"
	@echo "  make run ARGS='new my-api --template api-rest'"
	@echo "  make docker-run ARGS='--help'"
	@echo ""


