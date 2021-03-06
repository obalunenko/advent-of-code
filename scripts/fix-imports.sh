#!/usr/bin/env bash

set -Eeuo pipefail

function cleanup() {
  trap - SIGINT SIGTERM ERR EXIT
  echo "cleanup running"
}

trap cleanup SIGINT SIGTERM ERR EXIT

SCRIPT_NAME="$(basename "$(test -L "$0" && readlink "$0" || echo "$0")")"

echo "${SCRIPT_NAME} is running... "

if [[ -f "$(go env GOPATH)/bin/gogroup" ]] || [[ -f "/usr/local/bin/gogroup" ]]; then
    gogroup -order std,other,prefix=github.com/obalunenko/  -rewrite $(find . -type f -name "*.go" | grep -v "vendor/" |grep -v ".git" )
else
    printf "Cannot check gogroup, please run:
    make install-tools \n"
    exit 1
fi

echo "Done."