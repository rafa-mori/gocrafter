 ![GoCrafter Banner](/docs/assets/top_banner.png)

# **GoCrafter**

## A powerful Go project scaffolding and templating tool that helps you create production-ready Go projects with best practices, modern tooling, and customizable templates

[![Build](https://github.com/rafa-mori/gocrafter/actions/workflows/release.yml/badge.svg)](https://github.com/rafa-mori/gocrafter/actions/workflows/release.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/go-%3E=1.21-blue)](go.mod)
[![Releases](https://img.shields.io/github/v/release/rafa-mori/gocrafter?include_prereleases)](https://github.com/rafa-mori/gocrafter/releases)

---

## ✨ Features

- 🎯 **Interactive Project Creation** - Guided wizard for project setup
- 📦 **Multiple Templates** - API REST, CLI tools, microservices, gRPC services, and more
- ⚙️ **Smart Configuration** - Database, cache, authentication, and DevOps integration
- 🛠️ **Modern Tooling** - Docker, Kubernetes, CI/CD, Swagger documentation
- 🎨 **Customizable** - Extend with your own templates
- 🚀 **Production Ready** - Best practices and professional structure

## 🏃‍♂️ Quick Start

### Installation

```bash
# Using Go install
go install github.com/rafa-mori/gocrafter@latest

# Or download from releases
curl -sSL https://github.com/rafa-mori/gocrafter/releases/latest/download/gocrafter-linux-amd64.tar.gz | tar xz
```

### Create Your First Project

```bash
# Interactive mode (recommended for first time)
gocrafter new

# Quick mode
gocrafter new my-api --template api-rest

# List available templates
gocrafter list

# Get template details
gocrafter info api-rest
```

## 📦 Available Templates

| Template | Description | Features |
|----------|-------------|----------|
| **api-rest** | REST API server | Gin framework, middleware, health checks, Swagger |
| **cli-tool** | Command-line application | Cobra framework, subcommands, configuration |
| **microservice** | Microservice architecture | gRPC + HTTP, service discovery, metrics |
| **grpc-service** | Pure gRPC service | Protocol buffers, streaming, service mesh ready |
| **worker** | Background job processor | Queue integration, retry mechanisms, monitoring |
| **library** | Go library/package | Documentation, testing, CI/CD workflows |

## 🎯 Example: Creating a REST API

```bash
$ gocrafter new my-blog-api --template api-rest
🚀 Starting project generation...
✅ Project generated successfully!
📁 Location: my-blog-api

Next steps:
  cd my-blog-api
  make run    # Start the application
  make test   # Run tests
  make build  # Build the application
```

**Generated project structure:**

```
my-blog-api/
├── cmd/main.go              # Application entry point
├── internal/
│   ├── config/             # Configuration management
│   ├── handler/            # HTTP handlers
│   ├── middleware/         # HTTP middleware
│   ├── model/              # Data models
│   ├── repository/         # Data access layer
│   └── service/            # Business logic
├── pkg/                    # Public packages
├── Makefile               # Build automation
├── Dockerfile             # Container configuration
├── docker-compose.yml     # Development environment
├── .env.example           # Environment template
└── README.md              # Project documentation
```

## ⚙️ Configuration Options

GoCrafter supports extensive configuration through interactive prompts:

### Database Support

- **PostgreSQL** - Production-ready with connection pooling
- **MySQL** - High-performance relational database
- **MongoDB** - Document-based NoSQL database
- **SQLite** - Embedded database for development

### Caching

- **Redis** - In-memory data structure store
- **Memcached** - High-performance caching system
- **In-Memory** - Built-in Go cache

### Authentication

- **JWT** - JSON Web Token authentication
- **OAuth2** - Third-party authentication providers
- **API Keys** - Simple API key authentication

### DevOps Integration

- **Docker** - Containerization with multi-stage builds
- **Kubernetes** - Production deployment manifests
- **CI/CD** - GitHub Actions, GitLab CI, Jenkins, Azure DevOps

## 🛠️ Advanced Usage

### Interactive Mode

```bash
$ gocrafter new
🚀 Welcome to GoCrafter - Go Project Generator!
Let's craft your perfect Go project together...

? What's your project name? my-awesome-api
? What's your Go module name? github.com/username/my-awesome-api
? What type of project do you want to create? api-rest - REST API with HTTP server
? Which database do you want to use? postgres
? Do you want to add a cache layer? redis
? Which additional features do you want to include? [Use arrows to move, space to select]
  ◯ Authentication (JWT)
  ◉ API Documentation (Swagger)
  ◉ Health Checks
  ◉ Metrics (Prometheus)
  ◯ Distributed Tracing
```

### Quick Mode

```bash
# Create API with specific features
gocrafter new blog-api \
  --template api-rest \
  --output ./projects \
  --config api-config.json

# Create CLI tool
gocrafter new my-cli \
  --template cli-tool \
  --quick

# Create microservice
gocrafter new user-service \
  --template microservice
```

### Template Information

```bash
# List all templates with descriptions
gocrafter list

# Get detailed template information
gocrafter info api-rest

# Show template structure
gocrafter info microservice --show-structure
```

## 📚 Documentation

- 📖 [**User Guide**](docs/user-guide.md) - Complete usage documentation
- 🛠️ [**Template Development**](docs/template-development.md) - Create custom templates
- 🏗️ [**Architecture**](docs/architecture.md) - How GoCrafter works
- 🎯 [**Examples**](docs/examples/) - Project examples and tutorials
- 🤝 [**Contributing**](docs/CONTRIBUTING.md) - How to contribute

## 🌍 Language Support

- [🇺🇸 English](README.md)
- [🇧🇷 Português](docs/README.pt-BR.md)

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](docs/CONTRIBUTING.md) for details.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Survey](https://github.com/AlecAivazis/survey) - Interactive prompts
- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework
- [Logrus](https://github.com/sirupsen/logrus) - Structured logging

---

<p align="center">
  Made with ❤️ by <a href="https://github.com/rafa-mori">@rafa-mori</a>
</p>

<p align="center">
  <a href="https://github.com/rafa-mori/gocrafter">⭐ Give us a star if you find GoCrafter useful!</a>
</p>
