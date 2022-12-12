SHELL := env VERSION=$(VERSION) $(SHELL)
VERSION ?= $(shell git describe --tags $(git rev-list --tags --max-count=1))

APP_NAME?=aoc-cli
SHELL := env APP_NAME=$(APP_NAME) $(SHELL)

GOTOOLS_IMAGE_TAG?=v0.0.1
SHELL := env GOTOOLS_IMAGE_TAG=$(GOTOOLS_IMAGE_TAG) $(SHELL)

COMPOSE_CMD=docker compose -f deployments/docker-compose/go-tools-docker-compose.yml up --exit-code-from

TARGET_MAX_CHAR_NUM=20

## Show help
help:
	${call colored, help is running...}
	@echo ''
	@echo 'Usage:'
	@echo '  make <target>'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  %-$(TARGET_MAX_CHAR_NUM)s %s\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

## Build project.
build: compile-app
.PHONY: build

## Compile app.
compile-app:
	./scripts/build/app.sh
.PHONY: compile-app

## Test coverage report.
test-cover:
	$(COMPOSE_CMD) run-tests-coverage run-tests-coverage
.PHONY: test-cover

## Tests sonar report generate.
test-sonar-report:
	./scripts/tests/sonar-report.sh
.PHONY: test-sonar-report

## Open coverage report.
open-cover-report: test-cover
	./scripts/open-coverage-report.sh
.PHONY: open-cover-report

## Update readme coverage.
update-readme-cover: build test-cover
	$(COMPOSE_CMD) update-readme-coverage update-readme-coverage
.PHONY: update-readme-cover

## Run tests.
test:
	$(COMPOSE_CMD) run-tests run-tests
.PHONY: test

## Run regression tests.
test-regression:
	$(COMPOSE_CMD) run-tests-regression run-tests-regression
.PHONY: test-regression

## Sync vendor and install needed tools.
configure: sync-vendor install-tools

## Sync vendor with go.mod.
sync-vendor:
	./scripts/sync-vendor.sh
.PHONY: sync-vendor

## Fix imports sorting.
imports:
	$(COMPOSE_CMD) fix-imports fix-imports
.PHONY: imports

## Format code with go fmt.
fmt:
	$(COMPOSE_CMD) fix-fmt fix-fmt
.PHONY: fmt

## Format code and sort imports.
format-project: fmt imports
.PHONY: format-project

## Installs vendored tools.
install-tools:
	docker compose -f scripts/go-tools-docker-compose.yml pull
.PHONY: install-tools

## vet project
vet:
	./scripts/linting/run-vet.sh
.PHONY: vet

## Run full linting
lint-full:
	$(COMPOSE_CMD) lint-full lint-full
.PHONY: lint-full

## Run linting for build pipeline
lint-pipeline:
	$(COMPOSE_CMD) lint-pipeline lint-pipeline
.PHONY: lint-pipeline

## Run linting for sonar report
lint-sonar:
	$(COMPOSE_CMD) lint-sonar lint-sonar
.PHONY: lint-sonar

## recreate all generated code and documentation.
codegen:
	$(COMPOSE_CMD) go-generate go-generate
.PHONY: codegen

## recreate all generated code and swagger documentation and format code.
generate: codegen format-project vet
.PHONY: generate

## Release
release:
	./scripts/release/release.sh
.PHONY: release

## Release local snapshot
release-local-snapshot:
	./scripts/release/local-snapshot-release.sh
.PHONY: release-local-snapshot

## Check goreleaser config.
check-releaser:
	./scripts/release/check.sh
.PHONY: check-releaser

## Issue new release.
new-version: vet test-regression build
	./scripts/release/new-version.sh
.PHONY: new-release

## Open advent of code homepage in browser.
open-advent-homepage:
	./scripts/browser-opener.sh -u 'https://adventofcode.com/'

.DEFAULT_GOAL := help

