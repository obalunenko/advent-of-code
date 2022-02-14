# mango-coral

[![Latest Release](https://img.shields.io/github/release/muesli/mango-coral.svg)](https://github.com/muesli/mango-coral/releases)
[![Build Status](https://github.com/muesli/mango-coral/workflows/build/badge.svg)](https://github.com/muesli/mango-coral/actions)
[![Go ReportCard](https://goreportcard.com/badge/muesli/mango-coral)](https://goreportcard.com/report/muesli/mango-coral)
[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://pkg.go.dev/github.com/muesli/mango-coral)

[coral](https://github.com/muesli/coral/tree/coral) adapter for [mango](https://github.com/muesli/mango).

## Example

```go
import (
	"fmt"

	mcoral "github.com/muesli/mango-coral"
	"github.com/muesli/roff"
	"github.com/muesli/coral"
)

var (
    rootCmd = &coral.Command{
        Use:   "mango",
        Short: "A man-page generator",
    }
)

func main() {
    manPage, err := mcoral.NewManPage(1, rootCmd)
    if err != nil {
        panic(err)
    }

    manPage = manPage.WithSection("Copyright", "(C) 2022 Christian Muehlhaeuser.\n"+
        "Released under MIT license.")

    fmt.Println(manPage.Build(roff.NewDocument()))
}
```
