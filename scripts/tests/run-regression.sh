#!/bin/bash

set -eu pipefail

SCRIPT_NAME="$(basename "$0")"

echo "${SCRIPT_NAME} is running... "

export AOC_REGRESSION_ENABLED=true

REGRESSION_TESTS_PKG=$(go list -m)/tests

GOTEST="go test -v"
if command -v "gotestsum" &>/dev/null; then
  GOTEST="gotestsum --format testname --"
fi

${GOTEST} -count=1 -race "${REGRESSION_TESTS_PKG}"

unset AOC_REGRESSION_ENABLED

echo "${SCRIPT_NAME} done."
