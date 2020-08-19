NAME=aoc-cli
BIN_DIR=./bin

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
.PHONY: help


## app compile
compile:
	./scripts/compile.sh
.PHONY: compile

## Cross os compile
cross-compile:
	./scripts/cross-compile.sh
.PHONY: cross-compile

## lint project
lint:
	./scripts/run-linters.sh
.PHONY: lint

## lint-ci runs linters for ci.
lint-ci:
	./scripts/run-linters-ci.sh
.PHONY: lint-ci

## format markdown files in project
pretty-markdown:
	find . -name '*.md' -not -wholename './vendor/*' | xargs prettier --write
.PHONY: pretty-markdown

## Test all packages
test:
	./scripts/run-tests.sh
.PHONY: test

## Test coverage
test-cover:
	./scripts/coverage.sh
.PHONY: test-cover

## new-version releases new version with new tag
new-version: lint test compile
	./scripts/version.sh
.PHONY: new-version

## Release
release:
	./scripts/release.sh
.PHONY: release

## Fix imports sorting
imports:
	./scripts/fix-imports.sh
.PHONY: imports

## dependencies - fetch all dependencies for sripts
dependencies:
	./scripts/get-dependencies.sh
.PHONY: dependencies

## vendor-sync checks if all dependencies are correct in go.mod file and if vendor directory is up to date.
vendor-sync:
	./scripts/sync-vendor.sh
.PHONY: vendor-sync

## vendor-check if dependencies were not changed.
vendor-check:
	./scripts/check-vendor.sh
.PHONY: vendor-check

vet:
	./scripts/vet.sh
.PHONY: vet

.DEFAULT_GOAL := test