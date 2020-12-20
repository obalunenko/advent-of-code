package day02

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solution_Part1(t *testing.T) {
	type fields struct {
		year string
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
			name: "test example from description",
			fields: fields{
				year: year,
				name: puzzleName,
			},
			args: args{
				input: strings.NewReader("1-3 a: abcde\n1-3 b: cdefg\n2-9 c: ccccccccc"),
			},
			want:    "2",
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
		year string
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
			name: "test example from description",
			fields: fields{
				year: year,
				name: puzzleName,
			},
			args: args{
				input: strings.NewReader("1-3 a: abcde\n1-3 b: cdefg\n2-9 c: ccccccccc"),
			},
			want:    "1",
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
