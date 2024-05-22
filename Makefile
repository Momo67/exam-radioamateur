# Makefile for building and running a Go program

# Name of the output binary
BINARY_NAME=random_question

# Main source file
MAIN_FILE=main.go

# Default target
.PHONY: all
all: build

# Build the Go program
.PHONY: build
build:
	go build -o bin/$(BINARY_NAME) $(MAIN_FILE)

# Run the Go program
.PHONY: run
run: build
	./bin/$(BINARY_NAME)

# Clean up generated files
.PHONY: clean
clean:
	-rm -f bin/$(BINARY_NAME)

# Install dependencies (if any)
.PHONY: deps
deps:
	go mod tidy

# Format the code
.PHONY: fmt
fmt:
	go fmt ./...

# Run tests
.PHONY: test
test:
	go test ./...

# Lint the code
.PHONY: lint
lint:
	go vet ./...
