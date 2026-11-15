#!/bin/bash

set -Eeuo pipefail

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(dirname "$0")"
REPO_ROOT="$(cd "${SCRIPT_DIR}" && git rev-parse --show-toplevel)"
SCRIPTS_DIR="${REPO_ROOT}/scripts"

source "${SCRIPTS_DIR}/helpers-source.sh"

echo "${SCRIPT_NAME} is running... "

export AOC_REGRESSION_ENABLED=true

REGRESSION_TESTS_PKG=$(go list -m)/tests

go_tool gotestsum --format testname -- -count=1 -race "${REGRESSION_TESTS_PKG}"

unset AOC_REGRESSION_ENABLED

echo "${SCRIPT_NAME} done."
