#!/bin/bash

set -Eeuo pipefail

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(dirname "$0")"
REPO_ROOT="$(cd "${SCRIPT_DIR}" && git rev-parse --show-toplevel)"
SCRIPTS_DIR="${REPO_ROOT}/scripts"
COVER_DIR=${REPO_ROOT}/coverage

source "${SCRIPTS_DIR}/helpers-source.sh"

echo "${SCRIPT_NAME} is running... "

{
  echo "mode: atomic"
  tail -q -n +2 "${COVER_DIR}"/*.cov
} >>"${COVER_DIR}/full.cov"

go_tool gocov convert "${COVER_DIR}/full.cov" >"${COVER_DIR}/full.json"

go_tool gocov-html "${COVER_DIR}/full.json" >"${COVER_DIR}/full.html"

echo "${SCRIPT_NAME} is done... "
