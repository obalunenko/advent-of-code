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



## Build project.
build: compile-aoc-cli
.PHONY: build

## Compile aoc-cli.
compile-aoc-cli:
	./scripts/build/aoc-cli.sh
.PHONY: compile-spamassassin-parser-be

## Test coverage report.
test-cover:
	./scripts/tests/coverage.sh
.PHONY: test-cover

## Test regression.
test-regression:
	./scripts/tests/run-regression.sh
.PHONY: test-regression

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
	./scripts/update-readme-coverage.sh
.PHONY: update-readme-cover

## Run tests.
test:
	./scripts/tests/run.sh
.PHONY: test

## Sync vendor and install needed tools.
configure: sync-vendor install-tools

## Sync vendor with go.mod.
sync-vendor:
	./scripts/sync-vendor.sh
.PHONY: sync-vendor

## Fix imports sorting.
imports:
	./scripts/style/fix-imports.sh
.PHONY: imports

## Format code with go fmt.
fmt:
	./scripts/style/fmt.sh
.PHONY: fmt

## Format code and sort imports.
format-project: fmt imports
.PHONY: format-project

## Installs vendored tools.
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

