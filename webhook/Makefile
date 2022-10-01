PROJECT := $(shell basename $(shell pwd))
BIN := ./bin
SERVER_SRC := model.go server.go
CLIENT_SRC := model.go client.go
GO_ENV := CGO_ENABLED=0

.PHONY: run build clean

all: build

run:
	@trap 'exit' INT TERM; trap 'kill 0' EXIT; \
		$(GO_ENV) go run $(SERVER_SRC) & \
		$(GO_ENV) go run $(CLIENT_SRC) & \
		wait

build:
	$(GO_ENV) go build -o $(BIN)/$(PROJECT)-server $(SERVER_SRC)
	$(GO_ENV) go build -o $(BIN)/$(PROJECT)-client $(CLIENT_SRC)

clean:
	@-rm -rf $(BIN)
	go clean
