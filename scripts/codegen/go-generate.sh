#!/bin/bash

set -Eeuo pipefail

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(dirname "$0")"
REPO_ROOT="$(cd "${SCRIPT_DIR}" && git rev-parse --show-toplevel)"
SCRIPTS_DIR="${REPO_ROOT}/scripts"

source "${SCRIPTS_DIR}/helpers-source.sh"

# go:generate runs stringer from PATH.
STRINGER_BIN="$(tool_path stringer)"
export PATH="$(dirname "${STRINGER_BIN}"):${PATH}"

go generate -x ./...

echo "${SCRIPT_NAME} done."
