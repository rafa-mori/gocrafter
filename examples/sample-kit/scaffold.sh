#!/bin/bash

# Post-generation script for golang-basic-api kit
# This script runs after the project is generated

set -e

echo "🚀 Running post-generation setup for {{project_name}}..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go {{go_version}} or higher."
    exit 1
fi

# Check Go version
GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
REQUIRED_VERSION="{{go_version}}"

if [ "$(printf '%s\n' "$REQUIRED_VERSION" "$GO_VERSION" | sort -V | head -n1)" != "$REQUIRED_VERSION" ]; then
    echo "⚠️  Warning: Go version $GO_VERSION is lower than required $REQUIRED_VERSION"
fi

# Initialize Go module
echo "📦 Initializing Go module..."
go mod tidy

# Create necessary directories
echo "📁 Creating project directories..."
mkdir -p bin
mkdir -p logs
mkdir -p docs
mkdir -p internal/handler
mkdir -p internal/middleware
mkdir -p internal/model
mkdir -p internal/service
mkdir -p internal/config

# Create .gitignore
echo "📝 Creating .gitignore..."
cat > .gitignore << 'EOF'
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib
bin/
*.test
*.out

# Go workspace file
go.work

# Dependency directories
vendor/

# Go build cache
.cache/

# IDE directories
.vscode/
.idea/
*.swp
*.swo
*~

# OS generated files
.DS_Store
.DS_Store?
._*
.Spotlight-V100
.Trashes
ehthumbs.db
Thumbs.db

# Environment files
.env
.env.local
.env.*.local

# Log files
*.log
logs/

# Coverage files
coverage.out
coverage.html

# Air configuration
.air.toml

# Temporary files
tmp/
temp/
EOF

# Create basic project structure files
echo "🏗️  Creating project structure..."

# Create internal/config/config.go
cat > internal/config/config.go << 'EOF'
package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port     string
	GinMode  string
	LogLevel string
	AppName  string
	Version  string
}

func New() *Config {
	return &Config{
		Port:     getEnv("PORT", "8080"),
		GinMode:  getEnv("GIN_MODE", "debug"),
		LogLevel: getEnv("LOG_LEVEL", "info"),
		AppName:  getEnv("APP_NAME", "{{project_name}}"),
		Version:  getEnv("APP_VERSION", "1.0.0"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}
EOF

# Check if make is available
if command -v make &> /dev/null; then
    echo "🔧 Running initial build..."
    make deps
    make build
else
    echo "⚠️  Make is not available. Running go build directly..."
    go build -o bin/{{project_name}} main.go
fi

# Create development tools configuration
echo "⚙️  Creating development configuration..."

# Create .air.toml for hot reloading
cat > .air.toml << 'EOF'
root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
  args_bin = []
  bin = "./tmp/main"
  cmd = "go build -o ./tmp/main ."
  delay = 1000
  exclude_dir = ["assets", "tmp", "vendor", "testdata"]
  exclude_file = []
  exclude_regex = ["_test.go"]
  exclude_unchanged = false
  follow_symlink = false
  full_bin = ""
  include_dir = []
  include_ext = ["go", "tpl", "tmpl", "html"]
  include_file = []
  kill_delay = "0s"
  log = "build-errors.log"
  poll = false
  poll_interval = 0
  rerun = false
  rerun_delay = 500
  send_interrupt = false
  stop_on_root = false

[color]
  app = ""
  build = "yellow"
  main = "magenta"
  runner = "green"
  watcher = "cyan"

[log]
  main_only = false
  time = false

[misc]
  clean_on_exit = false

[screen]
  clear_on_rebuild = false
  keep_scroll = true
EOF

echo "✅ Post-generation setup completed successfully!"
echo ""
echo "📋 Next steps:"
echo "   1. cd {{project_name}}"
echo "   2. cp .env.example .env  # Configure environment variables"
echo "   3. make run              # Start the application"
echo "   4. Open http://localhost:{{port}}/health to test"
echo ""
echo "🛠️  Available commands:"
echo "   make help         # Show all available commands"
echo "   make run          # Run the application"
echo "   make test         # Run tests"
echo "   make build        # Build the application"
echo "   make docker-build # Build Docker image"
echo ""
echo "🎉 Happy coding!"
