PROJECT_ROOT := $(shell pwd)
PROJECT := $(shell basename $(PROJECT_ROOT))
PROJECT_BIN := ./bin
SRC := .
ROOT_DIR := $(shell git rev-parse --show-toplevel)
LINTER_CONFIG := $(ROOT_DIR)/.golangci.yaml

NO_FRWK_BIN := no-framework
DIG_BIN := dig
WIRE_BIN := wire
NO_FRWK_SRC := $(wildcard cmd/no-framework/*.go)
DIG_SRC := $(wildcard cmd/dig/*.go)
WIRE_SRC := cmd/wire/wire_gen.go cmd/wire/main.go
WIRE_SRC_DIR := cmd/wire

# Colorized print
BOLD := $(shell tput bold)
RED := $(shell tput setaf 1)
BLUE := $(shell tput setaf 4)
RESET := $(shell tput sgr0)

# Set up the default target all and set up phony targets.

.PHONY: all

all: run

##@ run: Run all targets

.PHONY: run

run: run-no-framework run-dig run-wire

##@ run-no-framework: Run no-framework

run-no-framework:
	@-echo "$(BOLD)$(BLUE)Running no framework...$(RESET)"
	go run $(NO_FRWK_SRC)

##@ run-dig: Run dig

run-dig:
	@-echo "$(BOLD)$(BLUE)Running dig...$(RESET)"
	go run $(DIG_SRC)

##@ run-wire: Run wire

run-wire:
	@-echo "$(BOLD)$(BLUE)Running wire...$(RESET)"
	cd $(WIRE_SRC_DIR) && wire
	go run $(WIRE_SRC)

##@ install: Install wire

.PHONY: install

install:
	go get github.com/google/wire/cmd/wire

##@ build: Build all binaries

build: build-no-framework build-dig build-wire

##@ build-no-framework: Build no framework

build-no-framework:
	@echo "$(BOLD)$(BLUE)Running no framework...$(RESET)"
	go build -o $(PROJECT_BIN)/$(NO_FRWK_BIN) $(NO_FRWK_SRC)

##@ build-dig: Build dig

build-dig:
	@echo "$(BOLD)$(BLUE)Running dig...$(RESET)"
	go build -o $(PROJECT_BIN)/$(DIG_BIN) $(DIG_SRC)

##@ build-wire: Build wire

build-wire:
	@echo "$(BOLD)$(BLUE)Running wire...$(RESET)"
	cd $(WIRE_SRC_DIR)
	go build -o $(PROJECT_BIN)/$(WIRE_BIN) $(WIRE_SRC)

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
