#!/bin/bash

set -Eeuo pipefail

function check_installed() {
  CMD=$1

  if ! command -v "${CMD}" &>/dev/null; then
    echo "Install ${CMD}"
    echo "run 'make install-tools' "
    exit 1
  fi
}

# Prints the binary path of a tool pinned in tools/<name>/go.mod, building it on
# first use. go tool runs inside the tool module: -modfile from the repo root
# clashes with vendor/. <NAME>_BIN (e.g. GORELEASER_BIN) overrides the path: CI
# sets it for tools installed by their official actions.
function tool_path() {
  local name="$1"

  local override
  override="$(echo "${name}" | tr '[:lower:]-' '[:upper:]_')_BIN"
  if [[ -n "${!override:-}" ]]; then
    echo "${!override}"
    return
  fi

  go tool -C "${REPO_ROOT}/tools/${name}" -n "${name}"
}

# Runs a tool pinned in tools/<name>/go.mod: go_tool <name> [args...]
function go_tool() {
  local name="$1"
  shift

  local bin
  bin="$(tool_path "${name}")"
  "${bin}" "$@"
}

function open_source() {
  OPEN_CMD=""
  OPEN_SRC=$1

  case "$OSTYPE" in
  darwin*)
    OPEN_CMD=open
    ;;
  linux*)
    OPEN_CMD=xdg-open
    ;;
  msys* | cygwin*)
    OPEN_CMD=start
    ;;
  *)
    echo "unknown: $OSTYPE"
    exit 1
    ;;
  esac

  ${OPEN_CMD} "${OPEN_SRC}"
}
