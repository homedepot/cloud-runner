SHELL := /bin/bash -o pipefail

SPINNAKERFILE_TEMPLATES ?= "../spinnakerfile-templates"

MANIFEST_FILES := $(shell find cd/spinnaker-cluster/cloud-runner/kustomize -name '*.yaml')

GENERATE_FILES := $(shell find . -name "*.go" -type f -print | xargs grep -l "//go:generate counterfeiter")
FAKE_FILES     := $(shell find . -name "*.go" -type f -print | xargs grep -l "//go:generate counterfeiter" | sed -e 's~\(.*\)/\(.*\)/\(.*\.go\)~\1/\2/\2fakes/fake_\3~')

IMAGE_TAG  := ${LOGNAME}-$(shell date +'%Y%m%d%H%M')
IMAGE_NAME := gcr.io/github-replication-sandbox/cd/cloud-runner:${IMAGE_TAG}

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

## lint: Runs linter
.PHONY: lint
lint:
	@echo "🧼 Running lint"
	@ci/scripts/go/lint.sh

## test: Runs unit tests
.PHONY: test
test: clean lint
	@echo "🚌 Running unit tests"
	@ci/scripts/go/unit-test.sh

## build: Builds application
.PHONY: build
build: clean test $(FAKE_FILES)
	@echo "🚧 Building application"
	@ci/scripts/go/build.sh -o build/cloud-runner

## docker: Dockerizes local version of cloud-runner and publishes to GCR
.PHONY: docker
docker: build
	@echo "🐳 Creating docker image"
	@docker build --file=./Dockerfile --tag=${IMAGE_NAME} .
	@echo "📚 Publishing docker image"
	@docker push ${IMAGE_NAME}

## manifest: Renders kubernetes manifests
.PHONY: manifest
manifest:
	@BUILD_DIR=./build \
	 IMAGE_TAG=${IMAGE_TAG} \
		make -f cd/spinnaker-cluster/cloud-runner/Makefile manifest

## deploy: Deploys local version of cloud-runner to kubernetes
.PHONY: deploy
deploy: docker manifest
	@echo "🚀 Deploying kubernetes manifests"
	@kubectl apply -f build/manifest-np-us-central1.yaml

## spinnaker: Renders Spinnakerfile (needs arm CLI)
.PHONY: spinnaker
spinnaker:
	@BUILD_DIR=./build make -f cd/spinnaker-cluster/cloud-runner/Makefile spinnaker

## submodules: Updates git submodules to most recent version
.PHONY: submodules
submodules:
	@echo "🐙 Updating ci-scripts submodule"
	@git submodule update --init --recursive
	@cd ci/scripts && git checkout master && git pull

# Will (re)create when .secrets is older than any of the manifest YAMLs
## .secrets: retrieves vault/k8s secrets required to run application
.secrets: $(MANIFEST_FILES)
	@rm -rf .secrets
	@make manifest
	@ci/scripts/secrets/vault-injector.sh
	@ci/scripts/secrets/k8s-injector.sh

## db-start: Starts a local MySQL instance running in Docker
.PHONY: db-start
db-start:
	@MYSQL_DATABASE=cloudrunner ci/scripts/db/mysql/run.sh --start

## db-start: Stops the local MySQL instance running in Docker
.PHONY: db-stop
db-stop:
	@MYSQL_DATABASE=cloudrunner ci/scripts/db/mysql/run.sh --stop

## run: Runs cloud-runner locally
.PHONY: run
run: .secrets $(FAKE_FILES) db-start
	@echo "🚧 Building application"
	@GOOS=darwin GOARCH=amd64 ci/scripts/go/build.sh -o build/cloud-runner
	@echo "🚀 Starting application"
	@SQL_HOST=localhost \
	  SQL_USER="root" \
	  SQL_PASS="password" \
		ci/scripts/go/run.sh build/cloud-runner

$(FAKE_FILES): $(GENERATE_FILES)
	@echo "🔋 Generating fakes"
	@go generate ./...

