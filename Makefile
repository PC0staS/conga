# Makefile for CONGA development helpers
# Not committed per user request

DOMAIN ?= conga.local
OUT_DIR ?= ./certs
BINARY ?= conga
BUILD_DIR ?= build

.PHONY: all generate-certs build-local build-all test lint clean install help

all: build-local

# Generate local certificates using tools/generate_certs.sh (mkcert required)
# Usage: make generate-certs DOMAIN=conga.local
generate-certs:
	@echo "==> Generating certificates for $(DOMAIN) into $(OUT_DIR)"
	@./tools/generate_certs.sh DOMAIN=$(DOMAIN)

# Build a local developer binary into $(BUILD_DIR)
build:
	@echo "==> Building local binary into $(BUILD_DIR)/$(BINARY)"
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY) .

# Build all release binaries using build.sh (multi-platform)
build-all:
	@echo "==> Building release binaries (multi-platform)"
	@./build.sh

# Run tests for the whole module
test:
	@echo "==> Running tests"
	@go test ./...

# Lint (no-op if tool not installed)
lint:
	@command -v golangci-lint >/dev/null 2>&1 || { echo "golangci-lint not found; skipping lint"; exit 0; }
	@golangci-lint run

# Install the local binary to GOPATH/bin or $GOBIN if set
install: build
	@echo "==> Installing binary to \
"
	@command -v go >/dev/null 2>&1 || { echo "go not found"; exit 1; }
	@cp $(BUILD_DIR)/$(BINARY) $(GOBIN)

# Clean build artifacts and generated certs
clean:
	@echo "==> Cleaning"
	@rm -f $(BUILD_DIR)/$(BINARY)
	@rm -rf $(BUILD_DIR)
	@rm -rf $(OUT_DIR)

help:
	@echo "Makefile targets:"
	@echo "  make generate-certs DOMAIN=conga.local  # generate mkcert certificates into ./certs"
	@echo "  make build                              # build a local binary into ./build/conga"
	@echo "  make build-all                          # run build.sh to produce release binaries"
	@echo "  make test                               # run go test ./..."
	@echo "  make lint                               # run golangci-lint if installed"
	@echo "  make clean                              # remove build/ and certs/"
	@echo "  make install                            # build and copy binary to GOPATH/bin or $GOBIN"
