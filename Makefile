.PHONY: help tools fmt lint test specs.fetch specs.combine generate verify-clean ci clean

# Go parameters
GOCMD=go
GOTEST=$(GOCMD) test
GOFMT=gofmt
GOLINT=golangci-lint

# Paths
BUILD_DIR=internal/build
SPECS_DIR=specs
GEN_DIR=gen

# Default target
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_.-]+:.*?## / {printf "  %-20s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

tools: ## Install or verify required tools
	@echo "==> Checking required tools..."
	@command -v docker >/dev/null 2>&1 || { echo "Error: docker is required but not installed"; exit 1; }
	@command -v curl >/dev/null 2>&1 || { echo "Error: curl is required but not installed"; exit 1; }
	@command -v redocly >/dev/null 2>&1 || { echo "Warning: redocly CLI not found. Install with: npm install -g @redocly/cli"; }
	@command -v $(GOLINT) >/dev/null 2>&1 || { echo "Warning: golangci-lint not found. Install from: https://golangci-lint.run/usage/install/"; }
	@command -v goimports >/dev/null 2>&1 || { echo "Installing goimports..."; $(GOCMD) install golang.org/x/tools/cmd/goimports@latest; }
	@echo "==> Tool check complete"

fmt: ## Format Go code with gofmt and goimports
	@echo "==> Formatting Go code..."
	@find . -name "*.go" -not -path "./gen/*" -exec $(GOFMT) -w {} \;
	@if command -v goimports >/dev/null 2>&1; then \
		find . -name "*.go" -not -path "./gen/*" -exec goimports -w -local github.com/catenaclearing/catena-sdk-go {} \; ; \
	fi
	@echo "==> Formatting complete"

lint: ## Run golangci-lint
	@echo "==> Running linters..."
	@if command -v $(GOLINT) >/dev/null 2>&1; then \
		$(GOLINT) run --timeout=5m; \
	else \
		echo "golangci-lint not found, skipping lint"; \
	fi

test: ## Run Go tests
	@echo "==> Running tests..."
	@$(GOTEST) -v -race -cover ./...

specs.fetch: ## Download OpenAPI specifications from remote URLs
	@echo "==> Fetching specs..."
	@chmod +x $(BUILD_DIR)/fetch-specs.sh
	@$(BUILD_DIR)/fetch-specs.sh

specs.combine: ## Combine and validate OpenAPI specifications
	@echo "==> Combining specs..."
	@chmod +x $(BUILD_DIR)/combine-specs.sh
	@$(BUILD_DIR)/combine-specs.sh

generate: specs.combine ## Generate Go client code from OpenAPI specs
	@echo "==> Generating code..."
	@chmod +x $(BUILD_DIR)/generate.sh
	@$(BUILD_DIR)/generate.sh
	@echo "==> Running go mod tidy..."
	@cd $(GEN_DIR)/integrations && $(GOCMD) mod tidy || true
	@cd $(GEN_DIR)/orgs && $(GOCMD) mod tidy || true
	@cd $(GEN_DIR)/telematics && $(GOCMD) mod tidy || true
	@cd $(GEN_DIR)/notifications && $(GOCMD) mod tidy || true
	@cd $(GEN_DIR)/authentication && $(GOCMD) mod tidy || true

verify-clean: ## Verify generated code matches committed code
	@chmod +x $(BUILD_DIR)/verify-clean.sh
	@$(BUILD_DIR)/verify-clean.sh

ci: tools specs.fetch generate fmt lint verify-clean test ## Run all CI checks (fetch, generate, format, lint, verify, test)
	@echo "==> CI checks passed!"

clean: ## Remove generated files
	@echo "==> Cleaning generated files..."
	@rm -rf $(GEN_DIR)
	@rm -rf $(SPECS_DIR)/combined
	@rm -rf pagination
	@echo "==> Clean complete"

verify: tools specs.fetch generate fmt lint test ## Verify everything
	@echo "==> Verifying clean state..."
	@git diff --exit-code || { echo "Error: Uncommitted changes found after generation"; exit 1; }
