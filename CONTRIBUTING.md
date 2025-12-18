# Contributing to Catena SDK for Go

Thank you for your interest in contributing to the Catena SDK for Go! This document provides guidelines and instructions for contributing to this project.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Workflow](#development-workflow)
- [Code Generation](#code-generation)
- [Coding Guidelines](#coding-guidelines)
- [Testing](#testing)
- [Pull Request Process](#pull-request-process)

## Code of Conduct

This project adheres to a Code of Conduct that all contributors are expected to follow. Please read [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) before contributing.

## Getting Started

### Prerequisites

Before you begin, ensure you have the following tools installed:

- **Go 1.23+**: [Download and install Go](https://golang.org/dl/)
- **Docker**: [Install Docker](https://docs.docker.com/get-docker/)
- **Node.js 20+**: [Install Node.js](https://nodejs.org/)
- **Make**: Usually pre-installed on macOS/Linux, or install via package manager

### Setting Up Your Development Environment

1. Fork the repository on GitHub
2. Clone your fork locally:

```bash
git clone https://github.com/YOUR_USERNAME/catena-sdk-go.git
cd catena-sdk-go
```

3. Add the upstream repository:

```bash
git remote add upstream https://github.com/catenaclearing/catena-sdk-go.git
```

4. Install required tools:

```bash
make tools
```

This will check for and help install:
- golangci-lint
- goimports
- Redocly CLI
- Docker (for OpenAPI Generator)

5. Fetch and generate the SDK:

```bash
make specs.fetch    # Download OpenAPI specs
make generate       # Generate Go client code
```

### Make Targets

| Target | Description |
|--------|-------------|
| `make help` | Show all available targets |
| `make tools` | Install or verify required tools |
| `make fmt` | Format Go code with gofmt and goimports |
| `make lint` | Run golangci-lint |
| `make test` | Run Go tests |
| `make specs.fetch` | Download OpenAPI specs from API |
| `make specs.combine` | Combine and validate specs |
| `make generate` | Generate Go client code |
| `make verify-clean` | Verify generated code matches committed code |
| `make ci` | Run all CI checks |
| `make clean` | Remove generated files |

## Development Workflow

### 1. Create a Branch

Create a branch for your work:

```bash
git checkout -b feature/your-feature-name
# or
git checkout -b fix/your-bug-fix
```

Use descriptive branch names:
- `feature/` for new features
- `fix/` for bug fixes
- `docs/` for documentation changes
- `chore/` for maintenance tasks

### 2. Make Your Changes

**Important Notes:**

- **Do NOT manually edit generated code** in the `gen/` directory
- Generated code is created from OpenAPI specs and will be overwritten
- If you need to change generated code, contact us to modify the source OpenAPI specs or generator configuration

For non-generated code:
- Follow Go best practices and idioms
- Add tests for new functionality
- Update documentation as needed

### 3. Format Your Code

Before committing, format your code:

```bash
make fmt
```

### 4. Run Linters

Ensure your code passes all linters:

```bash
make lint
```

### 5. Run Tests

Run the test suite:

```bash
make test
```

### 6. Commit Your Changes

Write clear, descriptive commit messages:

```bash
git add .
git commit -m "feat: add new feature X"
```

We follow [Conventional Commits](https://www.conventionalcommits.org/):

- `feat:` New feature
- `fix:` Bug fix
- `docs:` Documentation changes
- `chore:` Maintenance tasks
- `test:` Test updates
- `refactor:` Code refactoring

## Code Generation

This SDK uses OpenAPI Generator to create client code from OpenAPI specifications.

### Understanding the Generation Process

1. **Specs are fetched** from Catena API endpoints
2. **Specs are validated** using Redocly CLI
3. **Code is generated** using OpenAPI Generator (Docker)
4. **Code is formatted** with gofmt and goimports

### Regenerating the SDK

To regenerate the SDK from the latest OpenAPI specifications:

```bash
make specs.fetch    # Fetch latest specs
make generate       # Regenerate code
make ci            # Verify everything works
```

The generation process:

1. Downloads OpenAPI specs from the Catena API endpoints
2. Validates and bundles specs using Redocly CLI
3. Generates Go clients using OpenAPI Generator (pinned to v7.2.0)
4. Formats code with gofmt and goimports
5. Runs tests and linters

### Generated Code Structure

```
gen/
├── integrations/     # Integrations API client
├── orgs/            # Organizations API client
├── telematics/      # Telematics API client
└── notifications/   # Notifications API client
```

Each package contains:
- API client code
- Model definitions
- Configuration
- go.mod and go.sum

### Modifying Generation

If you need to change how code is generated:

1. Edit `internal/build/generate.sh`
2. Update generator flags or configuration
3. Regenerate and test:

```bash
make generate
make ci
```

## Coding Guidelines

### Go Code Standards

- Follow the [Effective Go](https://golang.org/doc/effective_go) guidelines
- Use `gofmt` and `goimports` for formatting
- Write clear, self-documenting code
- Add comments for exported functions and types
- Keep functions small and focused

### Linting Rules

The project uses `golangci-lint` with these enabled linters:

- errcheck
- gosimple
- govet
- ineffassign
- staticcheck
- unused
- revive
- gocritic
- bodyclose
- unconvert
- misspell
- gofmt
- goimports

Configuration: `.golangci.yml`

### Error Handling

- Always check and handle errors
- Provide context in error messages
- Don't panic in library code
- Use custom error types when appropriate

### Documentation

- Add package-level documentation
- Document all exported functions and types
- Include usage examples in godoc comments
- Keep README.md up-to-date

## Testing

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
go test -v -race -coverprofile=coverage.out ./...

# View coverage report
go tool cover -html=coverage.out
```

### Writing Tests

- Place tests in `*_test.go` files
- Use table-driven tests when possible
- Test error conditions
- Mock external dependencies
- Aim for high coverage on non-generated code

Example test structure:

```go
func TestMyFunction(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        want    string
        wantErr bool
    }{
        {
            name:    "valid input",
            input:   "test",
            want:    "expected",
            wantErr: false,
        },
        // More test cases...
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := MyFunction(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("MyFunction() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("MyFunction() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

## CI/CD

The repository uses GitHub Actions for CI/CD:

- **CI Workflow** (`.github/workflows/ci.yml`): Runs on every PR and push to main
  - Fetches specs
  - Generates code
  - Runs formatting, linting, and tests
  - Verifies generated code matches committed code

- **Regenerate Workflow** (`.github/workflows/regen.yml`): 
  - Runs weekly (Sunday at midnight UTC)
  - Can be triggered manually via workflow_dispatch
  - Fetches latest specs and regenerates SDK
  - Opens a PR if changes are detected

## Pull Request Process

### Before Submitting

1. Ensure all tests pass: `make test`
2. Run linters: `make lint`
3. Format code: `make fmt`
4. Run full CI: `make ci`
5. Update documentation if needed
6. Rebase on latest main:

```bash
git fetch upstream
git rebase upstream/main
```

### Submitting Your PR

1. Push your branch to your fork:

```bash
git push origin your-branch-name
```

2. Open a Pull Request on GitHub
3. Fill out the PR template completely
4. Link any related issues

### PR Title Format

Use Conventional Commits format:

- `feat: add new integration endpoint support`
- `fix: correct authentication header handling`
- `docs: update installation instructions`
- `chore: update dependencies`

### PR Description

Include:

- **What** changed
- **Why** it changed
- **How** to test it
- Screenshots (if UI-related)
- Related issues

### Code Review Process

1. Automated checks must pass (CI/CD)
2. At least one maintainer approval required
3. Address review comments
4. Keep the PR up-to-date with main
5. Squash commits if requested

### After Merge

1. Delete your branch
2. Update your local main:

```bash
git checkout main
git pull upstream main
```

## Questions or Need Help?

- Open an issue for bugs or feature requests
- Start a discussion for questions
- Check existing issues and discussions first

## License

By contributing, you agree that your contributions will be licensed under the Apache License 2.0.

Thank you for contributing to Catena SDK for Go! 🎉
