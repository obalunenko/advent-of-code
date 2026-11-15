SHELL := env VERSION=$(VERSION) $(SHELL)
VERSION ?= $(shell git describe --tags $(git rev-list --tags --max-count=1))

GOVERSION:=1.27.4

APP_NAME?=aoc-cli
SHELL := env APP_NAME=$(APP_NAME) $(SHELL)

AOC_PUZZLE_URL=
SHELL := env AOC_PUZZLE_URL=$(AOC_PUZZLE_URL) $(SHELL)

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

bump-go-version:
	./scripts/bump-go.sh $(GOVERSION)
.PHONY: bump-go-version

## Build project.
build: sync-vendor generate compile-app
.PHONY: build

## Compile app.
compile-app:
	./scripts/build/app.sh
.PHONY: compile-app


## Test coverage report.
test-cover:
	./scripts/tests/coverage.sh
.PHONY: test-cover

prepare-cover-report: test-cover
	./scripts/tests/prepare-cover-report.sh
.PHONY: prepare-cover-report

## Open coverage report.
open-cover-report: prepare-cover-report
	./scripts/open-coverage-report.sh
.PHONY: open-cover-report

## Update readme coverage.
update-readme-cover: build prepare-cover-report
	./scripts/update-readme-coverage.sh
.PHONY: update-readme-cover

## Run tests.
test:
	./scripts/tests/run.sh
.PHONY: test

## Run regression tests.
test-regression:
	./scripts/tests/run-regression.sh
.PHONY: test-regression

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

## Build tools pinned in tools/<name>/go.mod.
install-tools:
	./scripts/install/tools.sh
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

gen-boilerplate:
	./scripts/codegen/puzzle-boilerplate.sh
.PHONY: gen-boilerplate

.DEFAULT_GOAL := help

