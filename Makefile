# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
BINARY_NAME=jwt
BINARY_UNIX=$(BINARY_NAME)_unix
LINTER=golangci-lint

export GO111MODULE=on

all: setup lint-docker test build

setup:
		$(GOMOD) download
		$(GOMOD) tidy

build: clean
		$(GOBUILD) -o build/bin/$(BINARY_NAME) -v main.go

build-linux: clean
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o build/bin/$(BINARY_UNIX) -v

clean:
		$(GOCLEAN)
		rm -f build/bin/$(BINARY_NAME)
		rm -f build/bin/$(BINARY_UNIX)

test:
		rm -f coverage.out
		$(GOTEST) ./... -coverprofile=coverage.out

show-coverage: test
		$(GOCMD) tool cover -html=coverage.out

lint-docker:
		docker run --rm -v $(CURDIR):/app -w /app golangci/golangci-lint:v1.40.1 golangci-lint run -v

lint:
		$(LINTER) run

.PHONY: all setup build build-linux clean test show-coverage lint-docker lint