PROJECT_ROOT := $(shell pwd)
PROJECT := $(shell basename $(PROJECT_ROOT))
PROJECT_BIN := ./bin
SRC := .
ROOT_DIR := $(shell git rev-parse --show-toplevel)

NO_FRWK_BIN := no-framework
DIG_BIN := dig
WIRE_BIN := wire
NO_FRWK_SRC := $(wildcard cmd/no-framework/*.go)
DIG_SRC := $(wildcard cmd/dig/*.go)
WIRE_SRC := cmd/wire/wire_gen.go cmd/wire/main.go
WIRE_SRC_DIR := cmd/wire

# Colorized print
BOLD := $(shell tput bold)
RED := $(shell tput setaf 1)
BLUE := $(shell tput setaf 4)
RESET := $(shell tput sgr0)

# Set up the default target all and set up phony targets.

.PHONY: all
all: build

##@ run: Run all targets

.PHONY: run
run: run-no-framework

##@ install: Install wire

.PHONY: install
install:
	go install github.com/google/wire/cmd/wire@latest

##@ mod: Install modules

.PHONY: mod
mod:
	@-echo "$(BOLD)$(BLUE)Installing modules for $(PROJECT)...$(RESET)"
	$(GO_ENV) go mod tidy

##@ run-no-framework: Run no-framework

.PHONY: run-no-framework
run-no-framework: mod
	@-echo "$(BOLD)$(BLUE)Running no framework...$(RESET)"
	go run $(NO_FRWK_SRC)

##@ run-dig: Run dig

.PHONY: run-dig
run-dig: mod
	@-echo "$(BOLD)$(BLUE)Running dig...$(RESET)"
	go run $(DIG_SRC)

##@ run-wire: Run wire

.PHONY: run-wire
run-wire: install mod
	@-echo "$(BOLD)$(BLUE)Running wire...$(RESET)"
	cd $(WIRE_SRC_DIR) && wire
	go run $(WIRE_SRC)

##@ build: Build all binaries

.PHONY: build
build: build-no-framework

##@ build-no-framework: Build no framework

.PHONY: build-no-framework
build-no-framework: mod
	@echo "$(BOLD)$(BLUE)Building no framework...$(RESET)"
	go build -o $(PROJECT_BIN)/$(NO_FRWK_BIN) $(NO_FRWK_SRC)

##@ build-dig: Build dig

.PHONY: build-dig
build-dig: mod
	@echo "$(BOLD)$(BLUE)Building dig...$(RESET)"
	go build -o $(PROJECT_BIN)/$(DIG_BIN) $(DIG_SRC)

##@ build-wire: Build wire

.PHONY: build-wire
build-wire: install mod
	@echo "$(BOLD)$(BLUE)Building wire...$(RESET)"
	cd $(WIRE_SRC_DIR)
	go build -o $(PROJECT_BIN)/$(WIRE_BIN) $(WIRE_SRC)

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
