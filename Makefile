# GoCrafter Makefile
# Build automation for Go project scaffolding tool

# Variables
APP_NAME := $(shell echo $(basename $(CURDIR)) | tr '[:upper:]' '[:lower:]')
ROOT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
BINARY_NAME := $(ROOT_DIR)$(APP_NAME)
CMD_DIR := $(ROOT_DIR)cmd
VERSION := $(shell git describe --tags --always --dirty)

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


ARGUMENTS := $(MAKECMDGOALS)
INSTALL_SCRIPT=$(ROOT_DIR)support/install.sh
CMD_STR := $(strip $(firstword $(ARGUMENTS)))
ARGS := $(filter-out $(strip $(CMD_STR)), $(ARGUMENTS))

# Default target
.DEFAULT_GOAL := help

# Build the binary
.PHONY: build
build:
	$(call log,BUILD,Building $(BINARY_NAME))
	@bash $(INSTALL_SCRIPT) build $(ARGS)
	$(call log_success,Binary built successfully: $(BINARY_NAME))

# Install the binary to GOPATH/bin
# @go install $(LDFLAGS) ./cmd
.PHONY: install
install:
	$(call log,INSTALL,Installing $(BINARY_NAME))
	
	@bash $(INSTALL_SCRIPT) install $(ARGS)
	$(call log_success,$(BINARY_NAME) installed successfully)

# Run tests
.PHONY: test
test:
	$(call log,TEST,Running tests)
	@bash $(INSTALL_SCRIPT) test $(ARGS)
	$(call log_success,Tests completed successfully)

# Run tests with coverage report
# @go tool cover -html=coverage.out -o coverage.html
.PHONY: test-coverage
test-coverage: test
	$(call log,COVERAGE,Generating coverage report)
	@bash $(INSTALL_SCRIPT) test-coverage $(ARGS)
	$(call log_success,Coverage report generated: coverage.html)

# Run linting
# @golangci-lint run
.PHONY: lint
lint:
	$(call log,LINT,Running linters)
	@bash $(INSTALL_SCRIPT) lint $(ARGS)
	$(call log_success,Linting completed successfully)

# Format code
.PHONY: fmt
fmt:
	$(call log,FORMAT,Formatting code)
	@gofmt -s -w .
	@goimports -w .
	$(call log_success,Code formatted successfully)

# Clean build artifacts
# @rm -f $(BINARY_NAME)
# @rm -f coverage.out coverage.html
# @rm -rf dist/
.PHONY: clean
clean:
	$(call log,CLEAN,Cleaning build artifacts)
	@bash $(INSTALL_SCRIPT) clean $(ARGS)	
	$(call log_success,Cleanup completed)

# Run the application
# @./$(BINARY_NAME) $(ARGS)
.PHONY: run
run: build
	$(call log,RUN,Running $(BINARY_NAME))
	@bash $(INSTALL_SCRIPT) run $(ARGS)

# Run integration tests
# @./$(BINARY_NAME) list
# @./$(BINARY_NAME) info api-rest
# @./$(BINARY_NAME) new test-project --template api-rest --quick
# @cd test-project && go mod tidy && go build
# @rm -rf test-project
.PHONY: test-integration
test-integration: build
	$(call log,INTEGRATION,Running integration tests)
	@bash $(INSTALL_SCRIPT) test-integration $(ARGS)
	$(call log_success,Integration tests completed)

# Check dependencies for security vulnerabilities
# @go list -json -deps ./... | nancy sleuth
.PHONY: security
security:
	$(call log,SECURITY,Checking for security vulnerabilities)
	@bash $(INSTALL_SCRIPT) security $(ARGS)	
	$(call log_success,Security check completed)

# Generate documentation
# @go run cmd/main.go docs > docs/cli-reference.md
.PHONY: docs
docs:
	$(call log,DOCS,Generating documentation)
	@bash $(INSTALL_SCRIPT) docs $(ARGS)
	$(call log_success,Documentation generated)

# Docker build
# @docker build -t rafamori/gocrafter:latest .
# @docker build -t rafamori/gocrafter:$(VERSION) .
.PHONY: docker-build
docker-build:
	$(call log,DOCKER,Building Docker image)
	@docker build -t rafamori/gocrafter:latest -t rafamori/gocrafter:$(VERSION) .
	$(call log_success,Docker image built successfully)

# Show help
.PHONY: help
help:
	@echo ""
	@echo "$(COLOR_BLUE)GoCrafter - Go Project Scaffolding Tool$(COLOR_RESET)"
	@echo ""
	@echo "$(COLOR_YELLOW)Development Commands:$(COLOR_RESET)"
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
	@echo "$(COLOR_YELLOW)Docker:$(COLOR_RESET)"
	@echo "  docker-build    Build Docker image"
	@echo ""
	@echo "$(COLOR_YELLOW)Documentation:$(COLOR_RESET)"
	@echo "  docs            Generate documentation"
	@echo ""
	@echo "$(COLOR_YELLOW)Examples:$(COLOR_RESET)"
	@echo "  make run ARGS='list'"
	@echo "  make run ARGS='new my-api --template api-rest'"
	@echo "  make docker-run ARGS='--help'"
	@echo ""


