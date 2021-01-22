PROJECT_ROOT := $(shell pwd)
PROJECT := $(shell basename $(PROJECT_ROOT))
PROJECT_BIN := ./bin
SRC := .
GO_ENV := CGO_ENABLED=0
ROOT_DIR := $(shell git rev-parse --show-toplevel)
LINTER_CONFIG := $(ROOT_DIR)/.golangci.yaml

# Wasm Settings
JS_BOOTSRAP := "$(shell go env GOROOT)/misc/wasm/wasm_exec.js"
BIN_NAME := hello.wasm

# Colorized print
BOLD := $(shell tput bold)
RED := $(shell tput setaf 1)
BLUE := $(shell tput setaf 4)
RESET := $(shell tput sgr0)

# Set up the default target all and set up phony targets.

.PHONY: all

all: run

##@ run: Run program

.PHONY: run

run: build
	@echo "$(BOLD)$(BLUE)Running $(PROJECT)...$(RESET)"
	go get -u github.com/shurcooL/goexec
	goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'

##@ build: Build binary

.PHONY: build

build:
	@echo "$(BOLD)$(BLUE)Building $(PROJECT)...$(RESET)"
	GOOS=js GOARCH=wasm go build -o $(PROJECT_BIN)/$(BIN_NAME) $(SRC)
	@-cp $(JS_BOOTSRAP) $(PROJECT_BIN)

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
