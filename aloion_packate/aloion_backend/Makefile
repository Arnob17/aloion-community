.PHONY: run build test clean deps migrate dev air install-air

# Run the application
run:
	go run cmd/server/main.go

# Build the application
build:
	go build -o bin/server cmd/server/main.go

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -rf bin/ tmp/

# Download dependencies
deps:
	go mod download
	go mod tidy

# Run database migrations (auto-runs on server start)
migrate:
	@echo "Migrations run automatically on server start"

# Install dependencies
install: deps

# Development mode (without Air)
dev:
	ENV=development go run cmd/server/main.go

# Development mode with Air (live reload)
# Install Air first: go install github.com/cosmtrek/air@latest
air:
	@if ! command -v air > /dev/null; then \
		echo "Air is not installed. Installing..."; \
		go install github.com/cosmtrek/air@latest; \
	fi
	air

# Install Air tool
install-air:
	go install github.com/cosmtrek/air@latest
