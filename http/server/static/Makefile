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

# Colorized print
BOLD := $(shell tput bold)
RED := $(shell tput setaf 1)
BLUE := $(shell tput setaf 4)
RESET := $(shell tput sgr0)

# Set up the default target all and set up phony targets.

.PHONY: all
all: build

##@ mod: Install modules

.PHONY: mod
mod:
	@-echo "$(BOLD)$(BLUE)Installing modules for $(PROJECT)...$(RESET)"
	$(GO_ENV) go mod tidy

##@ web-copy: Copy the static directory

.PHONY: web-copy
web-copy:
	@mkdir -p $(GIN_SRC_DIR)/$(TARGET_STATIC_WEB)
	@mkdir -p $(HTTP_SRC_DIR)/$(TARGET_STATIC_WEB)
	@cp -rf $(SRC_STATIC_WEB)/* $(GIN_SRC_DIR)/$(TARGET_STATIC_WEB)
	@cp -rf $(SRC_STATIC_WEB)/* $(HTTP_SRC_DIR)/$(TARGET_STATIC_WEB)

##@ build: Build binary

.PHONY: build
build: build-http

##@ run: Run program

.PHONY: run
run: run-http

##@ build-http: Build binary implemented with the http module

.PHONY: build-http
build-http: mod web-copy
	@-echo "$(BOLD)$(BLUE)Building $(PROJECT) (http package)...$(RESET)"
	@cd $(HTTP_SRC_DIR); $(GO_ENV) go build -o $(PROJECT_BIN)/$(PROJECT) $(SRC); cd -

##@ build-gin: Build binary implemented with the gin module

.PHONY: build-gin
build-gin: mod web-copy
	@-echo "$(BOLD)$(BLUE)Building $(PROJECT) (http package)...$(RESET)"
	@cd $(GIN_SRC_DIR); $(GO_ENV) go build -o $(PROJECT_BIN)/$(PROJECT) $(SRC); cd -

##@ run-http: Run the web server implemented with the http module

.PHONY: run-http
run-http: mod web-copy
	@-echo "$(BOLD)$(BLUE)Running $(PROJECT)...$(RESET)"
	@cd $(HTTP_SRC_DIR); $(GO_ENV) go run $(SRC); cd -

##@ run-gin: Run the web server implemented with the gin module

.PHONY: run-gin
run-gin: mod web-copy
	@-echo "$(BOLD)$(BLUE)Running $(PROJECT)...$(RESET)"
	@cd $(GIN_SRC_DIR); $(GO_ENV) go run $(SRC); cd -

##@ help: Help

.PHONY: help
help: Makefile
	@-echo " Usage:\n  make $(BLUE)<target>$(RESET)"
	@-echo
	@-sed -n 's/^##@//p' $< | column -t -s ':' | sed -e 's/[^ ]*/ &/2'
