#!/bin/sh

set -eu pipefail

SCRIPT_NAME="$(basename "$0")"

echo "${SCRIPT_NAME} is running... "

export AOC_REGRESSION_ENABLED=true

REGRESSION_TESTS_PKG=$(go list -m)/tests

GOTEST="go test -v -race"
if [ -f "$(go env GOPATH)/bin/gotestsum" ] || [ -f "/usr/local/bin/gotestsum" ]; then
  GOTEST="gotestsum --format pkgname --"
fi

${GOTEST} "${REGRESSION_TESTS_PKG}"

unset AOC_REGRESSION_ENABLED

echo "${SCRIPT_NAME} done."
