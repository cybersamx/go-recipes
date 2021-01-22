# There's a difference between $(PWD) and $(shell pwd).
# $(PWD) = The current path where we invoke make.
# $(shell pwd) = The current path of the Makefile.
PROJECT_ROOT := $(shell pwd)
PROJECT := $(shell basename $(PROJECT_ROOT))
PROJECT_BIN := ./bin
SRC := .
GO_ENV := CGO_ENABLED=0
ROOT_DIR := $(shell git rev-parse --show-toplevel)
LINTER_CONFIG := $(ROOT_DIR)/.golangci.yaml

# Colorized print
BOLD := $(shell tput bold)
RED := $(shell tput setaf 1)
GREEN := $(shell tput setaf 2)
BLUE := $(shell tput setaf 4)
CYAN := $(shell tput setaf 6)
RESET := $(shell tput sgr0)

# Set up the default target all and set up phony targets ie. they are not files.

.PHONY: all

all: run

##@ run: Run program

.PHONY: run

# @  = Don't print command.
# @- = Don't print command and ignore error.
# @? = target
run:
	@-echo "$(BOLD)$(BLUE)Running $(PROJECT)...$(RESET)"
	$(GO_ENV) go run $(SRC)

##@ build: Build binary

.PHONY: build

build:
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

##@ clean: Clean output files and build cache

.PHONY: clean

clean:
	@-echo "$(BOLD)$(RED)Removing build cache and bin in $(PROJECT)...$(RESET)"
	@-rm -rf $(PROJECT_BIN)
	go clean

##@ help: Help

.PHONY: help

help: Makefile
	@-echo " Usage:\n  make $(BLUE)<target>$(RESET)"
	@-echo
	@-sed -n 's/^##@//p' $< | column -t -s ':' | sed -e 's/[^ ]*/ &/2'
