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
BLUE := $(shell tput setaf 4)
RESET := $(shell tput sgr0)

# Kafka settings
APP_NAME := 'Kafka'
PORT := 9092
RETRIES := 5
SLEEP := 3

# Set up the default target all and set up phony targets.

.PHONY: all

all: run-full

##@ run-full: Run Kafka and Zookeeper containers, and the Go program.

.PHONY: run-full

run-full:
	@-echo "$(BOLD)$(BLUE)Running $(PROJECT) in full...$(RESET)"
	@-docker-compose down > /dev/null 2>&1
	@docker-compose up > /dev/null 2>&1 &
	@printf 'Launching $(APP_NAME)'
	@for i in `seq 1 $(RETRIES)`; \
		do nc -v -w 1 localhost $(PORT) > /dev/null 2>&1 && s=0 && break || s=${?} && printf '.' && sleep $(SLEEP); \
		done; \
	echo && (exit ${s}) || echo '$(APP_NAME) failed to launch' && echo '$(APP_NAME) up'; \
	go run $(SRC)

##@ teardown: Teardown Kafa and Zookeeper containers

.PHONY: teardown

teardown:
	@-echo "$(BOLD)$(RED)Tearing down Kafka...$(RESET)"
	@docker-compose down

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
