# Makefile for KVSQL Database

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Binary names
BINARY_NAME=sqld
BINARY_UNIX=$(BINARY_NAME)_unix

# Build targets
.PHONY: all build clean test coverage deps help

all: test build

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/sqld

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v ./cmd/sqld

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

test: 
	$(GOTEST) -v ./...

coverage:
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out

deps:
	$(GOMOD) download
	$(GOMOD) tidy

# Run targets
run-cli:
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/sqld && ./$(BINARY_NAME) -mode=cli

run-server:
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/sqld && ./$(BINARY_NAME) -mode=server -port=8080

# Development targets
fmt:
	$(GOCMD) fmt ./...

lint:
	golangci-lint run

vet:
	$(GOCMD) vet ./...

# Docker targets (future)
docker-build:
	@echo "TODO: Add Docker build"

docker-run:
	@echo "TODO: Add Docker run"

# Help
help:
	@echo "Available targets:"
	@echo "  build      - Build the binary"
	@echo "  clean      - Clean build artifacts"  
	@echo "  test       - Run tests"
	@echo "  coverage   - Generate test coverage report"
	@echo "  deps       - Download and tidy dependencies"
	@echo "  run-cli    - Build and run in CLI mode"
	@echo "  run-server - Build and run in server mode"
	@echo "  fmt        - Format code"
	@echo "  lint       - Run linter"
	@echo "  vet        - Run go vet"