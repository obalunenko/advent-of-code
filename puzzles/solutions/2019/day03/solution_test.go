package day03

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solution_Part1(t *testing.T) {
	type fields struct {
		name string
	}

	type args struct {
		input io.Reader
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "distance 6",
			fields: fields{
				name: "part1",
			},
			args: args{
				input: strings.NewReader("U7,R6,D4,L4\n" +
					"R8,U5,L5,D3"),
			},
			want:    "6",
			wantErr: false,
		},
		{
			name: "distance 159",
			fields: fields{
				name: "part1",
			},
			args: args{
				input: strings.NewReader("R75,D30,R83,U83,L12,D49,R71,U7,L72\n" +
					"U62,R66,U55,R34,D71,R55,D58,R83"),
			},
			want:    "159",
			wantErr: false,
		},
		{
			name: "distance 159",
			fields: fields{
				name: "part1",
			},
			args: args{
				input: strings.NewReader("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\n" +
					"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"),
			},
			want:    "135",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			s := solution{
				name: tt.fields.name,
			}

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

func Test_findCross(t *testing.T) {
	type args struct {
		wm1 map[pos]bool
		wm2 map[pos]bool
	}

	tests := []struct {
		name string
		args args
		want []pos
	}{
		{
			name: "",
			args: args{
				wm1: map[pos]bool{
					{x: 0, y: 1}: true,
					{x: 3, y: 1}: true,
					{x: 4, y: 1}: true,
					{x: 5, y: 2}: true,
					{x: 6, y: 3}: true,
				},
				wm2: map[pos]bool{
					{x: 0, y: 1}: true,
					{x: 4, y: 1}: true,
					{x: 6, y: 3}: true,
				},
			},
			want: []pos{
				{x: 0, y: 1},
				{x: 4, y: 1},
				{x: 6, y: 3},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			got := findCross(tt.args.wm1, tt.args.wm2)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
