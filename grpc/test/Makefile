PROJECT := $(shell basename $(shell pwd))
SRC := .
BIN := $(SRC)/bin
GO_ENV := CGO_ENABLED=0
GENPROTO_DIR := genproto
PROTO_FILE := note_service.proto

.PHONY: install codegen build test clean

all: test

install:
	go get github.com/golang/protobuf/protoc-gen-go@latest

codegen: $(PROTO_FILE) install
	mkdir -p $(GENPROTO_DIR)
	protoc -I . $(PROTO_FILE) \
	--go_out=plugins=grpc:$(GENPROTO_DIR) --go_opt=paths=source_relative

build: codegen
	$(GO_ENV) go build -o $(BIN)/$(PROJECT) $(SRC)

test: codegen
	$(GO_ENV) go test -v $(SRC)

clean:
	@-rm -rf $(BIN)
	go clean
