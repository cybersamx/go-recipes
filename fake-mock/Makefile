PROJECT := $(shell basename $(shell pwd))
SRC := ./api

MOCK_PACKAGE_NAME := mocks
MOCK_SOURCE := ./model/account.go
MOCK_INTERFACE := AccountModel

.PHONY: install codegen test clean

all: test

install:
	go install github.com/golang/mock/mockgen@latest

codegen: $(MOCK_SOURCE)
	${GOPATH}/bin/mockgen -source=$(MOCK_SOURCE) -package=$(MOCK_PACKAGE_NAME) $(MOCK_INTERFACE) \
		> api/$(MOCK_PACKAGE_NAME)/model_mockgen.go

test: install codegen
	@-echo "$(BOLD)$(BLUE)Testing $(PROJECT)...$(RESET)"
	go test -v $(SRC)

clean:
	rm -rf $(BIN)
	go clean
