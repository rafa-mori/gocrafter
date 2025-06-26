# {{project_name}}

{{description}}

## Features

- REST API with Gin framework
- Structured logging with Logrus
- Environment configuration
- Docker support
- Health check endpoint

## Getting Started

### Prerequisites

- Go {{go_version}} or higher
- Make (optional, for build automation)
- Docker (optional, for containerization)

### Installation

1. Clone the repository:

```bash
git clone <repository-url>
cd {{project_name}}
```

2. Install dependencies:

```bash
go mod download
```

3. Create a `.env` file (optional):

```bash
cp .env.example .env
```

### Running the Application

#### Local Development

```bash
# Run directly
go run main.go

# Or using make
make run
```

#### Using Docker

```bash
# Build and run with Docker
make docker-run
```

### API Endpoints

- `GET /health` - Health check endpoint
- `GET /api/v1/hello?name=<name>` - Hello world endpoint
- `GET /api/v1/info` - Project information

### Examples

```bash
# Health check
curl http://localhost:{{port}}/health

# Hello endpoint
curl http://localhost:{{port}}/api/v1/hello?name=Developer

# Project info
curl http://localhost:{{port}}/api/v1/info
```

## Development

### Building

```bash
# Build binary
make build

# Build for different platforms
make build-linux
make build-windows
make build-darwin
```

### Testing

```bash
# Run tests
make test

# Run tests with coverage
make test-coverage
```

### Linting

```bash
# Run linter
make lint
```

## Docker

### Build Image

```bash
make docker-build
```

### Run Container

```bash
make docker-run
```

## Configuration

Environment variables:

- `PORT` - Server port (default: {{port}})
- `GIN_MODE` - Gin mode (debug, release, test)
- `LOG_LEVEL` - Log level (debug, info, warn, error)

## License

This project is licensed under the {{license}} License - see the [LICENSE](LICENSE) file for details.

## Author

{{author}}

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request
