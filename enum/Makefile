PROJECT := $(shell basename $(shell pwd))
SRC := .
BIN := $(SRC)/bin
GO_ENV := CGO_ENABLED=0

.PHONY: install generate run build clean

all: build

run:
	$(GO_ENV) go run $(SRC)

install:
	go install golang.org/x/tools/cmd/stringer@latest

generate: install
	go generate

build: generate
	$(GO_ENV) go build -o $(BIN)/$(PROJECT) $(SRC)

clean:
	rm -rf $(BIN)
	go clean
