## Instructions
## Comments that start with ##@ [target] [description] will be shown in help

GOBASE := $(shell pwd)
PROJECT := $(shell basename "$(PWD)")
SRC := ./api

MOCK_PACKAGE_NAME := mock_recipe
MOCK_SOURCE := ./model/account.go
MOCK_INTERFACE := AccountModel

run: install test

##@ run: Test program

test:
	@echo "Testing $(PROJECT)..."
	mockgen -source=$(MOCK_SOURCE) -package=$(MOCK_PACKAGE_NAME) $(MOCK_INTERFACE) > api/$(MOCK_PACKAGE_NAME)/model_mockgen.go
	go test -v $(SRC)

##@ install: Install mockgen

install:
	@echo "Installing mockgen..."
	GO111MODULE=on go get github.com/golang/mock/mockgen@latest

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
