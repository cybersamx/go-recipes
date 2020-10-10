# Comments that start with ##@ [target] [description] will be shown in help

# Project paths
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
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

.PHONY: all clean build help run test

all: test

##@ run: Run program

run:
	@echo "$(BOLD)$(GREEN)Running $(PROJECT)...$(RESET)"
	go run $(SRC)

##@ build: Build binary

build:
	@echo "$(BOLD)$(BLUE)Building $(PROJECT)...$(RESET)"
	go build -o $(GOBIN)/$(PROJECT) $(SRC)

##@ test: Test program

test:
	@echo "$(BOLD)$(BLUE)Test $(PROJECT)...$(RESET)"
	go test $(SRC)

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
