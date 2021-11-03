package day01

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solution_Part1(t *testing.T) {
	type fields struct {
		name string
		year string
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
			name: "+1, +1, +1` results in  `3`",
			fields: fields{
				name: "day01",
				year: "2018",
			},
			args: args{
				input: strings.NewReader("+1, +1, +1"),
			},
			want:    "3",
			wantErr: false,
		},
		{
			name: "`+1, +1, -2` results in  `0`",
			fields: fields{
				name: "day01",
				year: "2018",
			},
			args: args{
				input: strings.NewReader("+1, +1, -2"),
			},
			want:    "0",
			wantErr: false,
		},
		{
			name: "`1, -2, -3` results in `-6`",
			fields: fields{
				name: "day01",
				year: "2018",
			},
			args: args{
				input: strings.NewReader("1, -2, -3"),
			},
			want:    "-6",
			wantErr: false,
		},
		{
			name: "`+1, -2, +3, +1` results in `3`",
			fields: fields{
				name: "",
			},
			args: args{
				input: strings.NewReader("+1, -2, +3, +1"),
			},
			want:    "12",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			s := solution{
				year: tt.fields.year,
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

func Test_solution_Part2(t *testing.T) {
	type fields struct {
		name string
		year string
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
		{ // TODO(@obalunenko): Fill the tests
			name: "",
			fields: fields{
				name: "2018",
				year: "day01",
			},
			args: args{
				input: strings.NewReader(""),
			},
			want:    "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			s := solution{
				year: tt.fields.year,
				name: tt.fields.name,
			}

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
