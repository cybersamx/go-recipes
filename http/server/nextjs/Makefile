ROOT_DIR := $(shell git rev-parse --show-toplevel)/http/server/nextjs
APP_NAME := server
NEXTJS_DIR := $(ROOT_DIR)/web
APP_BIN_DIR := $(ROOT_DIR)/bin

.PHONY: all
all: run

###@ build: Build next.js and go applications
.PHONY: build
build:
	cd $(NEXTJS_DIR); npm install; npm run build
	@[ -f $(APP_BIN_DIR) ] || mkdir -p $(APP_BIN_DIR)
	go build -o $(APP_BIN_DIR)/$(APP_NAME) main.go


###@ run: Run the application
.PHONY: run
run: build
	@cd $(APP_BIN_DIR); ./$(APP_NAME)


###@ clean: Clean the project
.PHONY: clean
clean:
	@-rm -rf $(APP_BIN_DIR)
	@-rm -rf $(NEXTJS_DIR)/.next
	@-rm -rf $(NEXTJS_DIR)/dist
	@-rm -rf $(NEXTJS_DIR)/node_modules
	go clean -testcache


###@ help: Help
.PHONY: help
help: Makefile
	@-echo "Usage:\n  make $(BLUE)<target>$(RESET)"
	@-echo
	@-sed -n 's/^###@//p' $< | column -t -s ':' | sed -e 's/[^ ]*/ &/2'
