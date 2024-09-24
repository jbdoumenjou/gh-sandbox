help: ## Show this help.
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST) | column -tl 2

# Default build target
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

.PHONY: build
build:  ## Build the application.
	@echo "Building for OS: $(GOOS), Arch: $(GOARCH)"
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -a -installsuffix cgo -o gh-sandbox ./...

.PHONY: lint
lint: ## List all the linting issues.
	golangci-lint run