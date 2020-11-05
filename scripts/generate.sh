#!/usr/bin/env bash

  echo "generator is running..."
    if [[ -f "$(go env GOPATH)/bin/go-bindata" ]] || [[ -f "/usr/local/bin/go-bindata" ]]; then
        go generate -tags=input ./...
    else
        printf "Cannot check go-bindata, please run:
        go get -u -v github.com/shuLhan/go-bindata \n"
        exit 1
    fi
  echo "Done"
