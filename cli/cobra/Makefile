PROJECT := $(shell basename $(shell pwd))
BIN := ./bin
SRC := .
GO_ENV := CGO_ENABLED=0

.PHONY: mod run build
all: build

# Run program
run:
	$(GO_ENV) go run $(SRC)

# Build binary
.PHONY: build
build:
	$(GO_ENV) go build -o $(BIN)/$(PROJECT) $(SRC)
