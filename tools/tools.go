// +build tools

package tools

//go:generate go clean

//go:generate echo "installing gocov ..."
//go:generate go install -mod=vendor github.com/axw/gocov/gocov

//go:generate echo "installing golangci-lint ..."
//go:generate go install -mod=vendor github.com/golangci/golangci-lint/cmd/golangci-lint

//go:generate echo "installing gocov-html ..."
//go:generate go install -mod=vendor github.com/matm/gocov-html

//go:generate echo "installing gogroup ..."
//go:generate go install -mod=vendor github.com/vasi-stripe/gogroup/cmd/gogroup

//go:generate echo "installing cover ..."
//go:generate go install -mod=vendor golang.org/x/tools/cmd/cover

//go:generate echo "installing goveralls ..."
//go:generate go install -mod=vendor github.com/mattn/goveralls

import (
	_ "github.com/axw/gocov/gocov"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/matm/gocov-html"
	_ "github.com/mattn/goveralls"
	_ "github.com/vasi-stripe/gogroup/cmd/gogroup"
	_ "golang.org/x/tools/cmd/cover"
)
