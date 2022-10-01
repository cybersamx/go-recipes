PROJECT := $(shell basename $(shell pwd))
SRC := .
BIN := $(SRC)/bin
GO_ENV := CGO_ENABLED=0

.PHONY: run build clean

all: build

run:
	$(GO_ENV) go run $(SRC)

build:
	$(GO_ENV) go build -o $(BIN)/$(PROJECT) $(SRC)

clean:
	rm -rf $(BIN)
	go clean
