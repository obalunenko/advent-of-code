package internal

import (
	"fmt"
	"net"
	"net/netip"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	decimalBase = 10
	bitSize8    = 8
	bitSize16   = 16
	bitSize32   = 32
	bitSize64   = 64
	bitSize128  = 128
)

func getString(key string) (string, error) {
	env, ok := os.LookupEnv(key)
	if !ok || env == "" {
		return "", newErrNotSet(fmt.Sprintf("%q", key))
	}

	return env, nil
}

func getBool(key string) (bool, error) {
	env, err := getString(key)
	if err != nil {
		return false, err
	}

	val, err := strconv.ParseBool(env)
	if err != nil {
		return false, newErrInvalidValue(err.Error())
	}

	return val, nil
}

func getBoolSlice(key, sep string) ([]bool, error) {
	env, err := getString(key)
	if err != nil {
		return nil, err
	}

	if sep == "" {
		return nil, ErrInvalidValue
	}

	val := strings.Split(env, sep)

	b := make([]bool, 0, len(val))

	for _, s := range val {
		v, err := strconv.ParseBool(s)
		if err != nil {
			return nil, newErrInvalidValue(err.Error())
		}

		b = append(b, v)
	}

	return b, nil
}

func getStringSlice(key, sep string) ([]string, error) {
	env, err := getString(key)
	if err != nil {
		return nil, err
	}

	if sep == "" {
		return nil, ErrInvalidValue
	}

	val := strings.Split(env, sep)

	return val, nil
}

func parseNumberGen[T Number](raw string) (T, error) {
	var zero T

	switch any(zero).(type) {
	case int:
		return parseSignedNumber[T](raw, decimalBase, strconv.IntSize)
	case int8:
		return parseSignedNumber[T](raw, decimalBase, bitSize8)
	case int16:
		return parseSignedNumber[T](raw, decimalBase, bitSize16)
	case int32:
		return parseSignedNumber[T](raw, decimalBase, bitSize32)
	case int64:
		return parseSignedNumber[T](raw, decimalBase, bitSize64)
	case uint:
		return parseUnsignedNumber[T](raw, decimalBase, strconv.IntSize)
	case uint8:
		return parseUnsignedNumber[T](raw, decimalBase, bitSize8)
	case uint16:
		return parseUnsignedNumber[T](raw, decimalBase, bitSize16)
	case uint32:
		return parseUnsignedNumber[T](raw, decimalBase, bitSize32)
	case uint64:
		return parseUnsignedNumber[T](raw, decimalBase, bitSize64)
	case uintptr:
		return parseUnsignedNumber[T](raw, decimalBase, strconv.IntSize)
	case float32:
		return parseFloatNumber[T](raw, bitSize32)
	case float64:
		return parseFloatNumber[T](raw, bitSize64)
	default:
		return zero, ErrInvalidValue
	}
}

func parseSignedNumber[T Number](raw string, base, bits int) (T, error) {
	var zero T

	val, err := strconv.ParseInt(raw, base, bits)
	if err != nil {
		return zero, newErrInvalidValue(err.Error())
	}

	return T(val), nil
}

func parseUnsignedNumber[T Number](raw string, base, bits int) (T, error) {
	var zero T

	val, err := strconv.ParseUint(raw, base, bits)
	if err != nil {
		return zero, newErrInvalidValue(err.Error())
	}

	return T(val), nil
}

func parseFloatNumber[T Number](raw string, bits int) (T, error) {
	var zero T

	val, err := strconv.ParseFloat(raw, bits)
	if err != nil {
		return zero, newErrInvalidValue(err.Error())
	}

	return T(val), nil
}

func parseNumberSliceGen[T Number](raw []string) ([]T, error) {
	var zero []T

	val := make([]T, 0, len(raw))

	for _, s := range raw {
		v, err := parseNumberGen[T](s)
		if err != nil {
			return zero, err
		}

		val = append(val, v)
	}

	return val, nil
}

func getNumberSliceGen[T Number](key, sep string) ([]T, error) {
	env, err := getStringSlice(key, sep)
	if err != nil {
		return nil, err
	}

	return parseNumberSliceGen[T](env)
}

func getNumberGen[T Number](key string) (T, error) {
	env, err := getString(key)
	if err != nil {
		return 0, err
	}

	return parseNumberGen[T](env)
}

func getDuration(key string) (time.Duration, error) {
	env, err := getString(key)
	if err != nil {
		return 0, err
	}

	val, err := time.ParseDuration(env)
	if err != nil {
		return 0, newErrInvalidValue(err.Error())
	}

	return val, nil
}

func getTime(key, layout string) (time.Time, error) {
	env, err := getString(key)
	if err != nil {
		return time.Time{}, err
	}

	val, err := time.Parse(layout, env)
	if err != nil {
		return time.Time{}, newErrInvalidValue(err.Error())
	}

	return val, nil
}

func getTimeSlice(key, layout, sep string) ([]time.Time, error) {
	env, err := getStringSlice(key, sep)
	if err != nil {
		return nil, err
	}

	val := make([]time.Time, 0, len(env))

	for _, s := range env {
		v, err := time.Parse(layout, s)
		if err != nil {
			return nil, newErrInvalidValue(err.Error())
		}

		val = append(val, v)
	}

	return val, nil
}

func getDurationSlice(key, sep string) ([]time.Duration, error) {
	env, err := getStringSlice(key, sep)
	if err != nil {
		return nil, err
	}

	val := make([]time.Duration, 0, len(env))

	for _, s := range env {
		v, err := time.ParseDuration(s)
		if err != nil {
			return nil, newErrInvalidValue(err.Error())
		}

		val = append(val, v)
	}

	return val, nil
}

func getURL(key string) (url.URL, error) {
	env, err := getString(key)
	if err != nil {
		return url.URL{}, err
	}

	val, err := url.Parse(env)
	if err != nil {
		return url.URL{}, newErrInvalidValue(err.Error())
	}

	return *val, nil
}

func getURLSlice(key, sep string) ([]url.URL, error) {
	env, err := getStringSlice(key, sep)
	if err != nil {
		return nil, err
	}

	val := make([]url.URL, 0, len(env))

	for _, s := range env {
		v, err := url.Parse(s)
		if err != nil {
			return nil, newErrInvalidValue(err.Error())
		}

		val = append(val, *v)
	}

	return val, nil
}

func getIP(key string) (net.IP, error) {
	env, err := getString(key)
	if err != nil {
		return nil, err
	}

	val := net.ParseIP(env)
	if val == nil {
		return nil, ErrInvalidValue
	}

	return val, nil
}

func getIPSlice(key, sep string) ([]net.IP, error) {
	env, err := getStringSlice(key, sep)
	if err != nil {
		return nil, err
	}

	val := make([]net.IP, 0, len(env))

	for _, s := range env {
		v := net.ParseIP(s)
		if v == nil {
			return nil, ErrInvalidValue
		}

		val = append(val, v)
	}

	return val, nil
}

func getNetIPAddr(key string) (netip.Addr, error) {
	env, err := getString(key)
	if err != nil {
		return netip.Addr{}, err
	}

	val, err := netip.ParseAddr(env)
	if err != nil {
		return netip.Addr{}, newErrInvalidValue(err.Error())
	}

	return val, nil
}

func getNetIPAddrSlice(key, sep string) ([]netip.Addr, error) {
	env, err := getStringSlice(key, sep)
	if err != nil {
		return nil, err
	}

	val := make([]netip.Addr, 0, len(env))

	for _, s := range env {
		v, err := netip.ParseAddr(s)
		if err != nil {
			return nil, newErrInvalidValue(err.Error())
		}

		val = append(val, v)
	}

	return val, nil
}

func getNetIPPrefix(key string) (netip.Prefix, error) {
	env, err := getString(key)
	if err != nil {
		return netip.Prefix{}, err
	}

	val, err := netip.ParsePrefix(env)
	if err != nil {
		return netip.Prefix{}, newErrInvalidValue(err.Error())
	}

	return val, nil
}

func getNetIPPrefixSlice(key, sep string) ([]netip.Prefix, error) {
	env, err := getStringSlice(key, sep)
	if err != nil {
		return nil, err
	}

	val := make([]netip.Prefix, 0, len(env))

	for _, s := range env {
		v, err := netip.ParsePrefix(s)
		if err != nil {
			return nil, newErrInvalidValue(err.Error())
		}

		val = append(val, v)
	}

	return val, nil
}

func getHardwareAddr(key string) (net.HardwareAddr, error) {
	env, err := getString(key)
	if err != nil {
		return nil, err
	}

	val, err := net.ParseMAC(env)
	if err != nil {
		return nil, newErrInvalidValue(err.Error())
	}

	return val, nil
}

func getHardwareAddrSlice(key, sep string) ([]net.HardwareAddr, error) {
	env, err := getStringSlice(key, sep)
	if err != nil {
		return nil, err
	}

	val := make([]net.HardwareAddr, 0, len(env))

	for _, s := range env {
		v, err := net.ParseMAC(s)
		if err != nil {
			return nil, newErrInvalidValue(err.Error())
		}

		val = append(val, v)
	}

	return val, nil
}

func parseComplexGen[T Complex](raw string) (T, error) {
	var zero T

	val, err := strconv.ParseComplex(raw, complexBits[T]())
	if err != nil {
		return zero, newErrInvalidValue(err.Error())
	}

	return T(val), nil
}

func complexBits[T Complex]() int {
	var zero T

	switch any(zero).(type) {
	case complex64:
		return bitSize64
	case complex128:
		return bitSize128
	default:
		return 0
	}
}

func parseComplexSliceGen[T Complex](raw []string) ([]T, error) {
	var zero []T

	val := make([]T, 0, len(raw))

	for _, s := range raw {
		v, err := parseComplexGen[T](s)
		if err != nil {
			return zero, err
		}

		val = append(val, v)
	}

	return val, nil
}

func getComplexSliceGen[T Complex](key, sep string) ([]T, error) {
	env, err := getStringSlice(key, sep)
	if err != nil {
		return nil, err
	}

	return parseComplexSliceGen[T](env)
}

func getComplexGen[T Complex](key string) (T, error) {
	env, err := getString(key)
	if err != nil {
		return 0, err
	}

	return parseComplexGen[T](env)
}
