# Changelog

All notable changes to GoCrafter will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Initial GoCrafter implementation
- Interactive project generation wizard
- Multiple project templates (api-rest, cli-tool, microservice)
- Template-based project scaffolding
- Comprehensive documentation and examples

### Templates

- **api-rest**: Complete REST API template with Gin framework
  - Database support (PostgreSQL, MySQL, MongoDB, SQLite)
  - Caching layer (Redis, Memcached, In-Memory)
  - Authentication (JWT, OAuth2, API Keys)
  - API documentation with Swagger
  - Docker and Kubernetes support
  - Health checks and metrics
  - Comprehensive test suite

## [1.0.0] - 2025-06-21

### Added

- 🚀 **Core Features**
  - Interactive CLI wizard with `survey` library
  - Quick project generation with templates
  - Template listing and information commands
  - Modular and extensible architecture

- 📦 **Templates**
  - `api-rest`: Production-ready REST API server
  - `cli-tool`: Command-line application framework
  - `microservice`: Microservice with gRPC support

- ⚙️ **Configuration**
  - Dynamic project configuration
  - Template metadata system
  - Feature-based conditional generation
  - Environment-specific settings

- 🛠️ **Developer Experience**
  - Comprehensive error handling and validation
  - User-friendly prompts and feedback
  - Professional project structure
  - Best practices integration

- 📚 **Documentation**
  - Complete user guide and tutorials
  - Template development guide
  - Architecture documentation
  - Contributing guidelines
  - Multi-language support (English, Portuguese)

### Technical Details

- **Go Version**: Requires Go 1.21 or higher
- **Dependencies**:
  - `github.com/AlecAivazis/survey/v2` for interactive prompts
  - `github.com/spf13/cobra` for CLI framework
  - `github.com/sirupsen/logrus` for structured logging
- **Platforms**: Cross-platform support (Linux, macOS, Windows)

### Templates Features

#### API REST Template

- **Framework**: Gin HTTP framework
- **Database**: PostgreSQL, MySQL, MongoDB, SQLite support
- **Caching**: Redis, Memcached, in-memory options
- **Authentication**: JWT, OAuth2, API key authentication
- **Documentation**: Swagger/OpenAPI integration
- **DevOps**: Docker, Kubernetes, CI/CD configurations
- **Monitoring**: Health checks, metrics, logging
- **Testing**: Unit and integration test structure

#### CLI Tool Template

- **Framework**: Cobra command framework
- **Features**: Subcommands, configuration, colored output
- **Build**: Cross-platform builds and releases
- **Documentation**: Man pages and shell completion

#### Microservice Template

- **Protocols**: gRPC and HTTP support
- **Discovery**: Service discovery integration
- **Observability**: Metrics, tracing, logging
- **Deployment**: Kubernetes-ready manifests

### Project Structure

```
gocrafter/
├── cmd/                    # CLI application entry points
├── internal/               # Private packages and business logic
│   ├── cli/               # Command implementations
│   ├── generator/         # Template engine and generation logic
│   └── prompt/            # Interactive user prompts
├── templates/             # Project templates
├── docs/                  # Comprehensive documentation
├── logger/                # Logging utilities
└── version/               # Version management
```

### Installation Methods

- **Go Install**: `go install github.com/rafa-mori/gocrafter@latest`
- **Binary Releases**: GitHub releases with pre-built binaries
- **Docker**: Container image support (planned)

### Supported Platforms

- Linux (amd64, arm64)
- macOS (amd64, arm64)
- Windows (amd64)

---

### Migration Guide

This is the initial release of GoCrafter. No migration is required.

### Breaking Changes

None - this is the first release.

### Deprecations

None - this is the first release.

### Security Updates

None - this is the first release.

---

### Contributors

Special thanks to all contributors who made this release possible:

- [@rafa-mori](https://github.com/rafa-mori) - Lead developer and architect

### Acknowledgments

- [Cobra](https://github.com/spf13/cobra) - Powerful CLI framework
- [Survey](https://github.com/AlecAivazis/survey) - Interactive prompts
- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework
- [Logrus](https://github.com/sirupsen/logrus) - Structured logging

---

## Release Process

### Versioning Strategy

GoCrafter follows [Semantic Versioning](https://semver.org/):

- **MAJOR** version for incompatible API changes
- **MINOR** version for backward-compatible functionality additions
- **PATCH** version for backward-compatible bug fixes

### Release Schedule

- **Major releases**: Quarterly (every 3 months)
- **Minor releases**: Monthly or as needed for significant features
- **Patch releases**: As needed for critical bug fixes

### Release Criteria

For each release, we ensure:

- [ ] All tests pass
- [ ] Documentation is updated
- [ ] Templates are validated
- [ ] Cross-platform builds work
- [ ] Security scan completed
- [ ] Performance benchmarks maintained

### Upgrade Path

Each release will include:

- **Changelog**: Detailed list of changes
- **Migration Guide**: Steps for upgrading (if needed)
- **Breaking Changes**: Clear documentation of any breaking changes
- **Deprecation Notices**: Advance warning of future removals

---

## Support Policy

### Long Term Support (LTS)

- **Current Version**: Full support with all features and bug fixes
- **Previous Major**: Security fixes and critical bug fixes
- **Older Versions**: Community support only

### End of Life (EOL)

Versions reach EOL 12 months after the next major release.

---

## Roadmap Preview

### Upcoming Features (v1.1.0)

- Additional templates (worker, library, grpc-service)
- Template repository system
- Plugin architecture foundation
- Enhanced CLI experience

### Future Releases

- Web-based project generator
- Template marketplace
- Cloud deployment integration
- IDE extensions

---

*For the complete roadmap and feature requests, see our [GitHub Issues](https://github.com/rafa-mori/gocrafter/issues) and [Discussions](https://github.com/rafa-mori/gocrafter/discussions).*
