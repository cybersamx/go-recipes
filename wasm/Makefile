PROJECT := $(shell basename $(shell pwd))
SRC := .
BIN := $(SRC)/bin
GO_ENV := CGO_ENABLED=0

# Wasm Settings
JS_BOOTSRAP := "$(shell go env GOROOT)/misc/wasm/wasm_exec.js"
BIN_NAME := hello.wasm

.PHONY: install run build clean

all: build

install:
	go install github.com/shurcooL/goexec@latest

run: install build
	go get -u github.com/shurcooL/goexec
	goexec 'http.ListenAndServe(`:8000`, http.FileServer(http.Dir(`.`)))'

build:
	GOOS=js GOARCH=wasm go build -o $(BIN)/$(BIN_NAME) $(SRC)
	@-cp $(JS_BOOTSRAP) $(BIN)

clean:
	@-rm -rf $(BIN)
	go clean
