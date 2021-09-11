PROJECT_ROOT := $(shell pwd)
PROJECT := $(shell basename $(PROJECT_ROOT))
PROJECT_BIN := ./bin
STATIC_DIR := public
SRC := .
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

all: build

##@ run: Run the web server

.PHONY: run

run:
	@-echo "$(BOLD)$(BLUE)Running $(PROJECT)...$(RESET)"
	$(GO_ENV) go run $(SRC)

##@ web-build: Copy the static directory

.PHONY: web-build

web-build:
	@mkdir -p $(PROJECT_BIN)/$(STATIC_DIR)
	@cp -rf $(STATIC_DIR)/* $(PROJECT_BIN)/$(STATIC_DIR)

##@ build: Build binary

.PHONY: build

build: web-build
	@-echo "$(BOLD)$(BLUE)Building $(PROJECT)...$(RESET)"
	$(GO_ENV) go build -o $(PROJECT_BIN)/$(PROJECT) $(SRC)

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
