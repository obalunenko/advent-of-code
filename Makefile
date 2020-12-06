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
compile: generate
	./scripts/compile.sh
.PHONY: compile

build-ci:
	./scripts/compile.sh
.PHONY: build-ci

## Cross os compile
cross-compile: generate
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
release: generate
	./scripts/release.sh
.PHONY: release

## Fix imports sorting
imports:
	./scripts/fix-imports.sh
.PHONY: imports

## fetch all dependencies for scripts
dependencies:
	./scripts/get-dependencies.sh
.PHONY: dependencies

## Sync vendor
gomod:
	${call colored, gomod is running...}
	go mod tidy -v
	go mod verify
	go mod download
	go mod vendor
.PHONY: gomod

## Update dependencies
gomod-update:
	${call colored, gomod is running...}
	go get -u -v ./...
	make gomod
.PHONY: gomod-update

## Recreate generated files
generate:
	${call colored, generate is running...}
	./scripts/generate.sh
.PHONY: generate

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