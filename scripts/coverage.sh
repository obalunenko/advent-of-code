#!/bin/sh

set -eu

SCRIPT_NAME="$(basename "$0")"

echo "${SCRIPT_NAME} is running... "

rm -rf coverage
mkdir -p coverage

# go test --count=1 -tags=integration_test -coverprofile ./coverage/integration.cov -covermode=atomic ./...
go test --count=1 -coverprofile ./coverage/unit.cov -covermode=atomic ./...


{
echo "mode: atomic"
tail -q -n +2 ./coverage/*.cov
} >> ./coverage/full.cov

gocov convert ./coverage/full.cov > ./coverage/full.json
gocov report ./coverage/full.json
gocov-html ./coverage/full.json >./coverage/full.html
# open ./coverage/full.html

# go tool cover -html=./coverage/full.cov

echo "${SCRIPT_NAME} done."
