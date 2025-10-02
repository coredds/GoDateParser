# Contributing to GoDateParser

Thank you for your interest in contributing to GoDateParser! We welcome contributions from the community.

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/YOUR-USERNAME/GoDateParser.git`
3. Create a feature branch: `git checkout -b feature/your-feature-name`
4. Make your changes
5. Run tests: `go test ./...`
6. Run linter: `golangci-lint run ./...`
7. Commit your changes: `git commit -am 'Add new feature'`
8. Push to your fork: `git push origin feature/your-feature-name`
9. Create a Pull Request

## Development Guidelines

### Code Style

- Follow Go conventions and best practices
- Use `gofmt` to format your code
- Write clear, descriptive comments
- Keep functions focused and concise

### Testing

- Write tests for all new features
- Ensure all existing tests pass
- Aim for high test coverage
- Include edge cases in your tests

### Commit Messages

Use clear, descriptive commit messages:
- Start with a verb (Add, Fix, Update, Remove, etc.)
- Keep the first line under 72 characters
- Add detailed description if needed

### Pull Requests

- Describe what your PR does
- Reference any related issues
- Ensure CI passes before requesting review
- Be responsive to feedback

## Adding Language Support

To add support for a new language:

1. Create a new file in `translations/` (e.g., `french.go`)
2. Implement the `NewXxxTranslation()` function
3. Register it in `translations/registry.go`
4. Create comprehensive tests in `xxx_test.go`
5. Update the README and documentation

See `translations/portuguese.go` and `portuguese_test.go` as examples.

## Reporting Issues

- Use the issue tracker for bug reports and feature requests
- Search existing issues before creating a new one
- Provide detailed information and steps to reproduce
- Include Go version and OS information

## Questions?

Feel free to open an issue for any questions about contributing.

Thank you for helping make GoDateParser better!

