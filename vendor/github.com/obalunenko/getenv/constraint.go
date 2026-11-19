package getenv

import (
	"net"
	"net/netip"
	"net/url"
	"time"
)

// Parsable is a constraint for the types that Env and EnvOrDefault can parse.
// Every type is supported both as a single value and as a slice.
type Parsable interface {
	string | []string |
		int | []int | int8 | []int8 | int16 | []int16 | int32 | []int32 | int64 | []int64 |
		uint | []uint | uint8 | []uint8 | uint16 | []uint16 | uint32 | []uint32 | uint64 | []uint64 | uintptr | []uintptr |
		float32 | []float32 | float64 | []float64 |
		complex64 | []complex64 | complex128 | []complex128 |
		bool | []bool |
		time.Time | []time.Time | time.Duration | []time.Duration |
		url.URL | []url.URL |
		net.IP | []net.IP | net.HardwareAddr | []net.HardwareAddr |
		netip.Addr | []netip.Addr | netip.Prefix | []netip.Prefix
}

// Constraints of the number parsers.
type (
	signed interface {
		int | int8 | int16 | int32 | int64
	}

	unsigned interface {
		uint | uint8 | uint16 | uint32 | uint64 | uintptr
	}

	float interface {
		float32 | float64
	}

	complexNumber interface {
		complex64 | complex128
	}
)
