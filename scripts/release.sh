#!/usr/bin/env bash

set -Eeuo pipefail

function cleanup() {
  trap - SIGINT SIGTERM ERR EXIT
  echo "cleanup running"
}

trap cleanup SIGINT SIGTERM ERR EXIT

SCRIPT_NAME="$(basename "$(test -L "$0" && readlink "$0" || echo "$0")")"

echo "${SCRIPT_NAME} is running... "

# Get new tags from the remote
git fetch --tags -f

# Get the latest tag name
latestTag=$(git describe --tags $(git rev-list --tags --max-count=1))
echo ${latestTag}

export GOVERSION=$(go version | awk '{print $3;}')

goreleaser release --rm-dist
