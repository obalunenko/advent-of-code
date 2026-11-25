![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/obalunenko/getenv)
[![Go Reference](https://pkg.go.dev/badge/github.com/obalunenko/getenv.svg)](https://pkg.go.dev/github.com/obalunenko/getenv)
[![Go Report Card](https://goreportcard.com/badge/github.com/obalunenko/getenv)](https://goreportcard.com/report/github.com/obalunenko/getenv)
[![codecov](https://codecov.io/gh/obalunenko/getenv/branch/master/graph/badge.svg)](https://codecov.io/gh/obalunenko/getenv)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=obalunenko_getenv&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=obalunenko_getenv)

# getenv

Package getenv provides a simple way to get environment variables.
It's type-safe and supports built-in types and slices of them.
Empty environment variable values are treated as not set.

Slices are split by "," and time.Time values are parsed with the time.RFC3339 layout,
unless getenv.WithSeparator or getenv.WithTimeLayout sets another one.
Slice elements are not trimmed, so "1, 2" is not a valid []int,
and a trailing separator adds an empty element: "a,b," is []string{"a", "b", ""}.

Types supported, each both as a single value and as a slice (see `Parsable`):
- `string`
- `int`, `int8`, `int16`, `int32`, `int64`
- `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`
- `float32`, `float64`
- `complex64`, `complex128`
- `bool`
- `time.Time`, `time.Duration`
- `url.URL`
- `net.IP`, `net.HardwareAddr`, `netip.Addr`, `netip.Prefix`

## Examples

### Env

Env retrieves the value of the environment variable named by the key.
If the variable is present in the environment, the value will be parsed and returned.
Otherwise, an error will be returned.

```golang
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/obalunenko/getenv"
)

func main() {
	// In a real program the variable comes from the environment.
	if err := os.Setenv("PORT", "8080"); err != nil {
		log.Fatal(err)
	}

	port, err := getenv.Env[int]("PORT")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(port)
}

```

 Output:

```
8080
```

### Env: errors

```golang
package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/obalunenko/getenv"
)

func main() {
	if err := os.Setenv("WORKERS", "four"); err != nil {
		log.Fatal(err)
	}

	_, err := getenv.Env[int]("WORKERS")
	fmt.Println(errors.Is(err, getenv.ErrInvalidValue), err)

	_, err = getenv.Env[int]("GETENV_EXAMPLE_NOT_SET")
	fmt.Println(errors.Is(err, getenv.ErrNotSet), err)

}

```

 Output:

```
true environment variable "WORKERS": invalid value: strconv.ParseInt: parsing "four": invalid syntax
true environment variable "GETENV_EXAMPLE_NOT_SET": not set
```

### Env: slice

```golang
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/obalunenko/getenv"
)

func main() {
	if err := os.Setenv("HOSTS", "a.example.com,b.example.com"); err != nil {
		log.Fatal(err)
	}

	hosts, err := getenv.Env[[]string]("HOSTS")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hosts)
}

```

 Output:

```
[a.example.com b.example.com]
```

### Env: url.URL

```golang
package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/obalunenko/getenv"
)

func main() {
	if err := os.Setenv("DATABASE_URL", "postgres://user:secret@db:5432/app"); err != nil {
		log.Fatal(err)
	}

	dsn, err := getenv.Env[url.URL]("DATABASE_URL")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dsn.Host, dsn.Redacted())
}

```

 Output:

```
db:5432 postgres://user:xxxxx@db:5432/app
```

### EnvOrDefault

EnvOrDefault retrieves the value of the environment variable named by the key.
If the variable is set and its value can be parsed, the parsed value is returned.
Otherwise, defaultVal is returned: both when the variable is not set (or empty)
and when its value is invalid. Use Env to tell these cases apart.

```golang
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/obalunenko/getenv"
)

func main() {
	// "30" has no unit, so it is not a valid time.Duration: EnvOrDefault falls back
	// to the default, the same as for a variable that is not set. Use Env to catch it.
	if err := os.Setenv("TIMEOUT", "30"); err != nil {
		log.Fatal(err)
	}

	timeout := getenv.EnvOrDefault("TIMEOUT", 5*time.Second)
	retries := getenv.EnvOrDefault("GETENV_EXAMPLE_RETRIES", 3)

	fmt.Println(timeout, retries)
}

```

 Output:

```
5s 3
```

### WithSeparator

WithSeparator sets the separator of slice values. The default is ",".

```golang
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/obalunenko/getenv"
)

func main() {
	if err := os.Setenv("PORTS", "80|443"); err != nil {
		log.Fatal(err)
	}

	ports, err := getenv.Env[[]int]("PORTS", getenv.WithSeparator("|"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ports)
}

```

 Output:

```
[80 443]
```

### WithTimeLayout

WithTimeLayout sets the layout of time.Time values. The default is time.RFC3339.

```golang
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/obalunenko/getenv"
)

func main() {
	if err := os.Setenv("RELEASE_DATE", "2022-01-20"); err != nil {
		log.Fatal(err)
	}

	date, err := getenv.Env[time.Time]("RELEASE_DATE", getenv.WithTimeLayout(time.DateOnly))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date)
}

```

 Output:

```
2022-01-20 00:00:00 +0000 UTC
```
