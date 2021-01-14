## Instructions
## Comments that start with ##@ [target] [description] will be shown in help

GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
PROJECT := $(shell basename "$(PWD)")
SRC := .

##@ run: Run the program

run:
	@echo "Running $(PROJECT)..."
	go run $(SRC)

##@ build: Build binary

build:
	@echo "Building $(PROJECT)..."
	go build -o $(GOBIN)/$(PROJECT) $(SRC)

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
