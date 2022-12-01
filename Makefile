PACKAGES=$(shell go get)
SOURCE_FILES=$(shell find . -name '*.go' -not -path '*vendor*')
VERSION?=$(shell git describe --match 'v[0-9].*' --dirty='.m' --always --tags)
BINARY:=aoc_${VERSION}

.PHONY: all build check clean coverage fmt help lint test vet binaries files

binaries: clean build-darwin build-linux build-windows

all: check build test ## run fmt, vet, lint, build the binaries and run the tests

check: fmt vet lint ## run fmt, vet, lint

vet: ## run go vet
	@echo "Running $@"
	@test -z "$$(go vet ${PACKAGES} 2>&1 | tee /dev/stderr)"

fmt: ## run go fmt
	@echo "Running $@"
	@gofmt -s -l -w ${SOURCE_FILES}

build-windows: files ## build the go packages
	@echo "Running $@"
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-X main.Version=${VERSION}" -o bin/windows/${BINARY}.exe .

build-darwin: files ## build the go packages
	@echo "Running $@"
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.Version=${VERSION}" -o bin/darwin/${BINARY}-darwin-amd64 .

build-linux: files ## build the go packages for Linux (useful to copy the binary into docker)
	@echo "Running $@"
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=${VERSION}" -o bin/debian/${BINARY}-debian-amd64 .

test: ## run test
	@echo "Running $@"
	@go test ${PACKAGES}

coverage: ## run tests with coverage metrics
	@echo "Running $@"
	@go test -cover ${PACKAGES}

clean: ## clean up binaries
	@echo "Running $@"
	@rm -rf bin
	@rm -f langs.go

run: ## run cli
	@go run main.go ${ARGS}


help: ## this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

files: languages/boilerplate.go languages/boilerplate.py main.go
	echo 'package main' > langs.go
	echo "const golang = \"$$(cat languages/boilerplate.go | base64)\"" >> langs.go
	echo "const python = \"$$(cat languages/boilerplate.py | base64)\"" >> langs.go

