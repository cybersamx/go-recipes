## Instructions
## Comments that start with ##@ [target] [description] will be shown in help

GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
PROJECT := $(shell basename "$(PWD)")
SRC := main.go

APP_NAME := 'Pulsar'
PORT := 6650
RETRIES := 5
SLEEP := 3

##@ run-full: Run Kafka and Zookeeper containers, and the Go program

run-full:
	@echo "Running $(PROJECT) in full..."
	@-docker-compose down > /dev/null 2>&1
	@docker-compose up pulsar > /dev/null 2>&1 &
	@printf 'Launching $(APP_NAME)'
	@for i in `seq 1 $(RETRIES)`; \
		do nc -v -w 1 localhost $(PORT) > /dev/null 2>&1 && s=0 && break || s=${?} && printf '.' && sleep $(SLEEP); \
		done; \
	echo && (exit ${s}) || echo '$(APP_NAME) failed to launch' && echo '$(APP_NAME) up'; \
	go run $(SRC)

##@ teardown: Teardown Kafa and Zookeeper containers

teardown:
	@docker-compose down

##@ run: Run the program

run:
	@echo "Running $(PROJECT)..."
	go run $(SRC)

##@ build: Build binary

build:
	@echo "Building $(PROJECT)..."
	go build -o $(GOBIN)/$(PROJECT) $(SRC)

##@ clean: Clean output files and build cache

clean:
	@echo "Removing bin..."
	@-rm -rf $(GOBIN)
	@-$(MAKE) go-clean

go-clean:
	@echo "Cleaning build cache..."
	go clean

##@ help: Help

.PHONY: help
all: help
help: Makefile
	@echo " Usage:\n  make \033[36m<target>\033[0m"
	@echo
	@sed -n 's/^##@//p' $< | column -t -s ':' | sed -e 's/[^ ]*/ &/2'
