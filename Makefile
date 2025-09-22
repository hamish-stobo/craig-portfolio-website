# Go Portfolio Website Makefile

.PHONY: deps dev build-static clean kill-port-8080

# Function to kill processes on port 8080
kill-port-8080:
	@echo "Checking for processes running on port 8080..."
	@if lsof -ti:8080 > /dev/null 2>&1; then \
		echo "Found processes running on port 8080:"; \
		lsof -ti:8080 | while read pid; do \
			echo "  - Process ID: $$pid"; \
			ps -p $$pid -o pid,ppid,cmd --no-headers 2>/dev/null || echo "    (process already terminated)"; \
		done; \
		echo "Killing processes on port 8080..."; \
		lsof -ti:8080 | xargs kill -9 2>/dev/null || true; \
		sleep 1; \
		if lsof -ti:8080 > /dev/null 2>&1; then \
			echo "Warning: Some processes may still be running on port 8080"; \
		else \
			echo "Successfully killed all processes on port 8080"; \
		fi; \
	else \
		echo "No processes found running on port 8080"; \
	fi

# Install Go dependencies
deps:
	go mod tidy
	go mod download

# Start development server
dev: kill-port-8080
	@echo "Starting development server..."
	go run main.go

# Build static site
build-static: kill-port-8080
	@echo "Starting server for static build..."
	go run main.go &
	sleep 2
	@echo "Building static site..."
	curl http://localhost:8080/build
	@echo "Stopping server..."
	pkill -f "go run main.go" || true

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
