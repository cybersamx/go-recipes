PROJECT_ROOT := $(shell pwd)
PROJECT := $(shell basename "$(PROJECT_ROOT)")
SRC := ./api
ROOT_DIR := $(shell git rev-parse --show-toplevel)

MOCK_PACKAGE_NAME := mocks
MOCK_SOURCE := ./model/account.go
MOCK_INTERFACE := AccountModel

# Colorized print
BOLD := $(shell tput bold)
RED := $(shell tput setaf 1)
BLUE := $(shell tput setaf 4)
RESET := $(shell tput sgr0)

# Set up the default target all and set up phony targets.

.PHONY: all
all: test

##@ install: Install mockgen

.PHONY: install
install:
	@-echo "$(BOLD)$(BLUE)Installing mockgen...$(RESET)"
	go install github.com/golang/mock/mockgen@latest

##@ codegen: Generate mocks

.PHONY: codegen
codegen: $(MOCK_SOURCE)
	@-echo "$(BOLD)$(BLUE)Generating mocks for $(PROJECT)...$(RESET)"
	${GOPATH}/bin/mockgen -source=$(MOCK_SOURCE) -package=$(MOCK_PACKAGE_NAME) $(MOCK_INTERFACE) \
		> api/$(MOCK_PACKAGE_NAME)/model_mockgen.go

##@ install: Install packages

.PHONY: mod
mod:
	@-echo "$(BOLD)$(BLUE)Installing mods for $(PROJECT)...$(RESET)"
	go mod tidy

##@ test: Test program

.PHONY: test
test: install codegen mod
	@-echo "$(BOLD)$(BLUE)Testing $(PROJECT)...$(RESET)"
	go test -v $(SRC)

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
