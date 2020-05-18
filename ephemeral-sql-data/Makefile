## CREDIT: https://github.com/azer/go-makefile-example/blob/master/Makefile
##
## Instructions
## All comments that start ##@ [target] [description] will be shown in help

GO := go
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin

PROJECT := $(shell basename "$(PWD)")
SRC := $(wildcard *.go)

##@ run: Run program

run:
	@echo "Running $(PROJECT)..."
	$(GO) run $(SRC)

##@ build: Build binary

build:
	@echo "Building $(PROJECT)..."
	$(GO) build -o $(GOBIN)/$(PROJECT) $(SRC)

##@ clean: Clean output files and build cache

clean:
	@echo "Removing bin..."
	@-rm -rf $(GOBIN)
	@-$(MAKE) go-clean

go-clean:
	@echo "Cleaning build cache..."
	$(GO) clean

##@ help: Help

.PHONY: help
all: help
help: Makefile
	@echo " Usage:\n  make \033[36m<target>\033[0m"
	@echo
	@sed -n 's/^##@//p' $< | column -t -s ':' | sed -e 's/[^ ]*/ &/2'
