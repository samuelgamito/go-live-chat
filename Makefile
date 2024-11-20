# Define variables
APP_NAME := go-live-chat
COVERAGE_OUT := coverage.out

# Default target
.PHONY: all
all: test

# Run all tests
.PHONY: test
test:
	go test ./... -v

# Run tests and generate HTML coverage report
.PHONY: test-coverage
test-coverage:
	go test -coverprofile=$(COVERAGE_OUT) ./...
	go tool cover -html=$(COVERAGE_OUT) -o coverage.html
	@echo "Coverage report generated at coverage.html"

# Run the application
.PHONY: run
run:
	go run ./cmd/api.go

# Clean up generated files
.PHONY: clean
clean:
	rm -f $(COVERAGE_OUT) coverage.html
