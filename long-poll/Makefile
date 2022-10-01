PROJECT := $(shell basename $(shell pwd))
SRC := .
BIN := $(SRC)/bin
GO_ENV := CGO_ENABLED=0
SERVER_ENV := ENABLE_CLIENT=false
CLIENT_ENV := ENABLE_CLIENT=true
ROOT_DIR := $(shell git rev-parse --show-toplevel)

.PHONY: run build clean

all: build

run:
	@trap 'exit' INT TERM; trap 'kill 0' EXIT; \
  		$(GO_ENV) $(SERVER_ENV) go run $(SRC) & \
  		$(GO_ENV) $(CLIENT_ENV) go run $(SRC) & \
  		wait

build:
	$(GO_ENV) go build -o $(BIN)/$(PROJECT) $(SRC)

clean:
	@-rm -rf $(BIN)
	go clean
