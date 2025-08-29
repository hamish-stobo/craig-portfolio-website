# Go Portfolio Website Makefile

.PHONY: deps dev build-static clean

# Install Go dependencies
deps:
	go mod tidy
	go mod download

# Start development server
dev:
	go run main.go

# Build static site
build-static:
	go run main.go &
	sleep 2
	curl http://localhost:8080/build
	pkill -f "go run main.go"

# Clean build artifacts
clean:
	rm -rf dist/
	go clean

# Format code
fmt:
	go fmt ./...

# Run linter (requires golangci-lint)
lint:
	golangci-lint run

# Test
test:
	go test ./...

# Build binary
build:
	go build -o craig-portfolio main.go

# Full build process
all: clean deps fmt build-static
