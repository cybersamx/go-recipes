PROJECT := $(shell basename $(shell pwd))
SRC := .
BIN := $(SRC)/bin

NO_FRWK_BIN := no-framework
DIG_BIN := dig
WIRE_BIN := wire
NO_FRWK_SRC := $(wildcard cmd/no-framework/*.go)
DIG_SRC := $(wildcard cmd/dig/*.go)
WIRE_SRC := cmd/wire/wire_gen.go cmd/wire/main.go
WIRE_SRC_DIR := cmd/wire

.PHONY: install run-no-framework run-dig run-wire build-no-framework build-dig build-wire clean

all: build-no-framework

install:
	go install github.com/google/wire/cmd/wire@latest

run-no-framework:
	go run $(NO_FRWK_SRC)

run-dig:
	go run $(DIG_SRC)

run-wire: install
	cd $(WIRE_SRC_DIR) && wire
	go run $(WIRE_SRC)

build-no-framework:
	go build -o $(BIN)/$(NO_FRWK_BIN) $(NO_FRWK_SRC)

build-dig:
	go build -o $(BIN)/$(DIG_BIN) $(DIG_SRC)

build-wire: install
	cd $(WIRE_SRC_DIR)
	go build -o $(BIN)/$(WIRE_BIN) $(WIRE_SRC)

clean:
	rm -rf $(BIN)
	go clean
