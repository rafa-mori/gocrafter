# 📖 GoCrafter User Guide

This comprehensive guide will help you get the most out of GoCrafter, from basic usage to advanced features.

## Table of Contents

- [Installation](#installation)
- [Basic Usage](#basic-usage)
- [Interactive Mode](#interactive-mode)
- [Quick Mode](#quick-mode)
- [Templates](#templates)
- [Configuration](#configuration)
- [Advanced Features](#advanced-features)
- [Best Practices](#best-practices)
- [Troubleshooting](#troubleshooting)

## Installation

### Using Go Install (Recommended)

```bash
go install github.com/rafa-mori/gocrafter@latest
```

### Download Binary

```bash
# Linux
curl -sSL https://github.com/rafa-mori/gocrafter/releases/latest/download/gocrafter-linux-amd64.tar.gz | tar xz

# macOS
curl -sSL https://github.com/rafa-mori/gocrafter/releases/latest/download/gocrafter-darwin-amd64.tar.gz | tar xz

# Windows
curl -sSL https://github.com/rafa-mori/gocrafter/releases/latest/download/gocrafter-windows-amd64.zip -o gocrafter.zip
unzip gocrafter.zip
```

### Build from Source

```bash
git clone https://github.com/rafa-mori/gocrafter.git
cd gocrafter
go build -o gocrafter cmd/main.go cmd/wrpr.go cmd/usage.go
```

## Basic Usage

### Available Commands

```bash
gocrafter --help                    # Show help
gocrafter version                   # Show version
gocrafter list                      # List available templates
gocrafter info <template>           # Show template details
gocrafter new                       # Create new project (interactive)
gocrafter new [name] [flags]        # Create new project (quick)
```

### Your First Project

1. **Interactive Mode (Recommended for beginners):**
   ```bash
   gocrafter new
   ```

2. **Quick Mode:**
   ```bash
   gocrafter new my-api --template api-rest
   ```

3. **Navigate to your project:**
   ```bash
   cd my-api
   ```

4. **Start developing:**
   ```bash
   make run    # Start the application
   make test   # Run tests
   make build  # Build the application
   ```

## Interactive Mode

The interactive mode provides a guided experience for creating projects:

### Step-by-Step Process

1. **Project Information**
   - Project name
   - Go module name

2. **Template Selection**
   - Choose from available templates
   - View descriptions and features

3. **Database Configuration**
   - Select database type
   - Configure cache layer

4. **Feature Selection**
   - Authentication options
   - API documentation
   - Monitoring and metrics
   - Middleware options

5. **DevOps Integration**
   - Docker configuration
   - Kubernetes manifests
   - CI/CD pipeline

6. **Confirmation**
   - Review configuration
   - Confirm and generate

### Example Interactive Session

```bash
$ gocrafter new
🚀 Welcome to GoCrafter - Go Project Generator!
Let's craft your perfect Go project together...

? What's your project name? blog-api
? What's your Go module name? github.com/myuser/blog-api
? What type of project do you want to create? 
  ❯ api-rest - REST API with HTTP server (Gin/Fiber)
    cli-tool - Command-line application with Cobra
    microservice - Microservice with gRPC and HTTP
    grpc-service - Pure gRPC service
    worker - Background worker/job processor
    library - Go library/package

? Which database do you want to use? 
  ❯ postgres
    mysql
    mongodb
    sqlite
    none

? Do you want to add a cache layer? 
  ❯ redis
    memcached
    none

? Which additional features do you want to include? 
  ◉ Authentication (JWT)
  ◉ API Documentation (Swagger)
  ◯ Health Checks
  ◉ Metrics (Prometheus)
  ◯ Distributed Tracing
  ◯ Rate Limiting
  ◉ CORS Middleware
  ◯ Request Validation

? Include Docker configuration? Yes
? Include Kubernetes manifests? No
? Which CI/CD system do you want to use? 
  ❯ github
    gitlab
    jenkins
    azure
    none

📋 Project Configuration Summary:
  Name: blog-api
  Module: github.com/myuser/blog-api
  Template: api-rest
  Database: postgres
  Cache: redis
  Features: Authentication (JWT), API Documentation (Swagger), Metrics (Prometheus), CORS Middleware
  Docker: true
  Kubernetes: false
  CI/CD: github

? Does this look correct? Proceed with project generation? Yes

🚀 Starting project generation...
✅ Project generated successfully!
📁 Location: blog-api
```

## Quick Mode

Quick mode allows for rapid project creation with minimal prompts:

### Basic Quick Mode

```bash
# Create project with default settings
gocrafter new my-project --template api-rest

# Specify module name
gocrafter new my-project --template api-rest
# Then enter module name when prompted
```

### Advanced Quick Mode

```bash
# Create project in specific directory
gocrafter new my-api --template api-rest --output ./projects

# Use configuration file
gocrafter new my-api --config project-config.json

# Quick mode with minimal prompts
gocrafter new my-api --template api-rest --quick
```

### Flags and Options

| Flag | Short | Description |
|------|-------|-------------|
| `--template` | `-t` | Specify template name |
| `--output` | `-o` | Output directory |
| `--config` | `-c` | Configuration file |
| `--quick` | `-q` | Quick mode with minimal prompts |

## Templates

GoCrafter provides several built-in templates for different project types:

### API REST Template

**Description:** Full-featured REST API server with modern Go practices.

**Features:**
- Gin web framework
- Structured logging
- Environment configuration
- Database integration
- Authentication middleware
- API documentation
- Health checks
- Docker support

**Use Cases:**
- Web APIs
- Backend services
- RESTful microservices

**Generated Structure:**
```
my-api/
├── cmd/main.go
├── internal/
│   ├── config/
│   ├── handler/
│   ├── middleware/
│   ├── model/
│   ├── repository/
│   └── service/
├── pkg/
├── Makefile
├── Dockerfile
└── README.md
```

### CLI Tool Template

**Description:** Command-line application with Cobra framework.

**Features:**
- Cobra CLI framework
- Subcommands support
- Configuration management
- Structured logging
- Cross-platform builds

**Use Cases:**
- Command-line utilities
- Development tools
- System administration tools

### Microservice Template

**Description:** Microservice with gRPC and HTTP endpoints.

**Features:**
- gRPC server
- HTTP gateway
- Service discovery
- Health checks
- Metrics collection
- Distributed tracing

**Use Cases:**
- Distributed systems
- Service-oriented architecture
- Cloud-native applications

### gRPC Service Template

**Description:** Pure gRPC service with protocol buffers.

**Features:**
- Protocol buffers
- gRPC server
- Client generation
- Streaming support
- Service mesh ready

**Use Cases:**
- High-performance services
- Inter-service communication
- Real-time applications

### Worker Template

**Description:** Background job processor.

**Features:**
- Queue integration
- Job processing
- Retry mechanisms
- Monitoring hooks
- Graceful shutdown

**Use Cases:**
- Background processing
- Asynchronous tasks
- Data processing pipelines

### Library Template

**Description:** Go library/package template.

**Features:**
- Package structure
- Documentation templates
- Testing framework
- CI/CD workflows
- Version management

**Use Cases:**
- Reusable packages
- Open source libraries
- Internal utilities

## Configuration

### Environment Variables

Projects generated by GoCrafter use environment variables for configuration:

```bash
# Application settings
ENVIRONMENT=development
PORT=8080
LOG_LEVEL=info

# Database settings (if applicable)
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=secret
DB_NAME=myapp

# Cache settings (if applicable)
REDIS_URL=redis://localhost:6379
REDIS_PASSWORD=
REDIS_DB=0

# Authentication (if applicable)
JWT_SECRET=your-secret-key
JWT_EXPIRATION=3600
```

### Configuration Files

You can use configuration files for project generation:

**project-config.json:**
```json
{
  "name": "my-awesome-api",
  "module": "github.com/myuser/my-awesome-api",
  "template": "api-rest",
  "database": "postgres",
  "cache": "redis",
  "features": [
    "Authentication (JWT)",
    "API Documentation (Swagger)",
    "Health Checks"
  ],
  "docker": true,
  "kubernetes": false,
  "ci": "github"
}
```

**Usage:**
```bash
gocrafter new --config project-config.json
```

## Advanced Features

### Custom Output Directory

```bash
# Create project in specific directory
gocrafter new my-api --template api-rest --output /path/to/projects

# Create project in subdirectory
gocrafter new my-api --template api-rest --output ./workspace/apis
```

### Template Information

```bash
# List all templates
gocrafter list

# Get detailed template information
gocrafter info api-rest

# Show template structure (if available)
gocrafter info api-rest --show-structure
```

### Batch Project Creation

Create multiple projects using a script:

```bash
#!/bin/bash

# Create multiple microservices
services=("user-service" "order-service" "notification-service")

for service in "${services[@]}"; do
  gocrafter new "$service" \
    --template microservice \
    --output ./microservices \
    --quick
done
```

## Best Practices

### Project Naming

- Use lowercase with hyphens: `my-awesome-api`
- Be descriptive: `user-management-service`
- Avoid special characters: stick to letters, numbers, and hyphens

### Module Naming

- Follow Go module conventions: `github.com/username/project-name`
- Use your actual repository path
- Keep it consistent with your project structure

### Template Selection

- **API REST**: For HTTP APIs and web services
- **CLI Tool**: For command-line applications
- **Microservice**: For distributed systems
- **gRPC Service**: For high-performance inter-service communication
- **Worker**: For background processing
- **Library**: For reusable packages

### Feature Selection

- Start with basic features and add more later
- Consider your deployment environment
- Think about monitoring and observability early

### Development Workflow

1. Generate project with GoCrafter
2. Initialize Git repository
3. Set up development environment
4. Configure database and dependencies
5. Start with tests
6. Implement features incrementally
7. Set up CI/CD pipeline

## Troubleshooting

### Common Issues

#### Template Not Found

**Error:** `template 'api-rest' not found`

**Solution:**
- Check available templates: `gocrafter list`
- Ensure correct template name
- Verify GoCrafter installation

#### Permission Denied

**Error:** `permission denied`

**Solution:**
- Check directory permissions
- Ensure you have write access to output directory
- Run with appropriate permissions

#### Module Import Issues

**Error:** Module import problems in generated project

**Solution:**
- Verify module name is correct
- Run `go mod tidy` in project directory
- Check Go version compatibility

#### Binary Not Found

**Error:** `gocrafter: command not found`

**Solution:**
- Ensure `$GOPATH/bin` is in your `$PATH`
- Verify installation: `which gocrafter`
- Try absolute path to binary

### Getting Help

1. **Check documentation:** Read this guide and README
2. **Search issues:** Look for similar problems on GitHub
3. **Create issue:** Report bugs or request features
4. **Community:** Join discussions and ask questions

### Debug Mode

Enable debug logging for troubleshooting:

```bash
export GOCRAFTER_LOG_LEVEL=debug
gocrafter new my-project --template api-rest
```

## Tips and Tricks

### Aliases

Create shell aliases for common operations:

```bash
# Add to your .bashrc or .zshrc
alias gcnew='gocrafter new'
alias gclist='gocrafter list'
alias gcinfo='gocrafter info'
```

### Templates Directory

Find where templates are stored:

```bash
gocrafter info api-rest --show-path
```

### Quick Project Setup

Create a complete project setup script:

```bash
#!/bin/bash
PROJECT_NAME=$1
TEMPLATE=${2:-api-rest}

echo "Creating project: $PROJECT_NAME with template: $TEMPLATE"

gocrafter new "$PROJECT_NAME" --template "$TEMPLATE" --quick
cd "$PROJECT_NAME"

echo "Setting up Git repository..."
git init
git add .
git commit -m "Initial commit from GoCrafter"

echo "Installing dependencies..."
go mod tidy

echo "Running initial test..."
make test

echo "Project $PROJECT_NAME is ready!"
echo "Next steps:"
echo "  cd $PROJECT_NAME"
echo "  make run"
```

---

For more information, visit the [GoCrafter GitHub repository](https://github.com/rafa-mori/gocrafter).
