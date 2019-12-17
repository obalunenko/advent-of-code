package day02

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newComputer(t *testing.T) {
	type args struct {
		input io.Reader
	}

	tests := []struct {
		name    string
		args    args
		want    computer
		wantErr bool
	}{
		{
			name: "",
			args: args{
				input: strings.NewReader("1,9,10,3,2,3,11,0,99,30,40,50"),
			},
			want: computer{input: map[int]int{
				0:  1,
				1:  9,
				2:  10,
				3:  3,
				4:  2,
				5:  3,
				6:  11,
				7:  0,
				8:  99,
				9:  30,
				10: 40,
				11: 50,
			}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			got, err := newComputer(tt.args.input)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.EqualValues(t, tt.want.input, got.input)
		})
	}
}

func initComp(tb testing.TB, reader io.Reader) computer {
	c, err := newComputer(reader)
	require.NoError(tb, err)

	return c
}

func Test_computer_add(t *testing.T) {
	type args struct {
		aPos   int
		bPos   int
		resPos int
	}

	tests := []struct {
		name     string
		c        computer
		args     args
		expected computer
	}{
		{
			name: "",
			c:    initComp(t, strings.NewReader("1,9,10,3,2,3,11,0,99,30,40,50")),
			args: args{
				aPos:   9,
				bPos:   10,
				resPos: 3,
			},
			expected: initComp(t, strings.NewReader("1,9,10,70,2,3,11,0,99,30,40,50")),
		},
		{
			name: "",
			c:    initComp(t, strings.NewReader("1,0,0,0,99")),
			args: args{
				aPos:   0,
				bPos:   0,
				resPos: 0,
			},
			expected: initComp(t, strings.NewReader("2,0,0,0,99")),
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			err := tt.c.add(tt.args.aPos, tt.args.bPos, tt.args.resPos)
			require.NoError(t, err)

			assert.EqualValues(t, tt.expected, tt.c)
		})
	}
}

func Test_computer_mult(t *testing.T) {
	type args struct {
		aPos   int
		bPos   int
		resPos int
	}

	tests := []struct {
		name     string
		c        computer
		args     args
		expected computer
	}{
		{
			name: "",
			c:    initComp(t, strings.NewReader("1,9,10,70,2,3,11,0,99,30,40,50")),
			args: args{
				aPos:   3,
				bPos:   11,
				resPos: 0,
			},
			expected: initComp(t, strings.NewReader("3500,9,10,70,2,3,11,0,99,30,40,50")),
		},
		{
			name: "",
			c:    initComp(t, strings.NewReader("2,3,0,3,99")),
			args: args{
				aPos:   3,
				bPos:   0,
				resPos: 3,
			},
			expected: initComp(t, strings.NewReader("2,3,0,6,99")),
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			err := tt.c.mult(tt.args.aPos, tt.args.bPos, tt.args.resPos)
			require.NoError(t, err)

			assert.EqualValues(t, tt.expected, tt.c)
		})
	}
}

func Test_computer_calc(t *testing.T) {
	tests := []struct {
		name     string
		c        computer
		want     int
		wantComp computer
	}{
		{
			name:     "",
			c:        initComp(t, strings.NewReader("1,9,10,3,2,3,11,0,99,30,40,50")),
			want:     3500,
			wantComp: initComp(t, strings.NewReader("3500,9,10,70,2,3,11,0,99,30,40,50")),
		},
		{
			name:     "",
			c:        initComp(t, strings.NewReader("1,1,1,4,99,5,6,0,99")),
			want:     30,
			wantComp: initComp(t, strings.NewReader("30,1,1,4,2,5,6,0,99")),
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			comp := tt.c
			got, err := comp.calc()

			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
			assert.EqualValues(t, tt.wantComp, comp)
		})
	}
}

func Test_computer_replace(t *testing.T) {
	type args struct {
		data map[int]int
	}

	tests := []struct {
		name     string
		c        computer
		args     args
		wantComp computer
	}{
		{
			name: "",
			c:    initComp(t, strings.NewReader("1,1,1,4,99,5,6,0,99")),
			args: args{
				data: map[int]int{
					1: 9,
					2: 12,
				},
			},
			wantComp: initComp(t, strings.NewReader("1,9,12,4,99,5,6,0,99")),
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			comp := tt.c
			comp.replace(tt.args.data)
			assert.EqualValues(t, tt.wantComp, comp)
		})
	}
}