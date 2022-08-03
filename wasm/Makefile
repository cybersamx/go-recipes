PROJECT_ROOT := $(shell pwd)
PROJECT := $(shell basename $(PROJECT_ROOT))
PROJECT_BIN := ./bin
SRC := .
GO_ENV := CGO_ENABLED=0
ROOT_DIR := $(shell git rev-parse --show-toplevel)

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
all: build

##@ install: Install goexec

.PHONY: install
install:
	@-echo "$(BOLD)$(BLUE)Installing goexec...$(RESET)"
	go install github.com/shurcooL/goexec@latest

##@ mod: Install modules

.PHONY: mod
mod:
	@-echo "$(BOLD)$(BLUE)Installing modules for $(PROJECT)...$(RESET)"
	$(GO_ENV) go mod tidy

##@ run: Run program

.PHONY: run
run: install build
	@echo "$(BOLD)$(BLUE)Running $(PROJECT)...$(RESET)"
	go get -u github.com/shurcooL/goexec
	goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'

##@ build: Build binary

.PHONY: build
build: mod
	@echo "$(BOLD)$(BLUE)Building $(PROJECT)...$(RESET)"
	GOOS=js GOARCH=wasm go build -o $(PROJECT_BIN)/$(BIN_NAME) $(SRC)
	@-cp $(JS_BOOTSRAP) $(PROJECT_BIN)

##@ clean: Clean output files and build cache

.PHONY: clean
clean:
	@-echo "$(BOLD)$(RED)Removing build cache and bin from $(PROJECT)...$(RESET)"
	@-rm -rf $(PROJECT_BIN)
	go clean

##@ help: Help

.PHONY: help
help: Makefile
	@-echo " Usage:\n  make $(BLUE)<target>$(RESET)"
	@-echo
	@-sed -n 's/^##@//p' $< | column -t -s ':' | sed -e 's/[^ ]*/ &/2'
