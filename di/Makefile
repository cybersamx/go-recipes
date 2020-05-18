## CREDIT: https://github.com/azer/go-makefile-example/blob/master/Makefile
##
## Instructions
## All comments that start ##@ [target] [description] will be shown in help

GO := go
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin

NO_FRWK_BIN := no-framework
DIG_BIN := dig
WIRE_BIN := wire
NO_FRWK_SRC := $(wildcard cmd/no-framework/*.go)
DIG_SRC := $(wildcard cmd/dig/*.go)
WIRE_SRC := cmd/wire/wire_gen.go cmd/wire/main.go
WIRE_SRC_DIR := cmd/wire

all: install

##@ install: Install wire

install:
	$(GO) get github.com/google/wire/cmd/wire

##@ run: Run all targets

run: run-no-framework run-dig run-wire

##@ run-no-framework: Run no-framework

run-no-framework:
	@echo "Running no-framework..."
	$(GO) run $(NO_FRWK_SRC)

##@ run-dig: Run dig

run-dig:
	@echo "Running dig..."
	$(GO) run $(DIG_SRC)

##@ run-wire: Run wire

run-wire:
	@echo "Running wire..."
	cd $(WIRE_SRC_DIR) && wire
	$(GO) run $(WIRE_SRC)

##@ build: Build all binaries

build: build-no-framework build-dig build-wire

##@ build-no-framework: Build no-framework

build-no-framework:
	@echo "Building no-framework..."
	$(GO) build -o $(GOBIN)/$(NO_FRWK_BIN) $(NO_FRWK_SRC)

##@ build-dig: Build dig

build-dig:
	@echo "Building dig..."
	$(GO) build -o $(GOBIN)/$(DIG_BIN) $(DIG_SRC)

##@ build-wire: Build wire

build-wire:
	@echo "Building wire..."
	cd wire
	$(GO) build -o $(GOBIN)/$(WIRE_BIN) $(WIRE_SRC)

##@ clean: Clean output files and build cache

clean:
	@echo "Removing bin and generated code..."
	@-rm -rf $(GOBIN)
	@-rm -rf cmd/wire/wire_gen.go
	@-$(MAKE) go-clean

go-clean:
	@echo "Cleaning build cache..."
	$(GO) clean

##@ help: Help

.PHONY: help
help: Makefile
	@echo " Usage:\n  make \033[36m<target>\033[0m"
	@echo
	@sed -n 's/^##@//p' $< | column -t -s ':' | sed -e 's/[^ ]*/ &/2'
