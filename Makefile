PACKAGES=$(shell go get)
SOURCE_FILES=$(shell find . -name '*.go' -not -path '*vendor*')
VERSION?=$(shell cat .version 2> /dev/null || echo "local")
BINARY:=aoc

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

ENVIRONMENT=$(shell uname)
ifeq ($(ENVIRONMENT), Darwin)
	BASE64CALL = base64
endif

ifeq ($(UNAME), Linux)
	BASE64CALL = base64 -w 0
endif

TEMPLATE_DIR = templates
INIT_FILE = aoc-boilerplate/templates/init.go
files: $(TEMPLATE_DIR)/*_template.*
	echo "Running $@"
	echo 'package templates' > $(INIT_FILE)
	echo 'func Initialize() {' >> $(INIT_FILE)
	echo '  templates = make(map[string]string)' >> $(INIT_FILE)
	echo '  fileExtensions = make(map[string]string)' >> $(INIT_FILE)
	for file in $^ ; do \
		base_name=$$(basename $${file}) ; \
		template_name=$$(echo "$${base_name%.*}") ; \
        extension=$$(echo "$${file#*.}") ; \
        lang=$$(echo "$${base_name%_*}") ; \
		echo "  templates[\"$${lang}\"] = \"$$(${BASE64CALL} $${file})\"" >> $(INIT_FILE) ; \
		echo "  fileExtensions[\"$${lang}\"] = \"$${extension}\"" >> $(INIT_FILE) ; \
	done
	echo '}' >> $(INIT_FILE)



