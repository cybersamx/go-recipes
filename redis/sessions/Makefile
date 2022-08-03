PROJECT_ROOT := $(shell pwd)
PROJECT := $(shell basename $(PROJECT_ROOT))
PROJECT_BIN := ./bin
SRC := .
GO_ENV := CGO_ENABLED=0
ROOT_DIR := $(shell git rev-parse --show-toplevel)

# Colorized print
BOLD := $(shell tput bold)
RED := $(shell tput setaf 1)
BLUE := $(shell tput setaf 4)
RESET := $(shell tput sgr0)

# Redis Settings
CONTAINER := 'redis'
PORT := 6379
RETRIES := 5
SLEEP := 2

# Set up the default target all and set up phony targets.

.PHONY: all
all: build

##@ mod: Install modules

.PHONY: mod
mod:
	@-echo "$(BOLD)$(BLUE)Installing modules for $(PROJECT)...$(RESET)"
	$(GO_ENV) go mod tidy

##@ run: Run program and redis container

.PHONY: run
run: mod
	@-echo "$(BOLD)$(BLUE)Running $(PROJECT)...$(RESET)"
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

##@ teardown: Teardown redis container

.PHONY: teardown
teardown:
	@-echo "$(BOLD)$(RED)Tearing down $(CONTAINER)...$(RESET)"
	@docker-compose down

##@ build: Build binary

.PHONY: build
build: mod
	@-echo "$(BOLD)$(BLUE)Building $(PROJECT)...$(RESET)"
	$(GO_ENV) go build -o $(PROJECT_BIN)/$(PROJECT) $(SRC)

##@ clean: Clean output files and build cache

.PHONY: clean
clean:
	@-echo "$(BOLD)$(RED)Removing build cache and bin from $(PROJECT)...$(RESET)"
	@-rm -rf $(PROJECT_BIN)
	go clean

##@ help: Help

.PHONY: help
help: Makefile
	@-echo " Usage:\n  make $(BLUE)<target>$(RESET)"
	@-echo
	@-sed -n 's/^##@//p' $< | column -t -s ':' | sed -e 's/[^ ]*/ &/2'
