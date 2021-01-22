PROJECT_ROOT := $(shell pwd)
PROJECT := $(shell basename $(PROJECT_ROOT))
PROJECT_BIN := ./bin
SRC := .
GO_ENV := CGO_ENABLED=0
ROOT_DIR := $(shell git rev-parse --show-toplevel)
LINTER_CONFIG := $(ROOT_DIR)/.golangci.yaml
GENPROTO_DIR := genproto

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

##@ run: Run program

.PHONY: run

run: codegen
	@-echo "$(BOLD)$(BLUE)Running $(PROJECT)...$(RESET)"
	$(GO_ENV) go run $(SRC)

##@ build: Build binary

.PHONY: build

build: codegen
	@echo "$(BOLD)$(BLUE)Building $(PROJECT)...$(RESET)"
	$(GO_ENV) go build -o $(PROJECT_BIN)/$(PROJECT) $(SRC)

##@ run: Test program

.PHONY: test

test:
	@-echo "$(BOLD)$(BLUE)Testing $(PROJECT)...$(RESET)"
	$(GO_ENV) go test -v $(SRC)

##@ codegen: Generate code from proto file

.PHONY: codegen

codegen: hello.proto
	@echo "$(BOLD)$(BLUE)Code generate $(PROJECT)...$(RESET)"
	mkdir -p $(GENPROTO_DIR)
	protoc -I . hello.proto \
	--go_out=plugins=grpc:$(GENPROTO_DIR) --go_opt=paths=source_relative

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
