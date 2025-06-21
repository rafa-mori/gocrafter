# Security Policy

## Supported Versions

We actively support and provide security updates for the following versions of GoCrafter:

| Version | Supported          |
| ------- | ------------------ |
| 1.x.x   | :white_check_mark: |
| < 1.0   | :x:                |

## Reporting a Vulnerability

We take the security of GoCrafter seriously. If you have discovered a security vulnerability, please report it privately.

**Please do NOT report security vulnerabilities through public GitHub issues.**

### How to Report

Please use one of these private channels:

- **GitHub Security Advisories**: [Report a vulnerability](https://github.com/rafa-mori/gocrafter/security/advisories/new)
- **Email**: <security@rafamori.dev>

Include as much detail as possible:

- A description of the vulnerability
- Steps to reproduce or proof-of-concept
- The impact and affected versions
- Suggested fix (if you have one)

### What to Expect

- **Acknowledgment**: We'll acknowledge receipt within **48 hours**
- **Initial Assessment**: Initial assessment within **5 business days**
- **Status Updates**: Updates every **7 days** until resolution
- **Resolution**: We'll work to address critical issues quickly
- **Credit**: We'll credit you in release notes (with your permission)

## Security Best Practices

### When Using GoCrafter

- **Review Generated Code**: Always review generated projects before production use
- **Keep Dependencies Updated**: Regularly update dependencies in generated projects
- **Use Environment Variables**: Store sensitive configuration in environment variables
- **Enable Security Features**: Use authentication, validation, and security middleware in templates

### Generated Project Security

GoCrafter templates include security best practices:

- Input validation and sanitization
- Secure authentication patterns (JWT, OAuth2)
- HTTPS/TLS configuration
- CORS protection
- Rate limiting
- Parameterized database queries
- Security headers

## Contact

- **Security Issues**: <security@rafamori.dev> or [GitHub Security Advisories](https://github.com/rafa-mori/gocrafter/security/advisories)
- **General Questions**: [GitHub Issues](https://github.com/rafa-mori/gocrafter/issues) or [Discussions](https://github.com/rafa-mori/gocrafter/discussions)
