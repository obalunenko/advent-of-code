package puzzles

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResult_String(t *testing.T) {
	type fields struct {
		Year    string
		Name    string
		Part1   string
		Part2   string
		Metrics metrics
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "",
			fields: fields{
				Year:    "2020",
				Name:    "day01",
				Part1:   "12",
				Part2:   "10",
				Metrics: nil,
			},
			want: `
   2020/day01 puzzle answer:   
      part1                    12   
      part2                    10
`,
		},
		{
			name: "",
			fields: fields{
				Year:    "",
				Name:    "",
				Part1:   "",
				Part2:   "",
				Metrics: nil,
			},
			want: `
   unknown/unknown puzzle answer:   
      part1                         not solved   
      part2                         not solved
`,
		},
		{
			name: "",
			fields: fields{
				Year:  "2020",
				Name:  "day01",
				Part1: "12",
				Part2: "10",
				Metrics: metrics{
					{
						mType:    metricsFlagNone,
						metadata: "",
					},
				},
			},
			want: `
   2020/day01 puzzle answer:   
      part1                    12   
      part2                    10   
   metrics:                    
   none
`,
		},
		{
			name: "",
			fields: fields{
				Year:  "2020",
				Name:  "day01",
				Part1: "12",
				Part2: "10",
				Metrics: metrics{
					{
						mType:    metricsFlagElapsed,
						metadata: "12s",
					},
				},
			},
			want: `
   2020/day01 puzzle answer:   
      part1                    12   
      part2                    10   
   metrics:                    
      elapsed                  12s
`,
		},
		{
			name: "",
			fields: fields{
				Year:  "2020",
				Name:  "day01",
				Part1: "12",
				Part2: "10",
				Metrics: metrics{
					{
						mType:    metricsFlagBenchmark,
						metadata: "1600 alloc",
					},
				},
			},
			want: `
   2020/day01 puzzle answer:   
      part1                    12   
      part2                    10   
   metrics:                    
      benchmark                1600 alloc
`,
		},
		{
			name: "",
			fields: fields{
				Year:  "2020",
				Name:  "day01",
				Part1: "122",
				Part2: "10",
				Metrics: metrics{
					{
						mType:    metricsFlagBenchmark,
						metadata: "1600 alloc",
					},
					{
						mType:    metricsFlagElapsed,
						metadata: "16s",
					},
				},
			},
			want: `
   2020/day01 puzzle answer:   
      part1                    122   
      part2                    10    
   metrics:                    
      benchmark                1600 alloc   
      elapsed                  16s
`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d:%s", i, tt.name), func(t *testing.T) {
			r := Result{
				Year:    tt.fields.Year,
				Name:    tt.fields.Name,
				Part1:   tt.fields.Part1,
				Part2:   tt.fields.Part2,
				metrics: tt.fields.Metrics,
			}

			got := r.String()

			assert.Equal(t, tt.want, got)
		})
	}
}
