PROJECT := $(shell basename $(shell pwd))
SRC := .
BIN := $(SRC)/bin
GO_ENV := CGO_ENABLED=0

# Postgres Settings
CONTAINER := 'postgres'
PORT := 5432
RETRIES := 5
SLEEP := 2

.PHONY: run teardown build clean

all: build

run:
	@-docker-compose down > /dev/null 2>&1
	@docker-compose up > /dev/null 2>&1 &
	@printf 'Launching $(CONTAINER)'
	@for i in `seq 1 $(RETRIES)`; \
		do nc -v -w 1 localhost $(PORT) > /dev/null 2>&1 && s=0 && break || s=${?} && printf '.' && sleep $(SLEEP); \
		done; \
	echo && (exit ${s}) || echo '$(CONTAINER) failed to launch' && echo '$(CONTAINER) up'; \
	sleep 2; \
	$(GO_ENV) go run $(SRC)
	@docker-compose down

teardown:
	@docker-compose down

build:
	$(GO_ENV) go build -o $(BIN)/$(PROJECT) $(SRC)

clean:
	@-rm -rf $(BIN)
	go clean
