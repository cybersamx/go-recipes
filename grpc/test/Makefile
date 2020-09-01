## Instructions
## Comments that start with ##@ [target] [description] will be shown in help

GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
PROJECT := $(shell basename "$(PWD)")
SRC := .
PROTO_DIR := genproto

##@ run: Run program

run: codegen
	@echo "Running $(PROJECT)..."
	go run $(SRC)

##@ build: Build binary

build: codegen
	@echo "Building $(PROJECT)..."
	go build -o $(GOBIN)/$(PROJECT) $(SRC)

##@ test: Test program

test: codegen
	@echo "Test $(PROJECT)..."
	go test $(SRC)

##@ codegen: Generate code from proto file

codegen: note_service.proto
	@echo "Code generate $(PROJECT)..."
	mkdir -p $(PROTO_DIR)
	protoc -I . note_service.proto \
	--go_out=plugins=grpc:$(PROTO_DIR) --go_opt=paths=source_relative

##@ clean: Clean output files and build cache

clean: go-clean
	@echo "Removing bin..."
	@-rm -rf $(GOBIN)

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
