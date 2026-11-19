package getenv

import (
	"errors"
	"fmt"
	"net"
	"net/netip"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var errEmptySeparator = errors.New("empty separator")

// parseAs converts the raw value of an environment variable to T.
// Slices are split by params.separator, time.Time is parsed with params.layout.
// Every returned error wraps ErrInvalidValue.
func parseAs[T Parsable](raw string, p params) (T, error) {
	var res T

	if err := parseValue(&res, raw, p); err != nil {
		var zero T

		return zero, fmt.Errorf("%w: %w", ErrInvalidValue, err)
	}

	return res, nil
}

// parseValue parses raw into the value that dst points to.
//
//nolint:gocyclo,funlen // Flat dispatch: one case per supported type and its slice.
func parseValue(dst any, raw string, p params) error {
	var err error

	switch dst.(type) {
	case *string, *[]string:
		err = parseInto(dst, raw, p, parseString)
	case *int, *[]int:
		err = parseInto(dst, raw, p, parseInt[int])
	case *int8, *[]int8:
		err = parseInto(dst, raw, p, parseInt[int8])
	case *int16, *[]int16:
		err = parseInto(dst, raw, p, parseInt[int16])
	case *int32, *[]int32:
		err = parseInto(dst, raw, p, parseInt[int32])
	case *int64, *[]int64:
		err = parseInto(dst, raw, p, parseInt[int64])
	case *uint, *[]uint:
		err = parseInto(dst, raw, p, parseUint[uint])
	case *uint8, *[]uint8:
		err = parseInto(dst, raw, p, parseUint[uint8])
	case *uint16, *[]uint16:
		err = parseInto(dst, raw, p, parseUint[uint16])
	case *uint32, *[]uint32:
		err = parseInto(dst, raw, p, parseUint[uint32])
	case *uint64, *[]uint64:
		err = parseInto(dst, raw, p, parseUint[uint64])
	case *uintptr, *[]uintptr:
		err = parseInto(dst, raw, p, parseUint[uintptr])
	case *float32, *[]float32:
		err = parseInto(dst, raw, p, parseFloat[float32])
	case *float64, *[]float64:
		err = parseInto(dst, raw, p, parseFloat[float64])
	case *complex64, *[]complex64:
		err = parseInto(dst, raw, p, parseComplex[complex64])
	case *complex128, *[]complex128:
		err = parseInto(dst, raw, p, parseComplex[complex128])
	case *bool, *[]bool:
		err = parseInto(dst, raw, p, strconv.ParseBool)
	case *time.Time, *[]time.Time:
		err = parseInto(dst, raw, p, parseTime(p.layout))
	case *time.Duration, *[]time.Duration:
		err = parseInto(dst, raw, p, time.ParseDuration)
	case *url.URL, *[]url.URL:
		err = parseInto(dst, raw, p, parseURL)
	case *net.IP, *[]net.IP:
		err = parseInto(dst, raw, p, parseIP)
	case *netip.Addr, *[]netip.Addr:
		err = parseInto(dst, raw, p, netip.ParseAddr)
	case *netip.Prefix, *[]netip.Prefix:
		err = parseInto(dst, raw, p, netip.ParsePrefix)
	case *net.HardwareAddr, *[]net.HardwareAddr:
		err = parseInto(dst, raw, p, net.ParseMAC)
	default:
		// Unreachable from parseAs while every Parsable type has a case above.
		panic("getenv: no parser for " + reflect.TypeOf(dst).String())
	}

	return err
}

// parseInto parses raw into dst, which points to either E or []E.
func parseInto[E any](dst any, raw string, p params, parse func(string) (E, error)) error {
	var err error

	switch v := dst.(type) {
	case *E:
		*v, err = parse(raw)
	case *[]E:
		*v, err = parseSlice(raw, p.separator, parse)
	default:
		// Unreachable while every case of parseValue passes the parser of its own type.
		panic("getenv: parser does not match the destination type")
	}

	return err
}

// parseSlice splits raw by sep and parses every element with parse.
func parseSlice[T any](raw, sep string, parse func(string) (T, error)) ([]T, error) {
	if sep == "" {
		return nil, errEmptySeparator
	}

	parts := strings.Split(raw, sep)
	res := make([]T, 0, len(parts))

	for i, s := range parts {
		v, err := parse(s)
		if err != nil {
			return nil, fmt.Errorf("element %d: %w", i, err)
		}

		res = append(res, v)
	}

	return res, nil
}

func parseString(s string) (string, error) {
	return s, nil
}

func parseInt[T signed](s string) (T, error) {
	v, err := strconv.ParseInt(s, 10, bitSize[T]())
	if err != nil {
		return 0, err
	}

	return T(v), nil
}

func parseUint[T unsigned](s string) (T, error) {
	v, err := strconv.ParseUint(s, 10, bitSize[T]())
	if err != nil {
		return 0, err
	}

	return T(v), nil
}

func parseFloat[T float](s string) (T, error) {
	v, err := strconv.ParseFloat(s, bitSize[T]())
	if err != nil {
		return 0, err
	}

	return T(v), nil
}

func parseComplex[T complexNumber](s string) (T, error) {
	v, err := strconv.ParseComplex(s, bitSize[T]())
	if err != nil {
		return 0, err
	}

	return T(v), nil
}

// bitSize returns the size of T in bits, as the strconv parsers expect it.
func bitSize[T signed | unsigned | float | complexNumber]() int {
	return reflect.TypeFor[T]().Bits()
}

func parseTime(layout string) func(string) (time.Time, error) {
	return func(s string) (time.Time, error) {
		return time.Parse(layout, s)
	}
}

func parseURL(s string) (url.URL, error) {
	u, err := url.Parse(s)
	if err != nil {
		// url.Error quotes the whole input, which may hold credentials.
		if uerr, ok := errors.AsType[*url.Error](err); ok {
			err = uerr.Err
		}

		return url.URL{}, err
	}

	return *u, nil
}

func parseIP(s string) (net.IP, error) {
	ip := net.ParseIP(s)
	if ip == nil {
		return nil, fmt.Errorf("invalid IP address %q", s)
	}

	return ip, nil
}
