PROJECT := $(shell basename $(shell pwd))
SRC := .
BIN := $(SRC)/bin
HTTP_SRC_DIR := cmd/http
GIN_SRC_DIR := cmd/gin
SRC_STATIC_WEB := public
TARGET_STATIC_WEB := public
GO_ENV := CGO_ENABLED=0

.PHONY: all web-copy build run build-http build-gin run-http run-gin

all: build

web-copy:
	@mkdir -p $(GIN_SRC_DIR)/$(TARGET_STATIC_WEB)
	@mkdir -p $(HTTP_SRC_DIR)/$(TARGET_STATIC_WEB)
	@cp -rf $(SRC_STATIC_WEB)/* $(GIN_SRC_DIR)/$(TARGET_STATIC_WEB)
	@cp -rf $(SRC_STATIC_WEB)/* $(HTTP_SRC_DIR)/$(TARGET_STATIC_WEB)

build: build-http

run: run-http

build-http: web-copy
	@cd $(HTTP_SRC_DIR); $(GO_ENV) go build -o $(BIN)/$(PROJECT) $(SRC); cd -

build-gin: web-copy
	@cd $(GIN_SRC_DIR); $(GO_ENV) go build -o $(BIN)/$(PROJECT) $(SRC); cd -

run-http: web-copy
	@cd $(HTTP_SRC_DIR); $(GO_ENV) go run $(SRC); cd -

run-gin: web-copy
	@cd $(GIN_SRC_DIR); $(GO_ENV) go run $(SRC); cd -

clean:
	rm -rf $(BIN)
	go clean
