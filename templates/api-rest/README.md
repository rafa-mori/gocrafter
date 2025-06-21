# {{.ProjectName}}

{{.ProjectName}} is a REST API built with Go, generated using GoCrafter.

## Features

- ✅ REST API with Gin framework
- ✅ Structured logging with Logrus
- ✅ Environment-based configuration
- ✅ Graceful shutdown
- ✅ Health check endpoints
{{- if .DatabaseType}}
- ✅ {{.DatabaseType | title}} database integration
{{- end}}
{{- if .CacheType}}
- ✅ {{.CacheType | title}} caching
{{- end}}
{{- if hasFeature "Authentication"}}
- ✅ JWT authentication  
{{- end}}
{{- if hasFeature "API Documentation"}}
- ✅ Swagger API documentation
{{- end}}
{{- if .HasDocker}}
- ✅ Docker support
{{- end}}
{{- if .HasKubernetes}}
- ✅ Kubernetes manifests
{{- end}}

## Getting Started

### Prerequisites

- Go 1.21 or higher
{{- if .DatabaseType}}
- {{.DatabaseType | title}} database
{{- end}}
{{- if .CacheType}}
- {{.CacheType | title}} server
{{- end}}

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd {{.ProjectName}}
```

2. Install dependencies:
```bash
go mod download
```

3. Copy environment variables:
```bash
cp .env.example .env
```

4. Update the `.env` file with your configuration.

5. Run the application:
```bash
make run
```

The server will start on `http://localhost:8080`

## API Endpoints

### Health Check
- `GET /health` - Health check endpoint

### Users (Example)
- `GET /api/v1/users` - Get all users
- `POST /api/v1/users` - Create a new user
- `GET /api/v1/users/:id` - Get user by ID
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

{{- if hasFeature "API Documentation"}}
### API Documentation
- `GET /swagger/index.html` - Swagger UI
{{- end}}

## Configuration

The application uses environment variables for configuration:

| Variable | Description | Default |
|----------|-------------|---------|
| `ENVIRONMENT` | Application environment | `development` |
| `PORT` | Server port | `8080` |
| `LOG_LEVEL` | Log level | `info` |
{{- if .DatabaseType}}
| `DB_HOST` | Database host | `localhost` |
| `DB_PORT` | Database port | `{{- if eq .DatabaseType "postgres"}}5432{{- else if eq .DatabaseType "mysql"}}3306{{- else}}3306{{- end}}` |
| `DB_USER` | Database user | `{{- if eq .DatabaseType "postgres"}}postgres{{- else}}root{{- end}}` |
| `DB_PASSWORD` | Database password | `` |
| `DB_NAME` | Database name | `{{.PackageName}}` |
{{- end}}

## Development

### Running Tests
```bash
make test
```

### Building the Application
```bash
make build
```

### Linting
```bash
make lint
```

### Running with Docker
```bash
make docker-run
```

## Project Structure

```
{{.ProjectName}}/
├── cmd/
│   └── main.go          # Application entry point
├── internal/
│   ├── config/          # Configuration management
│   ├── handler/         # HTTP handlers
│   ├── middleware/      # HTTP middleware
│   ├── model/           # Data models
│   ├── repository/      # Data access layer
│   └── service/         # Business logic
├── pkg/                 # Public packages
├── docs/                # Documentation
├── scripts/             # Build and deployment scripts
├── .env.example         # Environment variables template
├── Dockerfile           # Docker configuration
├── Makefile            # Build automation
└── README.md           # This file
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for your changes
5. Run tests and ensure they pass
6. Submit a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
