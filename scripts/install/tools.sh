#!/bin/bash

set -Eeuo pipefail

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(dirname "$0")"
REPO_ROOT="$(cd "${SCRIPT_DIR}" && git rev-parse --show-toplevel)"
SCRIPTS_DIR="${REPO_ROOT}/scripts"

source "${SCRIPTS_DIR}/helpers-source.sh"

echo "${SCRIPT_NAME} is running... "

for dir in "${REPO_ROOT}"/tools/*/; do
  name="$(basename "${dir}")"
  echo "Building ${name}"
  tool_path "${name}" >/dev/null
done

echo "${SCRIPT_NAME} done."
