# Makefile for IoT Security Tool

# Variables
BINARY_NAME=iot-security-tool
BINARY_UNIX=$(BINARY_NAME)_unix
BINARY_WIN=$(BINARY_NAME).exe

# Go related variables
GOBASE=$(CURDIR)
GOBIN=$(GOBASE)\bin
GOFILES=$(wildcard *.go)

## build: Compile the binary.
build:
	@echo "  >  Building binary..."
	@go build -o $(BINARY_NAME).exe cmd\$(BINARY_NAME)\main.go

## run: Build and run the binary
run: build
	@echo "  >  Running binary..."
	@.\$(BINARY_NAME).exe

## run-web: Build and run the web server
run-web: build
	@echo "  >  Running web server..."
	@.\$(BINARY_NAME).exe web

## clean: Clean build files. Runs `go clean` internally.
clean:
	@echo "  >  Cleaning build cache"
	@go clean

## test: Run unit tests
test:
	@echo "  >  Running tests..."
	@go test .\...

## fmt: Run gofmt on all code
fmt:
	@echo "  >  Formatting code..."
	@go fmt .\...

## vet: Run go vet on all code
vet:
	@echo "  >  Vetting code..."
	@go vet .\...

## lint: Run golint on all code
lint:
	@echo "  >  Linting code..."
	@golint .\...

## build-all: Build for all platforms
build-all: build-linux build-windows

## build-linux: Build for Linux
build-linux:
	@echo "  >  Building binary for Linux..."
	@set GOOS=linux
	@set GOARCH=amd64
	@go build -o $(BINARY_UNIX) cmd\$(BINARY_NAME)\main.go

## build-windows: Build for Windows
build-windows:
	@echo "  >  Building binary for Windows..."
	@set GOOS=windows
	@set GOARCH=amd64
	@go build -o $(BINARY_WIN) cmd\$(BINARY_NAME)\main.go

## deps: Download dependencies
deps:
	@echo "  >  Downloading dependencies..."
	@go mod download

## help: Show this help message
help:
	@echo "Usage:"
	@findstr /r "^##" $(MAKEFILE_LIST)

.DEFAULT_GOAL := build
