package day03

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solution_Year(t *testing.T) {
	var s solution

	want := "2021"
	got := s.Year()

	assert.Equal(t, want, got)
}

func Test_solution_Day(t *testing.T) {
	var s solution

	want := "3"
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
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "test example from description",
			args: args{
				input: strings.NewReader("00100\n11110\n10110\n10111\n10101\n01111\n00111\n11100\n10000\n11001\n00010\n01010\n"),
			},
			want:    "198",
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Part1(tt.args.input)
			if !tt.wantErr(t, err, fmt.Sprintf("s.Part1()")) {
				return
			}

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
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "test example from description",
			args: args{
				input: strings.NewReader(""),
			},
			want:    "",
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Part2(tt.args.input)
			if !tt.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_findRates(t *testing.T) {
	type args struct {
		diagnostic []string
	}
	tests := []struct {
		name string
		args args
		want bitrates
	}{
		{
			name: "",
			args: args{
				diagnostic: []string{
					"00100",
					"11110",
					"10110",
					"10111",
					"10101",
					"01111",
					"00111",
					"11100",
					"10000",
					"11001",
					"00010",
					"01010",
				},
			},
			want: bitrates{
				gamma:   "10110",
				epsilon: "01001",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findRates(tt.args.diagnostic), "findGammaRate(%v)", tt.args.diagnostic)
		})
	}
}

func Test_bitrates_consumption(t *testing.T) {
	type fields struct {
		gamma   string
		epsilon string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "",
			fields: fields{
				gamma:   "10110",
				epsilon: "01001",
			},
			want:    "198",
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bitrates{
				gamma:   tt.fields.gamma,
				epsilon: tt.fields.epsilon,
			}
			got, err := b.consumption()
			if !tt.wantErr(t, err, fmt.Sprintf("consumption()")) {
				return
			}

			assert.Equalf(t, tt.want, got, "consumption()")
		})
	}
}
