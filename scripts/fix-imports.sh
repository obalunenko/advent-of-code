#!/usr/bin/env bash
set -e

if [[ -f "$(go env GOPATH)/bin/gci" ]] || [[ -f "/usr/local/bin/gci" ]]; then
  gci -w -local=github.com/obalunenko/  $(find . -type f -name "*.go" | grep -v "vendor/" | grep -v ".git")
else
  printf "Cannot check gogroup, please run:
    go get -u -v github.com/daixiang0/gci/... \n"
  exit 1
fi
