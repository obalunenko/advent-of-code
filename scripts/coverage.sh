#!/usr/bin/env bash

set -Eeuo pipefail

function cleanup() {
  trap - SIGINT SIGTERM ERR EXIT
  echo "cleanup running"
  rm -rf coverage.out.tmp
}

trap cleanup SIGINT SIGTERM ERR EXIT

SCRIPT_NAME="$(basename "$(test -L "$0" && readlink "$0" || echo "$0")")"

echo "${SCRIPT_NAME} is running... "


go test -race -coverpkg=./... -coverprofile coverage.out.tmp ./...

# shellcheck disable=SC2002
cat coverage.out.tmp | grep -v "cmd/" >coverage.out
gocov convert coverage.out >coverage.out.json
gocov report coverage.out.json
gocov-html coverage.out.json > coverage.out.html
go tool cover -html=coverage.out
