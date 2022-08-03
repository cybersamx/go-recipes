# Standard Go Makefile

PROJECT_ROOT := $(shell pwd)
PROJECT := $(shell basename $(PROJECT_ROOT))
PROJECT_BIN := ./bin
SRC := .
GO_ENV := CGO_ENABLED=0
ROOT_DIR := $(shell git rev-parse --show-toplevel)
GQL_FILE := ./gqlgen.yml

# Colorized print
BOLD := $(shell tput bold)
RED := $(shell tput setaf 1)
BLUE := $(shell tput setaf 4)
RESET := $(shell tput sgr0)

# Set up the default target all and set up phony targets.

.PHONY: all
all: build

##@ install: Install gqlgen

.PHONY: install
install:
	@-echo "$(BOLD)$(BLUE)Installing gqlgen...$(RESET)"
	go install github.com/99designs/gqlgen@latest

##@ mod: Install modules

.PHONY: mod
mod:
	@-echo "$(BOLD)$(BLUE)Installing modules for $(PROJECT)...$(RESET)"
	$(GO_ENV) go mod tidy

##@ run: Run program

.PHONY: run
run: codegen mod
	@-echo "$(BOLD)$(BLUE)Running $(PROJECT)...$(RESET)"
	$(GO_ENV) go run $(SRC)

##@ build: Build binary

.PHONY: build
build: codegen mod
	@-echo "$(BOLD)$(BLUE)Building $(PROJECT)...$(RESET)"
	$(GO_ENV) go build -o $(PROJECT_BIN)/$(PROJECT) $(SRC)

##@ test: Test program

.PHONY: test
test: mod
	@-echo "$(BOLD)$(BLUE)Testing $(PROJECT)...$(RESET)"
	go test $(SRC)

##@ codegen: Generate code skeletal code

.PHONY: codegen
codegen: install
	@echo "$(BOLD)$(BLUE)Generating code for $(PROJECT)...$(RESET)"
	test -s $(GQL_FILE) || gqlgen init	# Generate code if gql file doesn't exist

##@ clean: Clean output files and build cache

.PHONY: clean
clean:
	@-echo "$(BOLD)$(RED)Removing build cache and bin from $(PROJECT)...$(RESET)"
	@-rm -rf $(PROJECT_BIN)
	go clean
	rm -rf graph
	rm -f gqlgen.yml server.go

##@ help: Help

.PHONY: help
help: Makefile
	@-echo " Usage:\n  make $(BLUE)<target>$(RESET)"
	@-echo
	@-sed -n 's/^##@//p' $< | column -t -s ':' | sed -e 's/[^ ]*/ &/2'
