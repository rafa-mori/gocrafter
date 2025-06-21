# 🛠️ Template Development Guide

This guide explains how to create custom templates for GoCrafter.

## Table of Contents

- [Template Structure](#template-structure)
- [Template Variables](#template-variables)
- [Creating a Custom Template](#creating-a-custom-template)
- [Template Metadata](#template-metadata)
- [Template Functions](#template-functions)
- [Best Practices](#best-practices)
- [Testing Templates](#testing-templates)
- [Contributing Templates](#contributing-templates)

## Template Structure

A GoCrafter template is a directory containing:

```text
my-template/
├── template.json           # Template metadata
├── go.mod                 # Go module template
├── README.md              # Project documentation
├── Makefile              # Build automation
├── cmd/
│   └── main.go           # Main application
├── internal/
│   ├── config/
│   └── handler/
└── pkg/
```

### File Processing

GoCrafter processes files in two ways:

1. **Template Processing**: Files with template syntax `{{.Variable}}`
2. **Direct Copy**: Binary files and assets

### Supported File Types

Template processing is applied to:
- `.go` files
- `.mod`, `.sum` files
- `.yaml`, `.yml` files
- `.json` files
- `.toml` files
- `.md` files
- `.txt` files
- `.env` files
- `Dockerfile`, `Makefile` files

## Template Variables

Templates have access to these variables:

### Basic Variables

```go
type TemplateVars struct {
    ProjectName    string  // e.g., "my-awesome-api"
    ModuleName     string  // e.g., "github.com/user/my-awesome-api"
    PackageName    string  // e.g., "myawesomeapi"
    DatabaseType   string  // e.g., "postgres", "mysql", "mongodb"
    CacheType      string  // e.g., "redis", "memcached"
    QueueType      string  // e.g., "rabbitmq", "kafka"
    HasDocker      bool    // Docker configuration included
    HasKubernetes  bool    // Kubernetes manifests included
    HasMonitoring  bool    // Monitoring features included
    CIType         string  // e.g., "github", "gitlab"
    Features       []string // Selected features
    Custom         map[string]string // Custom variables
}
```

### Usage in Templates

```go
// go.mod template
module {{.ModuleName}}

go 1.21

// main.go template
package main

import (
    "{{.ModuleName}}/internal/config"
    {{- if .DatabaseType}}
    "{{.ModuleName}}/internal/database"
    {{- end}}
)

func main() {
    // {{.ProjectName}} application
    cfg := config.Load()
    {{- if .DatabaseType}}
    db := database.Connect(cfg)
    {{- end}}
}
```

## Creating a Custom Template

### Step 1: Create Template Directory

```bash
mkdir -p templates/my-custom-template
cd templates/my-custom-template
```

### Step 2: Create Template Metadata

Create `template.json`:

```json
{
  "name": "my-custom-template",
  "description": "Custom template for specialized applications",
  "version": "1.0.0",
  "author": "Your Name",
  "tags": ["custom", "specialized"],
  "features": [
    "Custom feature 1",
    "Custom feature 2",
    "Modern tooling"
  ]
}
```

### Step 3: Create Project Structure

```bash
# Create directories
mkdir -p cmd internal/config internal/handler pkg

# Create main.go
cat > cmd/main.go << 'EOF'
package main

import (
    "fmt"
    "{{.ModuleName}}/internal/config"
)

func main() {
    fmt.Println("Hello from {{.ProjectName}}!")
    cfg := config.Load()
    fmt.Printf("Environment: %s\n", cfg.Environment)
}
EOF
```

### Step 4: Create Configuration

```bash
cat > internal/config/config.go << 'EOF'
package config

import "os"

type Config struct {
    Environment string
    Port        string
    {{- if .DatabaseType}}
    DatabaseURL string
    {{- end}}
}

func Load() *Config {
    return &Config{
        Environment: getEnv("ENVIRONMENT", "development"),
        Port:        getEnv("PORT", "8080"),
        {{- if .DatabaseType}}
        DatabaseURL: getEnv("DATABASE_URL", ""),
        {{- end}}
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}
EOF
```

### Step 5: Create Go Module Template

```bash
cat > go.mod << 'EOF'
module {{.ModuleName}}

go 1.21

require (
    {{- if .DatabaseType}}
    {{- if eq .DatabaseType "postgres"}}
    github.com/lib/pq v1.10.9
    {{- else if eq .DatabaseType "mysql"}}
    github.com/go-sql-driver/mysql v1.7.1
    {{- end}}
    {{- end}}
)
EOF
```

### Step 6: Create Documentation

```bash
cat > README.md << 'EOF'
# {{.ProjectName}}

{{.ProjectName}} is a custom application generated with GoCrafter.

## Features

{{- range .Features}}
- {{.}}
{{- end}}

## Getting Started

```bash
# Install dependencies
go mod download

# Run the application
go run cmd/main.go
```

## Configuration

Set these environment variables:

- `ENVIRONMENT`: Application environment (default: development)
- `PORT`: Server port (default: 8080)
{{- if .DatabaseType}}
- `DATABASE_URL`: Database connection string
{{- end}}
EOF
```

## Template Metadata

The `template.json` file defines template metadata:

```json
{
  "name": "template-name",
  "description": "Template description",
  "version": "1.0.0",
  "author": "Author Name",
  "tags": ["tag1", "tag2"],
  "features": [
    "Feature description 1",
    "Feature description 2"
  ],
  "requirements": {
    "go_version": "1.21",
    "dependencies": [
      "github.com/gin-gonic/gin",
      "github.com/spf13/cobra"
    ]
  }
}
```

### Metadata Fields

| Field | Type | Description |
|-------|------|-------------|
| `name` | string | Template name (must match directory) |
| `description` | string | Short description |
| `version` | string | Template version (semver) |
| `author` | string | Template author |
| `tags` | array | Template tags for categorization |
| `features` | array | List of template features |
| `requirements` | object | Requirements and dependencies |

## Template Functions

GoCrafter provides template functions:

### String Functions

```go
{{.ProjectName | lower}}        // Convert to lowercase
{{.ProjectName | upper}}        // Convert to uppercase
{{.ProjectName | title}}        // Convert to title case
```

### Conditional Functions

```go
{{if hasFeature "Authentication"}}
// Include authentication code
{{end}}

{{if contains .Features "JWT"}}
// Include JWT handling
{{end}}
```

### Custom Functions

```go
// Check if a feature is selected
{{if hasFeature "API Documentation"}}
import "github.com/swaggo/swag"
{{end}}

// String manipulation
{{.ModuleName | replace "github.com/" ""}}
```

## Best Practices

### File Organization

1. **Follow Go conventions**: Use standard Go project layout
2. **Separate concerns**: Keep different functionalities in separate packages
3. **Use internal package**: Keep implementation details private

### Template Design

1. **Make it configurable**: Use template variables for customization
2. **Provide defaults**: Always have sensible default values
3. **Document well**: Include comprehensive README and comments

### Variable Usage

```go
// Good: Check if variable exists
{{- if .DatabaseType}}
// Database code here
{{- end}}

// Good: Provide defaults
{{.Port | default "8080"}}

// Avoid: Assuming variables exist
{{.DatabaseType}} // Might be empty
```

### Conditional Blocks

```go
// Database configuration
{{- if .DatabaseType}}
{{- if eq .DatabaseType "postgres"}}
import "github.com/lib/pq"
{{- else if eq .DatabaseType "mysql"}}
import "github.com/go-sql-driver/mysql"
{{- end}}
{{- end}}

// Feature-based inclusion
{{- if hasFeature "Authentication"}}
import "github.com/golang-jwt/jwt/v5"
{{- end}}
```

## Testing Templates

### Manual Testing

```bash
# Generate project with your template
gocrafter new test-project --template my-custom-template

# Verify generated structure
cd test-project
go mod tidy
go build
go test ./...
```

### Automated Testing

Create a test script:

```bash
#!/bin/bash
# test-template.sh

TEMPLATE_NAME="my-custom-template"
TEST_PROJECT="test-${TEMPLATE_NAME}"

echo "Testing template: $TEMPLATE_NAME"

# Clean up previous test
rm -rf "$TEST_PROJECT"

# Generate project
gocrafter new "$TEST_PROJECT" --template "$TEMPLATE_NAME" --quick

# Test the generated project
cd "$TEST_PROJECT"

# Check if it compiles
if go build ./...; then
    echo "✅ Build successful"
else
    echo "❌ Build failed"
    exit 1
fi

# Run tests
if go test ./...; then
    echo "✅ Tests passed"
else
    echo "❌ Tests failed"
    exit 1
fi

# Clean up
cd ..
rm -rf "$TEST_PROJECT"

echo "✅ Template test completed successfully"
```

### Integration Testing

Test with different configurations:

```bash
# Test with different databases
gocrafter new test-postgres --template my-template --quick
# Select postgres database

gocrafter new test-mysql --template my-template --quick
# Select mysql database

# Test with different features
gocrafter new test-features --template my-template --quick
# Select various features
```

## Contributing Templates

### Template Guidelines

1. **Follow naming conventions**: Use lowercase with hyphens
2. **Provide complete examples**: Include working code
3. **Add comprehensive tests**: Ensure template works correctly
4. **Document thoroughly**: Include detailed README and comments

### Submission Process

1. **Fork the repository**
2. **Create template directory** in `templates/`
3. **Add template files** and metadata
4. **Test thoroughly** with different configurations
5. **Submit pull request** with description

### Template Checklist

- [ ] Template metadata (`template.json`) is complete
- [ ] All template variables are documented
- [ ] Generated code compiles and runs
- [ ] README includes setup instructions
- [ ] Makefile includes common tasks
- [ ] Tests are included and pass
- [ ] Docker configuration (if applicable)
- [ ] CI/CD configuration (if applicable)

### Example Pull Request

```markdown
## Add New Template: GraphQL API

### Description
Adds a new template for GraphQL API servers using gqlgen.

### Features
- GraphQL server with gqlgen
- Schema-first development
- Resolver generation
- Authentication middleware
- Database integration
- Docker support

### Testing
- [x] Generated project compiles
- [x] Tests pass
- [x] Docker build works
- [x] Documentation is complete

### Files Added
- `templates/graphql-api/`
- `templates/graphql-api/template.json`
- `templates/graphql-api/schema.graphql`
- ... (other files)
```

## Advanced Features

### Multi-File Templates

For complex file generation:

```go
// In template file: internal/model/{{.PackageName}}.go
package model

type {{.ProjectName | title}} struct {
    ID   uint   `json:"id"`
    Name string `json:"name"`
}
```

### Environment-Specific Files

```go
// docker-compose.yml (only if Docker is enabled)
{{- if .HasDocker}}
version: '3.8'
services:
  {{.PackageName}}:
    build: .
    ports:
      - "8080:8080"
{{- end}}
```

### Dynamic Dependencies

```go
// go.mod with dynamic dependencies
module {{.ModuleName}}

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    {{- if hasFeature "Authentication"}}
    github.com/golang-jwt/jwt/v5 v5.0.0
    {{- end}}
    {{- if .DatabaseType}}
    {{- if eq .DatabaseType "postgres"}}
    gorm.io/driver/postgres v1.5.0
    gorm.io/gorm v1.25.0
    {{- end}}
    {{- end}}
)
```

---

For more information about template development, see the [GoCrafter documentation](https://github.com/rafa-mori/gocrafter/docs).
