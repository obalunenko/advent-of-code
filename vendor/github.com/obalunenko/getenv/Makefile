NAME=melsoft/bitech-go-shared

# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)


TARGET_MAX_CHAR_NUM=20


define colored
	@echo '${GREEN}$1${RESET}'
endef

## Show help
help:
	${call colored, help is running...}
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)


## Installs all tools.
install-tools: install-vendored-tools
.PHONY: install-tools

## Installs tools from vendor.
install-vendored-tools:
	./scripts/install/vendored-tools.sh
.PHONY: install-vendored-tools

## Sync vendor of root project and tools.
sync-vendor:
	./scripts/sync-vendor.sh
.PHONY: sync-vendor

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

## Test the project (excluding integration tests).
test:
	./scripts/tests/run.sh
.PHONY: test

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

## Update readme coverage badge.
update-readme-cover: test-cover
	./scripts/update-readme-coverage.sh
.PHONY: update-readme-cover

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

## Create release notes for GH release.
release:
	./scripts/release/release.sh
.PHONY: release

## Issue new release.
new-version: vet test
	./scripts/release/new-version.sh
.PHONY: new-release

.DEFAULT_GOAL := help