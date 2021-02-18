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

build-ci:
	./scripts/compile.sh
.PHONY: build-ci

## Cross os compile
cross-compile:
	./scripts/cross-compile.sh
.PHONY: cross-compile

## lint project
lint:
	./scripts/run-linters.sh
.PHONY: lint

## Lint in CI
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

## Increase version number
new-version: vet test compile
	./scripts/version.sh
.PHONY: new-version

## Release
release:
	./scripts/release.sh
.PHONY: release

## Release local snapshot
release-local-snapshot:
	${call colored, release is running...}
	./scripts/local-snapshot-release.sh
.PHONY: release-local-snapshot

## Fix imports sorting.
imports:
	${call colored, fix-imports is running...}
	./scripts/fix-imports.sh
.PHONY: imports

## Format code.
fmt:
	${call colored, fmt is running...}
	./scripts/fmt.sh
.PHONY: fmt

## Format code and sort imports.
format-project: fmt imports
.PHONY: format-project

## fetch all dependencies for scripts
install-tools:
	./scripts/get-dependencies.sh
.PHONY: install-tools

## Sync vendor
sync-vendor:
	${call colored, gomod is running...}
	./scripts/sync-vendor.sh
.PHONY: sync-vendor

## Update dependencies
gomod-update:
	${call colored, gomod is running...}
	go get -u -v ./...
	make sync-vendor
.PHONY: gomod-update


.DEFAULT_GOAL := test

## vendor-check if dependencies were not changed.
vendor-check:
	./scripts/check-vendor.sh
.PHONY: vendor-check

## vet project
vet:
	./scripts/vet.sh
.PHONY: vet

.DEFAULT_GOAL := test