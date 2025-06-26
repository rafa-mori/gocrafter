# Kit Development Guide

This guide explains how to create and share your own kits for GoCrafter.

## What is a Kit?

A kit is a pluggable project template that can be installed from external repositories. Unlike built-in templates, kits are distributed separately and can be maintained by the community.

## Kit Structure

A kit must have the following structure:

```
my-kit/
├── metadata.yaml       # Kit metadata and configuration
├── templates/          # Template files and directories
│   ├── main.go
│   ├── go.mod
│   ├── README.md
│   └── ...
└── scaffold.sh         # Optional post-generation script
```

### Required Files

#### metadata.yaml

The metadata file defines the kit's properties and requirements:

```yaml
name: "my-golang-kit"
description: "A sample Go project kit"
language: "go"
version: "1.0.0"
author: "Your Name"
repository: "https://github.com/user/my-golang-kit"
dependencies:
  - "go"
  - "make"
  - "docker"
placeholders:
  - "project_name"
  - "author"
  - "license"
  - "description"
  - "port"
tags:
  - "api"
  - "rest"
  - "go"
metadata:
  min_go_version: "1.24.4"
  features:
    - "HTTP server"
    - "Docker support"
    - "Makefile automation"
```

#### templates/

The templates directory contains all the files and directories that will be generated. Files can contain placeholders that will be replaced during generation.

### Optional Files

#### scaffold.sh

A post-generation script that runs after the project is created. This script can:

- Initialize dependencies
- Run setup commands
- Create additional files
- Configure the development environment

```bash
#!/bin/bash
set -e

echo "Setting up {{project_name}}..."

# Initialize Go module
go mod tidy

# Create directories
mkdir -p bin logs

# Install development tools
go install github.com/air-verse/air@latest

echo "Setup completed!"
```

## Placeholder System

GoCrafter supports a powerful placeholder system with the following features:

### Basic Placeholders

Use `{{placeholder_name}}` in your template files:

```go
package main

import "fmt"

func main() {
    fmt.Println("Welcome to {{project_name}}!")
    fmt.Println("Created by {{author}}")
}
```

### Built-in Placeholders

The following placeholders are automatically available:

- `{{project_name}}` - Project name
- `{{author}}` - Author name
- `{{license}}` - License type
- `{{current_year}}` - Current year
- `{{go_version}}` - Go version

### Derived Placeholders

Some placeholders are automatically derived from others:

- `{{package_name}}` - Derived from project_name (lowercase, no hyphens)
- `{{module_name}}` - Derived from project_name (lowercase with hyphens)
- `{{class_name}}` - Derived from project_name (title case)
- `{{const_name}}` - Derived from project_name (uppercase with underscores)

### Template Functions

Use template functions for advanced transformations:

```go
// String manipulation
{{upper project_name}}        // UPPERCASE
{{lower project_name}}        // lowercase  
{{title project_name}}        // Title Case
{{kebab project_name}}        // kebab-case
{{snake project_name}}        // snake_case
{{camel project_name}}        // camelCase
{{pascal project_name}}       // PascalCase

// Utilities
{{now}}                       // Current timestamp
{{date "2006-01-02"}}         // Formatted date
{{env "HOME"}}                // Environment variable
{{default "defaultValue" value}} // Default value if empty
```

### Conditional Logic

Use conditional logic in templates:

```go
{{if eq license "MIT"}}
// MIT License specific code
{{else if eq license "Apache-2.0"}}
// Apache License specific code
{{else}}
// Default license code
{{end}}
```

## File Processing

### Template Files

Files with the following extensions are automatically processed as templates:

- Code files: `.go`, `.js`, `.ts`, `.py`, `.rs`, `.java`, `.cpp`, `.c`, `.h`
- Configuration: `.yaml`, `.yml`, `.json`, `.toml`, `.env`
- Documentation: `.md`, `.txt`
- Build files: `Makefile`, `Dockerfile`, `.sh`

### Static Files

Binary files and other non-text files are copied as-is without processing.

### File Path Placeholders

You can use placeholders in file and directory names:

```
templates/
├── cmd/{{project_name}}/
│   └── main.go
├── internal/{{package_name}}/
│   └── service.go
└── {{project_name}}.yaml
```

## Best Practices

### 1. Comprehensive Metadata

Provide complete metadata including:

- Clear description
- Version information
- Dependencies list
- Required placeholders
- Relevant tags

### 2. Flexible Placeholders

Design your kit to work with minimal required placeholders:

```yaml
placeholders:
  - "project_name"    # Required
  - "author"          # Optional, has default
  - "license"         # Optional, defaults to MIT
```

### 3. Meaningful Defaults

Use the `default` template function for optional values:

```go
const DefaultPort = "{{default "8080" port}}"
const Author = "{{default "Developer" author}}"
```

### 4. Documentation

Include comprehensive documentation:

- README.md with setup instructions
- Code comments explaining key concepts
- Example usage and API documentation

### 5. Post-Generation Script

Use scaffold.sh for:

- Dependency installation
- Initial build/setup
- Development tool configuration
- Environment validation

### 6. Error Handling

Make your post-generation script robust:

```bash
#!/bin/bash
set -e  # Exit on error

# Check dependencies
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed"
    exit 1
fi

# Validate Go version
GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
if [ "$(printf '%s\n' "1.24.4" "$GO_VERSION" | sort -V | head -n1)" != "1.24.4" ]; then
    echo "Warning: Go version $GO_VERSION may not be supported"
fi
```

## Testing Your Kit

### 1. Local Testing

Test your kit locally before publishing:

```bash
# Add your local kit
gocrafter kit add /path/to/your/kit

# Generate a test project
gocrafter new test-project --kit your-kit-name

# Verify the generated project
cd test-project
make build
make test
```

### 2. Validation

Ensure your kit passes validation:

```bash
# The kit manager automatically validates:
# - metadata.yaml exists and is valid
# - templates/ directory exists
# - Required fields are present
# - Placeholders are properly defined
```

## Publishing Your Kit

### 1. Repository Setup

Create a Git repository with your kit:

```bash
git init
git add .
git commit -m "Initial kit version"
git tag v1.0.0
git push origin main --tags
```

### 2. Distribution

Users can install your kit using:

```bash
# From GitHub
gocrafter kit add https://github.com/user/your-kit

# From other Git repositories
gocrafter kit add https://gitlab.com/user/your-kit

# From archives
gocrafter kit add https://example.com/kits/your-kit.tar.gz
```

### 3. Versioning

Use semantic versioning for your kits:

- `v1.0.0` - Initial release
- `v1.0.1` - Bug fixes
- `v1.1.0` - New features
- `v2.0.0` - Breaking changes

## Example Kit

See the complete example kit in `examples/sample-kit/` for a working reference implementation.

## Troubleshooting

### Common Issues

1. **Kit validation fails**
   - Check metadata.yaml syntax
   - Ensure all required fields are present
   - Verify templates/ directory exists

2. **Placeholders not replaced**
   - Check placeholder names match metadata
   - Ensure file extensions are processed
   - Verify template syntax

3. **Post-generation script fails**
   - Make script executable
   - Use `set -e` for error handling
   - Test script independently

### Debug Mode

Enable debug logging for troubleshooting:

```bash
export GOCRAFTER_LOG_LEVEL=debug
gocrafter kit add your-kit-url
```

## Community

- Share your kits on GitHub with the `gocrafter-kit` topic
- Join discussions in the GoCrafter community
- Contribute improvements to the kit system

Happy kit development! 🎉
