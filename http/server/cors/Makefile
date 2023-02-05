PROJECT := $(shell basename $(shell pwd))
SRC := .
BIN := $(SRC)/bin
GO_ENV := CGO_ENABLED=0

.PHONY: all run-gin run-gin-cors run-gorilla run-gorilla-cors build clean

all: build

run: run-gin

# Run without cors support using the gin framework.
run-gin:
	$(GO_ENV) go run $(SRC) -framework=gin

# Run with cors support using the gin framework.
run-gin-cors:
	$(GO_ENV) go run $(SRC) -framework=gin -cors

# Run without cors support using the gorilla framework.
run-gorilla:
	$(GO_ENV) go run $(SRC) -framework=gorilla

# Run with cors support using the gorilla framework.
run-gorilla-cors:
	$(GO_ENV) go run $(SRC) -framework=gorilla -cors

build:
	$(GO_ENV) go build -o $(BIN)/$(PROJECT) $(SRC)

clean:
	rm -rf $(BIN)
	go clean
