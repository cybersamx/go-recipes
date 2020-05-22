## Instructions
## Comments that start with ##@ [target] [description] will be shown in help

GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin

TIMEOUT_BIN := timeout
DEADLINE_BIN := deadline
TIMEOUT_SRC := $(wildcard timeout/*.go)
DEADLINE_SRC := $(wildcard deadline/*.go)

##@ run: Run all targets

run: run-timeout run-deadline

##@ run-timeout: Run timeout

run-timeout:
	@echo "Running timeout..."
	go run $(TIMEOUT_SRC)

##@ run-deadline: Run deadline

run-deadline:
	@echo "Running deadline..."
	go run $(DEADLINE_SRC)

##@ build: Build all binaries

build: build-timeout build-deadline

##@ build-timeout: Build timeout

build-timeout:
	@echo "Building timeout..."
	go build -o $(GOBIN)/$(TIMEOUT_BIN) $(TIMEOUT_SRC)

##@ build-deadline: Build deadline

build-deadline:
	@echo "Building deadline..."
	go build -o $(GOBIN)/$(DEADLINE_BIN) $(DEADLINE_SRC)

##@ clean: Clean output files and build cache

clean:
	@echo "Removing bin..."
	@-rm -rf $(GOBIN)
	@-$(MAKE) go-clean

go-clean:
	@echo "Cleaning build cache..."
	go clean

##@ help: Help

.PHONY: help
all: help
help: Makefile
	@echo " Usage:\n  make \033[36m<target>\033[0m"
	@echo
	@sed -n 's/^##@//p' $< | column -t -s ':' | sed -e 's/[^ ]*/ &/2'
