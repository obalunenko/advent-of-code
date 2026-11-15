package internal

import (
	"net"
	"net/netip"
	"net/url"
	"time"
)

type (
	// EnvParsable is a constraint for types that can be parsed from environment variable.
	EnvParsable interface {
		String | Number | NumberSlice | Time | Bool | URL | Network | Complex | ComplexSlice
	}

	// String is a constraint for string and slice of strings.
	String interface {
		string | []string
	}

	// Number is a constraint for integer, float and unsigned integer.
	Number interface {
		Int | Float | Uint
	}

	// NumberSlice is a constraint for slice of integers, floats and unsigned integers.
	NumberSlice interface {
		IntSlice | FloatSlice | UintSlice
	}

	// Int is a constraint for integer scalar types.
	Int interface {
		int | int8 | int16 | int32 | int64
	}

	// IntSlice is a constraint for slice of integers.
	IntSlice interface {
		[]int | []int8 | []int16 | []int32 | []int64
	}

	// Uint is a constraint for unsigned integer scalar types.
	Uint interface {
		uint | uint8 | uint16 | uint32 | uint64 | uintptr
	}

	// UintSlice is a constraint for slice of unsigned integers.
	UintSlice interface {
		[]uint | []uint8 | []uint16 | []uint32 | []uint64 | []uintptr
	}

	// Float is a constraint for floating-point scalar types.
	Float interface {
		float32 | float64
	}

	// FloatSlice is a constraint for slice of floats.
	FloatSlice interface {
		[]float32 | []float64
	}

	// Time is a constraint for time.Time and time.Duration and slices of them.
	Time interface {
		time.Time | []time.Time | time.Duration | []time.Duration
	}

	// Bool is a constraint for bool and slice of bool.
	Bool interface {
		bool | []bool
	}

	// URL is a constraint for url.URL and slice of url.URL.
	URL interface {
		url.URL | []url.URL
	}

	// Network is a constraint for network address types and slices of them.
	Network interface {
		net.IP | []net.IP |
			net.HardwareAddr | []net.HardwareAddr |
			netip.Addr | []netip.Addr |
			netip.Prefix | []netip.Prefix
	}

	// ComplexSlice is a constraint for slice of complex.
	ComplexSlice interface {
		[]complex64 | []complex128
	}

	// Complex is a constraint for complex scalar types.
	Complex interface {
		complex64 | complex128
	}
)
