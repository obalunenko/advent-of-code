NAME=aoc-cli
BIN_DIR=./bin

SHELL := env VERSION=$(VERSION) $(SHELL)
VERSION ?= $(shell git describe --tags $(git rev-list --tags --max-count=1))

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



build: compile-aoc-cli
.PHONY: build

compile-aoc-cli:
	./scripts/build/aoc-cli.sh
.PHONY: compile-spamassassin-parser-be

## Test coverage report.
test-cover:
	./scripts/tests/coverage.sh
.PHONY: test-cover

## Tests sonar report generate.
test-sonar-report:
	./scripts/tests/sonar-report.sh
.PHONY: test-sonar-report

## Open coverage report.
open-cover-report: test-cover
	./scripts/open-coverage-report.sh
.PHONY: open-cover-report

update-readme-cover: build test-cover
	./scripts/update-readme-coverage.sh
.PHONY: update-readme-cover

test:
	./scripts/tests/run.sh
.PHONY: test


configure: sync-vendor

sync-vendor:
	./scripts/sync-vendor.sh
.PHONY: sync-vendor

## Fix imports sorting.
imports:
	${call colored, fix-imports is running...}
	./scripts/style/fix-imports.sh
.PHONY: imports

## Format code with go fmt.
fmt:
	./scripts/style/fmt.sh
.PHONY: fmt

## Format code and sort imports.
format-project: fmt imports
.PHONY: format-project

install-tools:
	./scripts/install/vendored-tools.sh
.PHONY: install-tools

## vet project
vet:
	./scripts/linting/run-vet.sh
.PHONY: vet

## Run full linting
lint-full:
	./scripts/linting/run-linters.sh
.PHONY: lint-full

## Run linting for build pipeline
lint-pipeline:
	./scripts/linting/golangci-pipeline.sh
.PHONY: lint-pipeline

## Run linting for sonar report
lint-sonar:
	./scripts/linting/golangci-sonar.sh
.PHONY: lint-sonar

## recreate all generated code and documentation.
codegen:
	./scripts/codegen/go-generate.sh
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

## Issue new release.
new-version: vet test build
	./scripts/release/new-version.sh
.PHONY: new-release

open-advent-homepage:
	./scripts/browser-opener.sh -u 'https://adventofcode.com/'

.DEFAULT_GOAL := help

