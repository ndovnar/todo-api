-include .makerc
SHELL = /bin/bash
PACKAGE ?= todo

#Output purposes
OUTPUT_DIR = $(CURDIR)/output
BIN_OUTPUT_DIR = $(OUTPUT_DIR)/bin
TEST_OUTPUT_DIR = $(OUTPUT_DIR)/test
DIRS=$(BIN_OUTPUT_DIR) $(TEST_OUTPUT_DIR)
$(shell mkdir -p $(DIRS))

#Build flags
# MAIN_PACKAGE: "main" package in Go, leave empty if there is no main package (for a lib)
MAIN_PACKAGE = ./cmd/app
LDFLAGS ?= "-X 'main.application=$(PACKAGE)'"
# EXTRA_BUILD_FLAGS and EXTRA_LINT_FLAGS can be used to make a difference between Dockerfile builds and local builds
BUILD_FLAGS ?= $(EXTRA_BUILD_FLAGS) $(MAIN_PACKAGE)
TEST_FLAGS ?= "-tags=unit"
PACKAGE_EXTENSION ?= $(shell if [ "$(GOOS)" = windows ]; then echo .exe; fi)
CGO_ENABLED=0

.NOTPARALLEL: ; # wait for this target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell
.PHONY: version build docs test scripts api cmd configs examples

build: ## Build the app
	go build -o $(BIN_OUTPUT_DIR)/$(PACKAGE)$(PACKAGE_EXTENSION) $(BUILD_FLAGS)

run: ## Run the app
	$(BIN_OUTPUT_DIR)/$(PACKAGE)$(PACKAGE_EXTENSION)

dev:
	go run cmd/app/main.go ${LDFLAGS}

test: ## Run tests
	go test $(TEST_FLAGS) ./... $(BUILD_FLAGS)