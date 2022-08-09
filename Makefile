# The master Makefile that controls the Makefiles of all recipes

# -f -name 'Makefile' = Remove all Makefiles in subdirectories
# -mindepth 2 = Exclude the root Makefile (this file)
# -path './.*/*' = Exclude directories that starts with '.'
MAKEFILES := $(shell find . -mindepth 2 -type f -name 'Makefile' !  -path './.*/*')
SUBDIRS := $(filter-out ./,$(dir $(MAKEFILES)))

# Colorized print
BOLD := $(shell tput bold)
RED := $(shell tput setaf 1)
CYAN := $(shell tput setaf 6)
RESET := $(shell tput sgr0)

# -= Executed once
# =  Executed at every invocation

# Set up the default target all and set up phony targets ie. they are not files.

.PHONY: all

all: build

##@ build: Build all recipes

.PHONY: build

# @  = Don't print command.
# @- = Don't print command and ignore error.
# @? = target
build:
	@-echo "$(BOLD)$(CYAN)Building all recipes...$(RESET)"
	@for subdir in $(SUBDIRS); do \
  	echo "Building $$(basename $$subdir)"; \
  	make -C $$subdir $@; \
  done

##@ clean: Clean output files and build cache

.PHONY: clean

clean:
	@-echo "$(BOLD)$(RED)Cleaning all recipes...$(RESET)"
	@for subdir in $(SUBDIRS); do \
    echo "Cleaning $$(basename $$subdir)"; \
    make -C $$subdir $@; \
  done

##@ help: Help

.PHONY: help

help: Makefile
	@-echo " Usage:\n  make <target>"
	@-echo
	@-sed -n 's/^##@//p' $< | column -t -s ':' | sed -e 's/[^ ]*/ &/2'
