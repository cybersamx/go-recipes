# Comments that start with ##@ [target] [description] will be shown in help

# Project paths
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
JS_BOOTSRAP := "$(shell go env GOROOT)/misc/wasm/wasm_exec.js"
BIN_NAME := hello.wasm
PROJECT := $(shell basename "$(PWD)")
SRC := .

# Colorized print
BOLD := $(shell tput bold)
RED := $(shell tput setaf 1)
GREEN := $(shell tput setaf 2)
BLUE := $(shell tput setaf 4)
CYAN := $(shell tput setaf 6)
RESET := $(shell tput sgr0)

# Target setup

.PHONY: all clean build help run

all: run

##@ run: Run program

run: build
	@echo "$(BOLD)$(GREEN)Running $(PROJECT)...$(RESET)"
	go get -u github.com/shurcooL/goexec
	goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'

##@ build: Build binary

build:
	@echo "$(BOLD)$(BLUE)Building $(PROJECT)...$(RESET)"
	GOOS=js GOARCH=wasm go build -o $(GOBIN)/$(BIN_NAME) $(SRC)
	@-cp $(JS_BOOTSRAP) $(GOBIN)

##@ clean: Clean output files and build cache

clean: go-clean
	@echo "$(BOLD)$(RED)Removing bin and $(GENPROTO_DIR)...$(RESET)"
	@-rm -rf $(GOBIN)

go-clean:
	@echo "$(BOLD)$(RED)Cleaning build cache...$(RESET)"
	go clean

##@ help: Help

help: Makefile
	@echo " Usage:\n  make $(CYAN)<target>$(RESET)"
	@echo
	@sed -n 's/^##@//p' $< | column -t -s ':' | sed -e 's/[^ ]*/ &/2'
