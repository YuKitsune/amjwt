
# Versioning information
GIT_COMMIT := $(shell git rev-parse HEAD)
GIT_BRANCH := $(shell git name-rev --name-only HEAD | sed "s/~.*//")

## Gets the current tag name or commit SHA
VERSION ?= $(shell git describe --tags ${COMMIT} 2> /dev/null || echo "$(GIT_COMMIT)")

## Gets the -ldflags for the go build command, this lets us set the version number in the binary
ROOT := github.com/yukitsune/amjwt
LD_FLAGS := -X '$(ROOT).Version=$(VERSION)'

## Whether the repo has uncommitted changes
GIT_DIRTY := false
ifneq ($(shell git status -s),)
	GIT_DIRTY=true
endif

# Commands

.DEFAULT_GOAL := help

.PHONY: help
help: ## Shows this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Builds all programs and places their binaries in the bin/ directory
	mkdir -p bin
	go build -ldflags="$(LD_FLAGS)" -o ./bin/  ./cmd/...

.PHONY: test
test: ## Runs all tests
	go test ./...

.PHONY: clean
clean: ## Removes the bin/ directory
	rm -rf bin
