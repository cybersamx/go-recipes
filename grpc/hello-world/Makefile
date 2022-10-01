PROJECT := $(shell basename $(shell pwd))
SRC := .
BIN := $(SRC)/bin
GO_ENV := CGO_ENABLED=0
GENPROTO_DIR := genproto
PROTO_FILE := hello.proto

.PHONY: install codegen run build clean

all: build

install:
	go install github.com/golang/protobuf/protoc-gen-go@latest

codegen: $(PROTO_FILE) install
	mkdir -p $(GENPROTO_DIR)
	protoc -I . $(PROTO_FILE) \
	--go_out=plugins=grpc:$(GENPROTO_DIR) --go_opt=paths=source_relative

run: codegen
	$(GO_ENV) go run $(SRC)

build: codegen
	$(GO_ENV) go build -o $(BIN)/$(PROJECT) $(SRC)

clean:
	@-rm -rf $(BIN)
	go clean
