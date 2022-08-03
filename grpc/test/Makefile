PROJECT_ROOT := $(shell pwd)
PROJECT := $(shell basename $(PROJECT_ROOT))
PROJECT_BIN := ./bin
SRC := .
GO_ENV := CGO_ENABLED=0
ROOT_DIR := $(shell git rev-parse --show-toplevel)
GENPROTO_DIR := genproto
PROTO_FILE := note_service.proto

# Colorized print
BOLD := $(shell tput bold)
RED := $(shell tput setaf 1)
BLUE := $(shell tput setaf 4)
RESET := $(shell tput sgr0)

# Set up the default target all and set up phony targets.

.PHONY: all
all: test

##@ install: Install protoc-gen-go

.PHONY: install
install:
	@-echo "$(BOLD)$(BLUE)Installing protoc-gen-go...$(RESET)"
	go get github.com/golang/protobuf/protoc-gen-go@latest

##@ mod: Install modules

.PHONY: mod
mod:
	@-echo "$(BOLD)$(BLUE)Installing modules for $(PROJECT)...$(RESET)"
	$(GO_ENV) go mod tidy

##@ codegen: Generate code from proto file

.PHONY: codegen
codegen: $(PROTO_FILE) install
	@echo "$(BOLD)$(BLUE)Generating code for $(PROJECT)...$(RESET)"
	mkdir -p $(GENPROTO_DIR)
	protoc -I . $(PROTO_FILE) \
	--go_out=plugins=grpc:$(GENPROTO_DIR) --go_opt=paths=source_relative

##@ build: Build binary

.PHONY: build
build: mod codegen
	@echo "$(BOLD)$(BLUE)Building $(PROJECT)...$(RESET)"
	$(GO_ENV) go build -o $(PROJECT_BIN)/$(PROJECT) $(SRC)

##@ run: Test program

.PHONY: test
test: mod codegen
	@-echo "$(BOLD)$(BLUE)Testing $(PROJECT)...$(RESET)"
	$(GO_ENV) go test -v $(SRC)

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
