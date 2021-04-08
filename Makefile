SHELL := /bin/bash -o pipefail

GENERATE_FILES := $(shell find . -name "*.go" -type f -print | xargs grep -l "//go:generate counterfeiter")
FAKE_FILES     := $(shell find . -name "*.go" -type f -print | xargs grep -l "//go:generate counterfeiter" | sed -e 's~\(.*\)/\(.*\)/\(.*\.go\)~\1/\2/\2fakes/fake_\3~')

.PHONY: all
all: help

.PHONY: help
help: Makefile
	@echo
	@echo " Choose a command to run in cloud-runner:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

.PHONY: clean
clean:
	@echo "🧹 Cleaning old build"
	@rm -rf build
	@mkdir build

## test: Runs unit tests
.PHONY: test
test: clean
	@echo "🚌 Running unit tests"
	@go test -v ./...

## build: Builds application
.PHONY: build
build: clean test $(FAKE_FILES)
	@echo "🚧 Building application"
	@go build -o build/cloud-runner cmd/cloud-runner/cloud-runner.go

## run: Runs cloud-runner locally
.PHONY: run
run: $(FAKE_FILES) cloud-runner
	@echo "🚀 Starting application"
	@GOOS=darwin GOARCH=amd64 go run cmd/cloud-runner/cloud-runner.go

$(FAKE_FILES): $(GENERATE_FILES)
	@echo "🔋 Generating fakes"
	@go generate ./...

