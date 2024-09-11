APPLICATION_NAME := $(shell grep "const ApplicationName " version.go | sed -E 's/.*"(.+)"$$/\1/')
BIN_NAME=${APPLICATION_NAME}

# grep the version from the mix file
BASE_VERSION := $(shell grep "const Version " version.go | sed -E 's/.*"(.+)"$$/\1/')
VERSION="${BASE_VERSION}.$(shell date +%s | head -c 9)"

GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)

# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

default: help

build: ## Build project for development
	@echo "building ${BIN_NAME} ${BASE_VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags "-X main.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X main.VersionPrerelease=DEV" -o bin/${BIN_NAME} ./

dep: ## Install projects dependecies with Glide
	go mod tidy
	go mod vendor

clean: ## Clean build project
	@test ! -e bin/${BIN_NAME} || rm bin/${BIN_NAME}

run-test:  ## Run project tests
	mkdir -p ./test/cover
#	go test -race -coverpkg=./... -coverprofile=./test/cover/cover.out
	go test -coverpkg=./... -coverprofile=./test/cover/cover.out
	go tool cover -html=./test/cover/cover.out -o ./test/cover/cover.html
