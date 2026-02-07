.PHONY: build test lint clean install release help

# Variables
BINARY_NAME=statusline
VERSION=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS=-s -w -X main.Version=$(VERSION) -X main.Commit=$(COMMIT) -X main.Date=$(DATE)

# Default target
all: build

## build: Build the binary
build:
	go build -ldflags="$(LDFLAGS)" -o $(BINARY_NAME) .

## build-all: Build for all platforms
build-all:
	GOOS=darwin GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o dist/$(BINARY_NAME)-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o dist/$(BINARY_NAME)-darwin-arm64 .
	GOOS=linux GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o dist/$(BINARY_NAME)-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o dist/$(BINARY_NAME)-linux-arm64 .
	GOOS=windows GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o dist/$(BINARY_NAME)-windows-amd64.exe .

## test: Run tests
test:
	go test -v -race -coverprofile=coverage.out ./...

## coverage: Show test coverage
coverage: test
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

## lint: Run linter
lint:
	golangci-lint run

## fmt: Format code
fmt:
	gofmt -s -w .

## clean: Clean build artifacts
clean:
	rm -f $(BINARY_NAME)
	rm -f coverage.out coverage.html
	rm -rf dist/

## install: Install to ~/.claude/statusline-go/
install: build
	mkdir -p ~/.claude/statusline-go
	cp $(BINARY_NAME) ~/.claude/statusline-go/
	@echo "Installed to ~/.claude/statusline-go/$(BINARY_NAME)"

## run: Run with sample input
run: build
	@echo '{"model":{"display_name":"Claude Sonnet 4"},"session_id":"test","workspace":{"current_dir":"$(PWD)"}}' | ./$(BINARY_NAME)

## menu: Run interactive theme menu
menu: build
	./$(BINARY_NAME) --menu

## themes: List available themes
themes: build
	./$(BINARY_NAME) --list-themes

## version: Show version
version: build
	./$(BINARY_NAME) --version

## help: Show this help
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | column -t -s ':' | sed 's/^/  /'
