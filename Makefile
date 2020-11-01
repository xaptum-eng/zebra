BASEDIR = $(shell pwd)
APPNAME = $(shell basename $(BASEDIR))

PACKAGE = github.com/<username>/$(APPNAME)

SOURCE_FILES?=./...
TEST_PATTERN?=.
TEST_OPTIONS?=-v

export GO111MODULE := on

.DEFAULT_GOAL := help

.PHONY: clean tools build build test fmt fmtcheck lint cover clean

tools: ## setup
	GO111MODULE=on go install github.com/golangci/golangci-lint/cmd/golangci-lint

# Uncomment this rule for building executable
#build: fmtcheck lint ## build
#	GOOS=linux GOARCH=amd64 go build -v -o $(GOPATH)/bin/$(APPNAME)-linux-amd64 $(PACKAGE)
#	GOOS=darwin GOARCH=amd64 go build -v -o $(GOPATH)/bin/$(APPNAME)-darwin-amd64 $(PACKAGE)
#	GOOS=windows GOARCH=amd64 go build -v -o $(GOPATH)/bin/$(APPNAME)-windows-amd64 $(PACKAGE)

# Uncomment this rule for building library
#build: fmtcheck lint ## build
#	go build ./...

# If building a library, test should depend on build-lib
test: build ## test
	go test $(TEST_OPTIONS) -failfast -race -coverpkg=./... -covermode=atomic -coverprofile=coverage.txt $(SOURCE_FILES) -run $(TEST_PATTERN) -timeout=2m

fmt: ## format
	@echo "==> Fixing source code with gofmt..."
	gofmt -s -w ./$(PKG_NAME)

fmtcheck: ## check format
	@sh -c "'$(CURDIR)/gofmtcheck.sh'"

lint: ## lint
	@echo "==> Checking source code against linters..."
	@GOGC=30 golangci-lint run ./...

cover: test ## generate coverage report
	go tool cover -html=coverage.txt

clean: ## clean
	rm $(GOPATH)/bin/$(APPNAME)*

help: ## help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
