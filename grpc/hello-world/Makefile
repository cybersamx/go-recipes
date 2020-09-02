# Comments that start with ##@ [target] [description] will be shown in help

# Project paths
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
PROJECT := $(shell basename "$(PWD)")
SRC := .
GENPROTO_DIR := genproto

# Colorized print
BOLD := $(shell tput bold)
RED := $(shell tput setaf 1)
GREEN := $(shell tput setaf 2)
BLUE := $(shell tput setaf 4)
CYAN := $(shell tput setaf 6)
RESET := $(shell tput sgr0)

# Target setup

.PHONY: all clean codegen build help run

all: run

##@ run: Run program

run: codegen
	@echo "$(BOLD)$(GREEN)Running $(PROJECT)...$(RESET)"
	go run $(SRC)

##@ build: Build binary

build: codegen
	@echo "$(BOLD)$(BLUE)Building $(PROJECT)...$(RESET)"
	go build -o $(GOBIN)/$(PROJECT) $(SRC)

##@ codegen: Generate code from proto file

codegen: hello.proto
	@echo "$(BOLD)$(BLUE)Code generate $(PROJECT)...$(RESET)"
	mkdir -p $(GENPROTO_DIR)
	protoc -I . hello.proto \
	--go_out=plugins=grpc:$(GENPROTO_DIR) --go_opt=paths=source_relative

##@ clean: Clean output files and build cache

clean: go-clean
	@echo "$(BOLD)$(RED)Removing bin and $(GENPROTO_DIR)...$(RESET)"
	@-rm -rf $(GOBIN)
	@-rm -rf $(GENPROTO_DIR)

go-clean:
	@echo "$(BOLD)$(RED)Cleaning build cache...$(RESET)"
	go clean

##@ help: Help

help: Makefile
	@echo " Usage:\n  make $(CYAN)<target>$(RESET)"
	@echo
	@sed -n 's/^##@//p' $< | column -t -s ':' | sed -e 's/[^ ]*/ &/2'
