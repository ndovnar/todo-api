-include .makerc
SHELL = /bin/bash
PACKAGE ?= todoauth

OUTPUT_DIR = $(CURDIR)/output
BIN_OUTPUT_DIR = $(OUTPUT_DIR)/bin
TEST_OUTPUT_DIR = $(OUTPUT_DIR)/test
DIRS=$(BIN_OUTPUT_DIR) $(TEST_OUTPUT_DIR)
$(shell mkdir -p $(DIRS))

MAIN_PACKAGE = ./cmd/app
LDFLAGS ?= "-X 'main.application=$(PACKAGE)'"

BUILD_FLAGS ?= $(MAIN_PACKAGE)
TEST_FLAGS ?= "-tags=unit"
PACKAGE_EXTENSION ?= $(shell if [ "$(GOOS)" = windows ]; then echo .exe; fi)
CGO_ENABLED=0

.NOTPARALLEL: ; # wait for this target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell
.PHONY: version build docs test scripts api cmd configs examples

build:
	go build -o $(BIN_OUTPUT_DIR)/$(PACKAGE)$(PACKAGE_EXTENSION) $(BUILD_FLAGS)

run:
	$(BIN_OUTPUT_DIR)/$(PACKAGE)$(PACKAGE_EXTENSION)

dev:
	go run cmd/app/main.go ${LDFLAGS}

test:
	go test $(TEST_FLAGS) ./... $(BUILD_FLAGS)