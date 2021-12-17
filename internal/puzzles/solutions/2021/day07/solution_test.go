package day07

import (
	"errors"
	"io"
	"strings"
	"testing"
	"testing/iotest"

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

	want := "7"
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
				input: strings.NewReader("16,1,2,0,4,2,7,1,2,14"),
			},
			want:    "37",
			wantErr: assert.NoError,
		},
		{
			name: "",
			args: args{
				input: iotest.ErrReader(errors.New("custom error")),
			},
			want:    "",
			wantErr: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Part1(tt.args.input)
			if !tt.wantErr(t, err) {
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
		{
			name: "",
			args: args{
				input: iotest.ErrReader(errors.New("custom error")),
			},
			want:    "",
			wantErr: assert.Error,
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

func Test_makeMatrix(t *testing.T) {
	type args struct {
		crabs []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				crabs: []int{
					1, 45, 3, 4,
				},
			},
			want: [][]int{
				{undef, 1, 45, 3, 4},
				{1, 0, 0, 0, 0},
				{45, 0, 0, 0, 0},
				{3, 0, 0, 0, 0},
				{4, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, makeMatrix(tt.args.crabs), "makeMatrix(%v)", tt.args.crabs)
		})
	}
}

func Test_swarm_calcDistances(t *testing.T) {
	type fields struct {
		crabsMatrix [][]int
	}
	tests := []struct {
		name     string
		fields   fields
		expected *swarm
	}{
		{
			name: "",
			fields: fields{
				crabsMatrix: [][]int{
					{undef, 1, 2, 3, 4},
					{1, 0, 0, 0, 0},
					{2, 0, 0, 0, 0},
					{3, 0, 0, 0, 0},
					{4, 0, 0, 0, 0},
				},
			},
			expected: &swarm{
				crabsMatrix: [][]int{
					{undef, 1, 2, 3, 4},
					{1, 0, 1, 2, 3},
					{2, 1, 0, 1, 2},
					{3, 2, 1, 0, 1},
					{4, 3, 2, 1, 0},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &swarm{
				crabsMatrix: tt.fields.crabsMatrix,
			}
			s.calcDistances()

			assert.Equal(t, tt.expected, s)
		})
	}
}

func Test_swarm_minDistanceCost(t *testing.T) {
	type fields struct {
		crabsMatrix [][]int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "",
			fields: fields{
				crabsMatrix: [][]int{
					{undef, 1, 2, 3, 4},
					{1, 0, 1, 2, 3},
					{2, 1, 0, 1, 2},
					{3, 2, 1, 0, 1},
					{4, 3, 2, 1, 0},
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := swarm{
				crabsMatrix: tt.fields.crabsMatrix,
			}

			assert.Equalf(t, tt.want, s.minDistanceCost(), "minDistanceCost()")
		})
	}
}
