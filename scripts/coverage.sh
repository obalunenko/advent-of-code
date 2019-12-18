#!/usr/bin/env bash

go test -race -coverpkg=./... -v -coverprofile coverage.out.tmp ./...
cat coverage.out.tmp | grep -v "cmd/aoc-cli" > coverage.out
rm -rf coverage.out.tmp
gocov convert coverage.out | gocov report
go tool cover -html=coverage.out
