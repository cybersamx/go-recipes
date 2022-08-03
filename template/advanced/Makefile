PROJECT_ROOT := $(shell pwd)
PROJECT := $(shell basename $(PROJECT_ROOT))
PROJECT_BIN := ./bin
SRC := .
GO_ENV := CGO_ENABLED=0
ROOT_DIR := $(shell git rev-parse --show-toplevel)

# Colorized print
BOLD := $(shell tput bold)
RED := $(shell tput setaf 1)
BLUE := $(shell tput setaf 4)
RESET := $(shell tput sgr0)

# Set up the default target all and set up phony targets.

.PHONY: all
all: build

##@ mod: Install modules

.PHONY: mod
mod:
	@-echo "$(BOLD)$(BLUE)Installing modules for $(PROJECT)...$(RESET)"
	$(GO_ENV) go mod tidy

##@ web-build: Build the web application.

.PHONY: web-build
web-build:
	@-echo "$(BOLD)$(BLUE)Building web application...$(RESET)"
	@cd public && \
	npm install && \
	npm run build && \
	cp dist/bundle.js dist/bundle.css . && \
	cd -

##@ run: Run program

.PHONY: run
run: mod web-build
	@-echo "$(BOLD)$(BLUE)Running $(PROJECT)...$(RESET)"
	$(GO_ENV) go run $(SRC)

##@ build: Build binary

.PHONY: build
build: mod web-build
	@-echo "$(BOLD)$(BLUE)Building $(PROJECT)...$(RESET)"
	$(GO_ENV) go build -o $(PROJECT_BIN)/$(PROJECT) $(SRC)
	@mkdir -p $(PROJECT_BIN)/public
	@cp public/index.html public/bundle.js public/bundle.css $(PROJECT_BIN)/public

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
