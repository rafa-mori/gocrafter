# Contributing to GoCrafter

Thank you for your interest in contributing to **GoCrafter**! We welcome contributions from the community and are excited to have you join us in making GoCrafter the best Go project scaffolding tool.

## Code of Conduct

By participating in this project, you agree to abide by our [Code of Conduct](../CODE_OF_CONDUCT.md). Please be respectful and professional in all interactions.

## How to Contribute

There are many ways to contribute to GoCrafter:

### 🐛 Report Bugs

- Found a bug? Open an issue with detailed information
- Include steps to reproduce, expected vs actual behavior
- Provide your Go version, OS, and GoCrafter version

### 💡 Suggest Features

- Have an idea for improvement? Open an issue with the `enhancement` label
- Describe the problem you're trying to solve
- Suggest a possible solution or approach

### 📝 Improve Documentation

- Fix typos, improve clarity, or add examples
- Translate documentation to other languages
- Create tutorials or guides

### 🔧 Submit Code Changes

- Fix bugs or implement new features
- Add new templates or improve existing ones
- Improve performance or code quality

### 🧪 Test and Review

- Test pull requests and provide feedback
- Review code changes and suggest improvements
- Help validate new features

## Getting Started

### Prerequisites

Before contributing, ensure you have:

- **Go 1.21+** installed ([Download Go](https://golang.org/dl/))
- **Git** for version control
- A **GitHub account** for pull requests

Quick Go installation (optional):

```bash
# Using go-installer (easy way)
curl -sSfL 'https://raw.githubusercontent.com/faelmori/go-installer/refs/heads/main/go.sh' | bash
```

### Setting Up Your Development Environment

1. **Fork and Clone**

   ```bash
   # Fork the repository on GitHub first
   git clone https://github.com/YOUR_USERNAME/gocrafter.git
   cd gocrafter
   
   # Add upstream remote
   git remote add upstream https://github.com/rafa-mori/gocrafter.git
   ```

2. **Install Dependencies**

   ```bash
   # Download Go modules
   go mod download
   
   # Install development tools
   go install golang.org/x/tools/cmd/goimports@latest
   go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
   ```

3. **Build and Test**

   ```bash
   # Build the project
   go build -o gocrafter .
   
   # Run tests
   go test ./...
   
   # Test the CLI
   ./gocrafter list
   ```

4. **Verify Everything Works**

   ```bash
   # Create a test project
   ./gocrafter new test-project --template api-rest
   cd test-project
   go mod tidy
   go build
   ```

## Development Workflow

### Making Changes

1. **Create a Feature Branch**

   ```bash
   git checkout -b feature/your-feature-name
   # or
   git checkout -b fix/issue-description
   ```

2. **Write Code**
   - Follow Go best practices and conventions
   - Add appropriate comments and documentation
   - Write tests for new functionality

3. **Test Your Changes**

   ```bash
   # Run all tests
   go test ./...
   
   # Run linting
   golangci-lint run
   
   # Format code
   goimports -w .
   ```

4. **Commit Changes**

   ```bash
   git add .
   git commit -m "type: brief description of changes"
   ```

   Use conventional commit types:
   - `feat:` - New features
   - `fix:` - Bug fixes
   - `docs:` - Documentation changes
   - `style:` - Code style changes
   - `refactor:` - Code refactoring
   - `test:` - Adding tests
   - `chore:` - Maintenance tasks

5. **Push and Create PR**

   ```bash
   git push origin feature/your-feature-name
   ```

   Then create a Pull Request on GitHub.

## Code Standards

### Go Style Guide

We follow the standard Go style guide:

- Use `gofmt` and `goimports` for formatting
- Follow [Effective Go](https://golang.org/doc/effective_go.html) guidelines
- Use meaningful variable and function names
- Add comments for exported functions and types
- Keep functions small and focused

### Project Structure

```
gocrafter/
├── cmd/                    # CLI entry points
├── internal/               # Private packages
│   ├── cli/               # CLI commands
│   ├── generator/         # Project generation logic
│   └── prompt/            # Interactive prompts
├── templates/             # Project templates
├── docs/                  # Documentation
└── tests/                 # Integration tests
```

### Template Development

When creating or modifying templates:

1. Follow the [Template Development Guide](template-development.md)
2. Include proper `template.json` metadata
3. Use template variables consistently
4. Test generation with various configurations
5. Document template features and usage

## Testing

### Unit Tests

- Write tests for all new functions
- Aim for good test coverage
- Use table-driven tests where appropriate
- Mock external dependencies

```bash
# Run tests with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Integration Tests

- Test CLI commands end-to-end
- Verify template generation works correctly
- Test with different configurations

```bash
# Run integration tests
go test -tags=integration ./tests/...
```

## Pull Request Guidelines

### Before Submitting

- [ ] Code follows Go style guidelines
- [ ] Tests pass (`go test ./...`)
- [ ] Linting passes (`golangci-lint run`)
- [ ] Documentation is updated
- [ ] Commit messages follow conventional format

### PR Description

Please include:

1. **What** - What changes were made?
2. **Why** - Why were these changes needed?
3. **How** - How do the changes work?
4. **Testing** - How were the changes tested?

### Example PR Template

```markdown
## Description
Brief description of the changes

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation update

## Testing
- [ ] Tests pass
- [ ] Manual testing completed
- [ ] New tests added

## Checklist
- [ ] Code follows style guidelines
- [ ] Self-review completed
- [ ] Documentation updated
```

## Adding New Templates

To add a new template:

1. Create template directory in `templates/`
2. Add `template.json` with metadata
3. Create template files with Go template syntax
4. Test template generation
5. Update documentation
6. Add to template list in CLI

See [Template Development Guide](template-development.md) for details.

## Documentation

### Writing Documentation

- Use clear, concise language
- Include code examples
- Add screenshots where helpful
- Keep documentation up-to-date with code changes

### Documentation Structure

- `README.md` - Main project overview
- `docs/user-guide.md` - Complete user documentation
- `docs/template-development.md` - Template creation guide
- `docs/CONTRIBUTING.md` - This file
- `docs/examples/` - Example projects and tutorials

## Community Guidelines

### Code of Conduct

We are committed to providing a welcoming and inclusive environment. Please:

- Be respectful and professional
- Welcome newcomers and help them succeed
- Focus on constructive feedback
- Respect different viewpoints and experiences

### Getting Help

- 📖 Check the documentation first
- 🔍 Search existing issues
- 💬 Ask questions in discussions
- 🐛 Report bugs with detailed information

### Recognition

Contributors are recognized in:

- Release notes
- Contributors section
- Special mentions for significant contributions

## Release Process

Releases are managed by maintainers:

1. Version bumping follows semantic versioning
2. Changelog is updated automatically
3. GitHub Actions handle builds and releases
4. Templates are versioned with the main project

## Questions?

If you have questions about contributing:

- Open a GitHub Discussion
- Check existing documentation
- Ask in pull request comments
- Contact maintainers directly

Thank you for contributing to GoCrafter! 🚀

2. **Document Your Changes**  
   Update the `README.md` or documentation, if necessary, to include your changes.

3. **Add Tests When Possible**  
   Ensure any new functionality is accompanied by tests.

4. **Be Clear in Issue Reports**  
   When opening an issue, be detailed and provide as much context as possible.

---

## **Where to Get Help**

If you need assistance, feel free to:

- Open an issue with the `question` tag.
- Contact me via the email or LinkedIn listed in the `README.md`.

---

## **Our Commitment**

We commit to reviewing pull requests and issues as quickly as possible. We value your contribution and appreciate the time dedicated to the project!
