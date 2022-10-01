PROJECT := $(shell basename $(shell pwd))
SRC := .
BIN := $(SRC)/bin
GO_ENV := CGO_ENABLED=0
GQL_FILE := $(SRC)/gqlgen.yml

.PHONY: install mod run build codegen clean

all: build

install:
	go install github.com/99designs/gqlgen@latest

mod:
	go mod tidy

run: codegen mod
	$(GO_ENV) go run $(SRC)

build: codegen mod
	$(GO_ENV) go build -o $(BIN)/$(PROJECT) $(SRC)

codegen: install
	test -s $(GQL_FILE) || gqlgen init	# Generate code only if gql file doesn't exist

clean:
	@-rm -rf $(BIN)
	go clean
	rm -rf graph
	rm -f gqlgen.yml server.go
