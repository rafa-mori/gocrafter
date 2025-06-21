# GoCrafter Architecture

This document provides an overview of GoCrafter's architecture, design decisions, and internal workings.

## Overview

GoCrafter is a CLI-based project scaffolding tool designed with modularity, extensibility, and ease of use in mind. The architecture follows Go best practices and clean architecture principles.

## Project Structure

```plain
gocrafter/
├── cmd/                        # CLI entry points
│   ├── main.go                # Application entry point
│   ├── usage.go               # Usage documentation
│   └── wrpr.go                # Wrapper logic
├── internal/                   # Private packages
│   ├── cli/                   # CLI command implementations
│   │   ├── commands.go        # Command definitions
│   │   ├── new.go             # Project creation command
│   │   └── list.go            # Template listing command
│   ├── generator/             # Core generation logic
│   │   ├── config.go          # Configuration management
│   │   └── generator.go       # Template engine
│   └── prompt/                # Interactive prompts
│       └── interactive.go     # Survey-based prompts
├── templates/                  # Project templates
│   ├── api-rest/              # REST API template
│   ├── cli-tool/              # CLI tool template
│   └── microservice/          # Microservice template
├── logger/                     # Logging utilities
├── version/                    # Version management
└── docs/                       # Documentation
```

## Core Components

### 1. CLI Layer (`cmd/`)

The CLI layer serves as the user interface and application entry point.

**Key Components:**

- `main.go` - Application bootstrap and configuration
- `wrpr.go` - Command routing and execution wrapper
- `usage.go` - Help and usage information

**Responsibilities:**

- Parse command-line arguments
- Route commands to appropriate handlers
- Handle global flags and configuration
- Provide user feedback and error messages

### 2. Command Layer (`internal/cli/`)

Implements specific CLI commands and their business logic.

**Key Components:**

- `commands.go` - Command definitions and registration
- `new.go` - Project creation workflow
- `list.go` - Template listing and information

**Design Patterns:**

- Command Pattern for CLI operations
- Factory Pattern for command creation
- Strategy Pattern for different command types

### 3. Generation Engine (`internal/generator/`)

The core template processing and project generation engine.

**Key Components:**

- `generator.go` - Main template engine
- `config.go` - Project configuration model

**Features:**

- Go template engine integration
- Dynamic content generation
- File and directory structure creation
- Variable substitution and processing

### 4. Interactive System (`internal/prompt/`)

Handles user interaction and input collection.

**Key Components:**

- `interactive.go` - Survey-based prompts

**Features:**

- Multi-step interactive wizards
- Input validation and sanitization
- Conditional prompting logic
- User experience optimization

### 5. Template System (`templates/`)

Modular template definitions for different project types.

**Structure:**

```
templates/
└── template-name/
    ├── template.json          # Template metadata
    ├── go.mod.tmpl           # Go module template
    ├── main.go.tmpl          # Main application template
    └── internal/             # Template directory structure
        └── config/
            └── config.go.tmpl
```

## Design Principles

### 1. Modularity

- **Separation of Concerns**: Each package has a single, well-defined responsibility
- **Loose Coupling**: Components interact through well-defined interfaces
- **High Cohesion**: Related functionality is grouped together

### 2. Extensibility

- **Template System**: Easy to add new project templates
- **Plugin Architecture**: Designed for future plugin support
- **Configuration Driven**: Behavior controlled through configuration

### 3. User Experience

- **Interactive Design**: Guided workflows for beginners
- **Power User Support**: Quick commands for experienced users
- **Consistent Interface**: Uniform command structure and feedback

### 4. Reliability

- **Error Handling**: Comprehensive error checking and recovery
- **Input Validation**: Robust validation of user input
- **Safe Operations**: Non-destructive operations with confirmation

## Data Flow

### 1. Project Creation Flow

```plain
User Input → CLI Parser → Command Router → Interactive Prompts 
    ↓
Configuration Builder → Template Engine → File Generator → Project Output
```

### 2. Template Processing Flow

```plain
Template Selection → Metadata Loading → Variable Collection 
    ↓
Template Parsing → Content Generation → File System Operations
```

## Template Engine

### Template Structure

Templates use Go's `text/template` package with custom functions:

```go
// Template variables
type TemplateData struct {
    ProjectName   string
    ModuleName    string
    Database      string
    Features      []string
    // ... other fields
}
```

### Template Functions

Custom template functions provide additional functionality:

- `title` - Convert to title case
- `lower` - Convert to lowercase
- `contains` - Check if slice contains value
- `join` - Join strings with separator

### Conditional Generation

Templates support conditional content generation:

```go
{{if .Features.Contains "database"}}
// Database-related code
{{end}}
```

## Configuration Management

### Project Configuration

```go
type ProjectConfig struct {
    Name         string              `json:"name"`
    ModuleName   string              `json:"module_name"`
    Template     string              `json:"template"`
    Database     string              `json:"database,omitempty"`
    Features     []string            `json:"features,omitempty"`
    Additional   map[string]string   `json:"additional,omitempty"`
}
```

### Template Metadata

```json
{
    "name": "api-rest",
    "description": "REST API with HTTP server",
    "version": "1.0.0",
    "author": "GoCrafter Team",
    "tags": ["api", "rest", "http", "server"],
    "features": ["database", "cache", "auth", "swagger"],
    "dependencies": {
        "go_version": ">=1.21"
    }
}
```

## Error Handling Strategy

### 1. Error Types

- **User Errors**: Invalid input, missing requirements
- **System Errors**: File system issues, permission problems
- **Template Errors**: Invalid templates, parsing failures

### 2. Error Recovery

- **Graceful Degradation**: Continue when possible
- **Clear Messages**: User-friendly error descriptions
- **Suggestions**: Provide actionable next steps

### 3. Logging

- **Structured Logging**: Using logrus for consistent log format
- **Log Levels**: Debug, Info, Warning, Error
- **Context**: Include relevant context in log messages

## Performance Considerations

### 1. Template Caching

- Templates are parsed once and cached
- Metadata is loaded on demand
- File operations are batched when possible

### 2. Memory Management

- Stream processing for large files
- Cleanup of temporary resources
- Efficient string operations

### 3. Concurrency

- Parallel file operations where safe
- Async operations for independent tasks
- Resource pooling for expensive operations

## Security Considerations

### 1. Input Validation

- Path traversal prevention
- Input sanitization
- File name validation

### 2. Template Security

- Restricted template functions
- Safe variable substitution
- Controlled file system access

### 3. Dependency Management

- Vetted dependencies only
- Regular security updates
- Minimal dependency surface

## Testing Strategy

### 1. Unit Tests

- Individual component testing
- Mock external dependencies
- Edge case coverage

### 2. Integration Tests

- End-to-end command testing
- Template generation validation
- File system operation testing

### 3. Template Tests

- Template parsing validation
- Generated project verification
- Cross-platform compatibility

## Future Architecture Considerations

### 1. Plugin System

- Dynamic plugin loading
- Plugin API definition
- Security and isolation

### 2. Remote Templates

- Template repository support
- Version management
- Caching and updates

### 3. GUI Interface

- Web-based interface
- Desktop application
- Mobile companion app

### 4. Cloud Integration

- Cloud project deployment
- CI/CD pipeline integration
- Remote collaboration features

## Monitoring and Observability

### 1. Metrics

- Command usage statistics
- Template popularity
- Performance metrics

### 2. Health Checks

- System dependency validation
- Template integrity checks
- Resource availability monitoring

### 3. Telemetry

- Anonymous usage data
- Error reporting
- Performance profiling

## Deployment Architecture

### 1. Distribution

- Binary releases for multiple platforms
- Package manager integration
- Container images

### 2. Updates

- Version checking
- Automatic updates
- Rollback capability

### 3. Configuration

- User configuration files
- System-wide defaults
- Environment-specific settings

This architecture provides a solid foundation for GoCrafter's current functionality while allowing for future growth and enhancement.
