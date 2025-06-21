# GoCrafter Examples

This directory contains practical examples and tutorials for using GoCrafter effectively.

## Quick Start Examples

### 1. Creating a Simple REST API

```bash
# Interactive mode (recommended for beginners)
gocrafter new blog-api

# Follow the prompts:
# - Template: api-rest
# - Database: PostgreSQL
# - Features: Authentication, Swagger, Health Checks
```

**Generated structure:**

```
blog-api/
├── cmd/main.go
├── internal/
│   ├── config/
│   ├── handler/
│   ├── middleware/
│   ├── model/
│   ├── repository/
│   └── service/
├── Makefile
├── Dockerfile
└── README.md
```

**Next steps:**

```bash
cd blog-api
make setup    # Install dependencies
make run      # Start the development server
make test     # Run tests
```

### 2. Creating a CLI Tool

```bash
# Quick mode with specific template
gocrafter new my-cli --template cli-tool

cd my-cli
go run cmd/main.go --help
```

### 3. Creating a Microservice

```bash
# With custom configuration
gocrafter new user-service \
  --template microservice \
  --output ./services/
```

## Real-World Examples

### Blog API with Full Features

This example shows how to create a complete blog API with authentication, database, and API documentation.

```bash
# Create the project
gocrafter new blog-api --template api-rest

# Answer the prompts:
# - Database: PostgreSQL
# - Cache: Redis
# - Features: JWT Auth, Swagger, Metrics, Docker
```

**Features included:**

- REST endpoints for posts, users, comments
- JWT authentication and authorization
- PostgreSQL database with migrations
- Redis caching layer
- API documentation with Swagger
- Docker containerization
- Health checks and metrics
- Comprehensive test suite

**Project structure:**

```
blog-api/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go          # Configuration management
│   ├── handler/
│   │   ├── auth.go            # Authentication endpoints
│   │   ├── post.go            # Blog post endpoints
│   │   └── user.go            # User management endpoints
│   ├── middleware/
│   │   ├── auth.go            # JWT middleware
│   │   ├── cors.go            # CORS handling
│   │   └── logging.go         # Request logging
│   ├── model/
│   │   ├── post.go            # Post model
│   │   └── user.go            # User model
│   ├── repository/
│   │   ├── post.go            # Post data access
│   │   └── user.go            # User data access
│   └── service/
│       ├── auth.go            # Authentication logic
│       ├── post.go            # Post business logic
│       └── user.go            # User business logic
├── pkg/
│   ├── database/              # Database utilities
│   ├── redis/                 # Redis utilities
│   └── validation/            # Input validation
├── migrations/
│   ├── 001_create_users.sql
│   └── 002_create_posts.sql
├── docs/
│   └── api.yaml              # OpenAPI specification
├── docker-compose.yml        # Development environment
├── Dockerfile               # Production container
├── Makefile                 # Build automation
└── README.md               # Project documentation
```

### CLI Tool with Subcommands

Create a feature-rich CLI tool with multiple subcommands and configuration.

```bash
gocrafter new devtools --template cli-tool
```

**Generated CLI features:**

- Multiple subcommands (init, build, deploy)
- Configuration file support
- Colored output and progress bars
- Shell completion
- Cross-platform builds

### Microservice with gRPC

Create a production-ready microservice with gRPC and HTTP support.

```bash
gocrafter new user-service --template microservice
```

**Features:**

- gRPC service definitions
- HTTP gateway integration
- Service discovery ready
- Observability built-in
- Kubernetes deployment manifests

## Advanced Usage Examples

### Custom Template Configuration

Using a configuration file to pre-define project settings:

```json
{
  "name": "my-ecommerce-api",
  "module_name": "github.com/mycompany/ecommerce-api",
  "template": "api-rest",
  "database": "postgres",
  "cache": "redis",
  "features": [
    "jwt-auth",
    "swagger",
    "metrics",
    "docker",
    "kubernetes"
  ],
  "additional": {
    "payment_gateway": "stripe",
    "email_service": "sendgrid"
  }
}
```

```bash
gocrafter new --config project-config.json
```

### Batch Project Creation

Create multiple related services:

```bash
# API Gateway
gocrafter new api-gateway --template api-rest

# User Service
gocrafter new user-service --template microservice

# Order Service
gocrafter new order-service --template microservice

# Notification Service
gocrafter new notification-service --template worker
```

### Template Customization

Extending existing templates for organization-specific needs:

```bash
# Copy existing template
cp -r templates/api-rest templates/company-api

# Customize template.json
# Add company-specific configurations
# Modify template files as needed

# Use custom template
gocrafter new company-project --template company-api
```

## Development Workflow Examples

### Full Development Cycle

```bash
# 1. Create project
gocrafter new my-api --template api-rest

cd my-api

# 2. Set up development environment
make setup
docker-compose up -d postgres redis

# 3. Run database migrations
make migrate

# 4. Start development server
make dev

# 5. Run tests
make test

# 6. Build for production
make build

# 7. Deploy
make deploy
```

### Testing Generated Projects

```bash
# Unit tests
make test

# Integration tests
make test-integration

# Load testing
make test-load

# Security testing
make test-security
```

### Continuous Integration

Generated projects include CI/CD configurations:

**GitHub Actions:**

```yaml
# .github/workflows/ci.yml
name: CI
on: [push, pull_request]
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v3
        with:
          go-version: '1.21'
      - run: make test
      - run: make build
```

## Best Practices Examples

### Project Organization

```
my-service/
├── cmd/                    # Main applications
├── internal/              # Private code
├── pkg/                   # Public libraries
├── api/                   # API definitions
├── configs/               # Configuration files
├── deployments/           # Deployment configurations
├── docs/                  # Documentation
├── scripts/               # Build and deployment scripts
├── test/                  # Integration tests
└── vendor/                # Dependencies (if using vendor)
```

### Configuration Management

```go
// internal/config/config.go
type Config struct {
    Server   ServerConfig   `yaml:"server"`
    Database DatabaseConfig `yaml:"database"`
    Redis    RedisConfig    `yaml:"redis"`
    Auth     AuthConfig     `yaml:"auth"`
}

// Environment-specific configs
// config/development.yaml
// config/production.yaml
// config/testing.yaml
```

### Error Handling

```go
// pkg/errors/errors.go
type APIError struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    Details string `json:"details,omitempty"`
}

func (e APIError) Error() string {
    return e.Message
}
```

### Testing Patterns

```go
// internal/handler/post_test.go
func TestCreatePost(t *testing.T) {
    tests := []struct {
        name           string
        input          CreatePostRequest
        expectedStatus int
        expectedError  string
    }{
        {
            name: "valid post creation",
            input: CreatePostRequest{
                Title:   "Test Post",
                Content: "Test content",
            },
            expectedStatus: http.StatusCreated,
        },
        // ... more test cases
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

## Troubleshooting Examples

### Common Issues and Solutions

**Issue: Module name conflicts**

```bash
# Problem: Module name already exists
gocrafter new my-project --module github.com/existing/repo

# Solution: Use a unique module name
gocrafter new my-project --module github.com/myusername/unique-name
```

**Issue: Template not found**

```bash
# Problem: Custom template not recognized
gocrafter new project --template my-template

# Solution: Check template exists
gocrafter list
```

**Issue: Permission errors**

```bash
# Problem: Cannot create files in directory
# Solution: Check permissions or use different output directory
gocrafter new project --output ~/projects/
```

## Performance Examples

### Optimized Project Structure

```
high-performance-api/
├── cmd/main.go
├── internal/
│   ├── handler/
│   │   └── handler.go      # Optimized handlers
│   ├── middleware/
│   │   └── cache.go        # Caching middleware
│   └── pool/
│       └── worker.go       # Worker pools
├── pkg/
│   ├── metrics/            # Performance metrics
│   └── profiling/          # Profiling utilities
└── configs/
    └── performance.yaml    # Performance tuning
```

### Monitoring and Observability

Generated projects include observability features:

```go
// Metrics collection
prometheus.NewCounterVec(
    prometheus.CounterOpts{
        Name: "http_requests_total",
        Help: "Total number of HTTP requests",
    },
    []string{"method", "endpoint", "status"},
)

// Distributed tracing
span := opentracing.StartSpan("http_request")
defer span.Finish()
```

## Integration Examples

### Database Integration

```go
// Multiple database support
switch config.Database.Type {
case "postgres":
    db = setupPostgreSQL(config.Database.URL)
case "mysql":
    db = setupMySQL(config.Database.URL)
case "mongodb":
    db = setupMongoDB(config.Database.URL)
}
```

### External Services

```go
// Service integrations
type Services struct {
    EmailService    email.Service
    PaymentService  payment.Service
    StorageService  storage.Service
}
```

These examples demonstrate the versatility and power of GoCrafter for creating production-ready Go applications across different domains and use cases.
