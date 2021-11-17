package puzzles

import (
	"fmt"
	"strings"
	"testing"
	"text/tabwriter"
	"time"
)

type metricsFlag uint32

const (
	metricsFlagNone metricsFlag = 1 << iota
	metricsFlagElapsed
	metricsFlagBenchmark

	elapsed   = "elapsed"
	benchmark = "benchmark"
	none      = "none"
)

func (m metricsFlag) String() string {
	var s string

	if m.HasFlag(metricsFlagNone) {
		s = none
	}

	maybePfx := func(s string) string {
		if s != "" {
			return "|"
		}

		return ""
	}

	if m.HasFlag(metricsFlagElapsed) {
		s += maybePfx(s) + elapsed
	}

	if m.HasFlag(metricsFlagBenchmark) {
		s += maybePfx(s) + benchmark
	}

	return s
}

func (f metricsFlag) HasFlag(flag metricsFlag) bool { return f&flag != 0 }
func (f *metricsFlag) AddFlag(flag metricsFlag)     { *f |= flag }
func (f *metricsFlag) ClearFlag(flag metricsFlag)   { *f &= ^flag }
func (f *metricsFlag) ToggleFlag(flag metricsFlag)  { *f ^= flag }

type metrics []*metric

func (m metrics) String() string {
	var buf strings.Builder

	w := tabwriter.NewWriter(&buf, 0, 0, 1, ' ', tabwriter.TabIndent)

	_, err := fmt.Fprintf(w, `
metrics:
`)
	if err != nil {
		panic(err)
	}

	for i := range m {
		_, err := fmt.Fprintf(w, `| %s |`, m[i].String())
		if err != nil {
			panic(err)
		}
	}

	if err := w.Flush(); err != nil {
		panic(err)
	}

	return buf.String()
}

type metric struct {
	mType    metricsFlag
	metadata string
}

func (m *metric) String() string {
	if m == nil {
		return undefined
	}

	if m.mType.HasFlag(metricsFlagNone) {
		return fmt.Sprintf("[%s]", m.mType.String())
	}

	return fmt.Sprintf("%s: %s", m.mType.String(), m.metadata)
}

func (m *metric) elapsed() func() {
	start := time.Now()

	m.mType = metricsFlagElapsed
	m.metadata = inProgress

	return func() {
		m.metadata = timeElapsed(start)
	}
}

type benchFunc func() error

func (m *metric) bench(f benchFunc) func() {
	m.mType = metricsFlagBenchmark
	m.metadata = inProgress

	return func() {
		m.metadata = bench(f)
	}
}

func timeElapsed(start time.Time) string {
	elapsed := time.Since(start)

	return elapsed.String()
}

func bench(f benchFunc) string {
	b := testing.Benchmark(func(b *testing.B) {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			if err := f(); err != nil {
				b.FailNow()
			}
		}
	})

	result := fmt.Sprintf("(N=%d, %d ns/op, %d bytes/op, %d allocs/op)",
		b.N, b.NsPerOp(), b.AllocedBytesPerOp(), b.AllocsPerOp())

	return result
}
