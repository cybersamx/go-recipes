## Instructions
## Comments that start with ##@ [target] [description] will be shown in help

GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
PROJECT := $(shell basename "$(PWD)")
SRC := ./cmd

##@ run-http: Run program built using the http package

run-http:
	@echo "Running $(PROJECT) (http)..."
	go run $(SRC)/http

##@ run-gin: Run program built using the gin-gonic/gin package

run-gin:
	@echo "Running $(PROJECT) (gin)..."
	go run $(SRC)/gin

##@ build-http: Build binary using the http package

build-http:
	@echo "Building $(PROJECT) (http)..."
	go build -o $(GOBIN)/$(PROJECT)-http $(SRC)/http

##@ build-gin: Build binary using the gin package

build-gin:
	@echo "Building $(PROJECT) (gin)..."
	go build -o $(GOBIN)/$(PROJECT)-gin $(SRC)/http

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
