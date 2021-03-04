SHELL := /bin/bash -o pipefail

SPINNAKERFILE_TEMPLATES ?= "../spinnakerfile-templates"

MANIFEST_FILES := $(shell find cd/kustomize -name '*.yaml')

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

## test: Runs unit tests
.PHONY: test
test: clean
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
manifest: clean
	@echo "📦 Generating kubernetes manifests"
	@kubectl kustomize cd/kustomize/overlays/github-replication-sandbox/us-central1 | \
		sed -e "s~\$${#toInt(parameters\[replicas\])}~1~g" \
		    -e "s~\$${imageTag}~${IMAGE_TAG}~g" > build/manifest.yaml

## deploy: Deploys local version of cloud-runner to kubernetes
.PHONY: deploy
deploy: docker manifest
	@echo "🚀 Deploying kubernetes manifests"
	@kubectl apply -f build/manifest.yaml

## spinnaker: Renders Spinnakerfile (needs arm CLI)
.PHONY: spinnaker
spinnaker: clean
	@echo "⛵️ Validating Spinnakerfile"
	@cat cd/spinnaker/Spinnakerfile | \
		sed -e "s~cd/spinnaker/modules/~../cd/spinnaker/modules/~g" > build/Spinnakerfile.module
	@arm dinghy render build/Spinnakerfile.module --modules ${SPINNAKERFILE_TEMPLATES} --output build/Spinnakerfile

## submodules: Updates git submodules to most recent version
.PHONY: submodules
submodules:
	@echo "🐙 Updating ci-scripts submodule"
	@git submodule update --init
	@cd ci/scripts && git fetch && git merge

# Will (re)create when .secrets is older than any of the manifest YAMLs
.secrets: $(MANIFEST_FILES)
	@rm -rf .secrets
	@make manifest
	@ci/scripts/secrets/vault-injector.sh
	@ci/scripts/secrets/k8s-injector.sh

## db-start: Starts a local MySQL instance running in Docker
.PHONY: db-start
db-start:
	@MYSQL_DATABASE=cloud-runner ci/scripts/db/mysql/run.sh --start

## db-start: Stops the local MySQL instance running in Docker
.PHONY: db-stop
db-stop:
	@MYSQL_DATABASE=cloud-runner ci/scripts/db/mysql/run.sh --stop

## run: Runs cloud-runner locally
.PHONY: run
run: .secrets $(FAKE_FILES) cloud-runner db-start
	@echo "🚧 Building application"
	@GOOS=darwin GOARCH=amd64 ci/scripts/go/build.sh -o build/cloud-runner
	@echo "🚀 Starting application"
	@ci/scripts/go/run.sh build/cloud-runner

$(FAKE_FILES): $(GENERATE_FILES)
	@echo "🔋 Generating fakes"
	@go generate ./...
	
