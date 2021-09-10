PROJECT_ROOT := $(shell pwd)
PROJECT := $(shell basename $(PROJECT_ROOT))
PROJECT_BIN := ./bin
SRC := .
HTTP_SRC_DIR := cmd/http
GIN_SRC_DIR := cmd/gin
SRC_STATIC_WEB := public
TARGET_STATIC_WEB := public
GO_ENV := CGO_ENABLED=0
ROOT_DIR := $(shell git rev-parse --show-toplevel)
LINTER_CONFIG := $(ROOT_DIR)/.golangci.yaml

# Colorized print
BOLD := $(shell tput bold)
RED := $(shell tput setaf 1)
BLUE := $(shell tput setaf 4)
RESET := $(shell tput sgr0)

# Set up the default target all and set up phony targets.

.PHONY: all

all: run

##@ run-http: Run the web server implemented with the http package

.PHONY: run-http

run-http: web-build
	@-echo "$(BOLD)$(BLUE)Running $(PROJECT)...$(RESET)"
	@cd $(HTTP_SRC_DIR); $(GO_ENV) go run $(SRC); cd -

##@ run-gin: Run the web server implemented with the gin package

.PHONY: run-gin

run-gin: web-build
	@-echo "$(BOLD)$(BLUE)Running $(PROJECT)...$(RESET)"
	@cd $(GIN_SRC_DIR); $(GO_ENV) go run $(SRC); cd -

##@ web-build: Copy the static directory

.PHONY: web-build

web-build:
	@mkdir -p $(GIN_SRC_DIR)/$(TARGET_STATIC_WEB)
	@mkdir -p $(HTTP_SRC_DIR)/$(TARGET_STATIC_WEB)
	@cp -rf $(SRC_STATIC_WEB)/* $(GIN_SRC_DIR)/$(TARGET_STATIC_WEB)
	@cp -rf $(SRC_STATIC_WEB)/* $(HTTP_SRC_DIR)/$(TARGET_STATIC_WEB)

##@ lint: Run linter

.PHONY: lint

lint:
	@-echo "$(BOLD)$(BLUE)Linting $(PROJECT)...$(RESET)"
	golangci-lint run -v -c $(LINTER_CONFIG)

##@ format: Run gofmt

.PHONY: format

format:
	@-echo "$(BOLD)$(BLUE)Formatting $(PROJECT)...$(RESET)"
	gofmt -e -s -w $(SRC)

##@ help: Help

.PHONY: help

help: Makefile
	@-echo " Usage:\n  make $(BLUE)<target>$(RESET)"
	@-echo
	@-sed -n 's/^##@//p' $< | column -t -s ':' | sed -e 's/[^ ]*/ &/2'
