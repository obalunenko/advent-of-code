package day01

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solution_Year(t *testing.T) {
	var s solution

	want := "2018"
	got := s.Year()

	assert.Equal(t, want, got)
}

func Test_solution_Day(t *testing.T) {
	var s solution

	want := "day01"
	got := s.Day()

	assert.Equal(t, want, got)
}

func Test_solution_Part1(t *testing.T) {
	var s solution

	type args struct {
		input io.Reader
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "+1, +1, +1` results in  `3`",
			args: args{
				input: strings.NewReader("+1\n+1\n+1"),
			},
			want:    "3",
			wantErr: false,
		},
		{
			name: "`+1, +1, -2` results in  `0`",
			args: args{
				input: strings.NewReader("+1\n+1\n-2"),
			},
			want:    "0",
			wantErr: false,
		},
		{
			name: "`-1, -2, -3` results in `-6`",
			args: args{
				input: strings.NewReader("-1\n-2\n-3"),
			},
			want:    "-6",
			wantErr: false,
		},
		{
			name: "`+1, -2, +3, +1` results in `3`",
			args: args{
				input: strings.NewReader("+1\n-2\n+3\n+1"),
			},
			want:    "3",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Part1(tt.args.input)
			if tt.wantErr {
				assert.Error(t, err)

				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_solution_Part2(t *testing.T) {
	var s solution

	type args struct {
		input io.Reader
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "`+1, -2, +3, +1` results in `2`",
			args: args{
				input: strings.NewReader("+1\n-2\n+3\n+1"),
			},
			want:    "2",
			wantErr: false,
		},
		{
			name: "`+1, -1` first reaches `0` twice.",
			args: args{
				input: strings.NewReader("+1\n-1"),
			},
			want:    "0",
			wantErr: false,
		},
		{
			name: "`+3, +3, +4, -2, -4` first reaches `10` twice.",
			args: args{
				input: strings.NewReader("+3\n+3\n+4\n-2\n-4"),
			},
			want:    "10",
			wantErr: false,
		},
		{
			name: "`-6, +3, +8, +5, -6` first reaches `5` twice.",
			args: args{
				input: strings.NewReader("-6\n+3\n+8\n+5\n-6"),
			},
			want:    "5",
			wantErr: false,
		},
		{
			name: "`+7, +7, -2, -7, -4` first reaches `14` twice.",
			args: args{
				input: strings.NewReader("+7\n+7\n-2\n-7\n-4"),
			},
			want:    "14",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Part2(tt.args.input)
			if tt.wantErr {
				assert.Error(t, err)

				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getFreqDelta(t *testing.T) {
	type args struct {
		line string
	}

	tests := []struct {
		name    string
		args    args
		want    freqDelta
		wantErr bool
	}{
		{
			name: "",
			args: args{
				line: "+2",
			},
			want: freqDelta{
				sign: "+",
				d:    2,
			},
			wantErr: false,
		},
		{
			name: "",
			args: args{
				line: "2",
			},
			want: freqDelta{
				sign: "",
				d:    0,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getFreqDelta(tt.args.line)
			if tt.wantErr {
				assert.Error(t, err)

				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
