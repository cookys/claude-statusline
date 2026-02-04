# Contributing to Claude Statusline

Thank you for your interest in contributing to Claude Statusline!

## How to Contribute

### Reporting Bugs

1. Check if the issue already exists in [GitHub Issues](https://github.com/kevinlincg/claude-statusline/issues)
2. If not, create a new issue with:
   - Clear description of the problem
   - Steps to reproduce
   - Expected vs actual behavior
   - Your environment (OS, Go version, Claude Code version)

### Suggesting Features

1. Open a [GitHub Issue](https://github.com/kevinlincg/claude-statusline/issues) with the "enhancement" label
2. Describe the feature and its use case
3. If possible, include mockups or examples

### Pull Requests

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/my-feature`
3. Make your changes
4. Ensure code passes linting: `golangci-lint run`
5. Run tests: `go test ./...`
6. Commit with clear messages
7. Push and create a Pull Request

### Code Style

- Follow standard Go conventions
- Run `gofmt -s` before committing
- Keep functions focused and well-documented
- Add tests for new functionality

### Creating Themes

Want to add a new theme? Check the existing themes in `themes/` directory for examples. Each theme must implement the `Theme` interface:

```go
type Theme interface {
    Name() string
    Description() string
    Render(data StatusData) string
}
```

## Development Setup

```bash
# Clone the repo
git clone https://github.com/kevinlincg/claude-statusline.git
cd claude-statusline

# Build
go build -o statusline .

# Run tests
go test ./...

# Run linter
golangci-lint run
```

## Questions?

Feel free to open an issue for any questions!
