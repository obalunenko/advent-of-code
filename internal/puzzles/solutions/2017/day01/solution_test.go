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
			name: `1122 produces a sum of 3 (1 + 2) because the first digit (1) matches the second digit and the third digit (2) matches the fourth digit`,
			fields: fields{
				name: "",
				year: "",
			},
			args: args{
				input: strings.NewReader("1122"),
			},
			want:    "3",
			wantErr: false,
		},
		{
			name: `1111 produces 4 because each digit (all 1) matches the next.`,
			fields: fields{
				name: "",
				year: "",
			},
			args: args{
				input: strings.NewReader("1111"),
			},
			want:    "4",
			wantErr: false,
		},
		{
			name: `1234 produces 0 because no digit matches the next`,
			fields: fields{
				name: "",
				year: "",
			},
			args: args{
				input: strings.NewReader("1234"),
			},
			want:    "0",
			wantErr: false,
		},
		{
			name: `91212129 produces 9 because the only digit that matches the next one is the last digit, 9`,
			fields: fields{
				name: "",
				year: "",
			},
			args: args{
				input: strings.NewReader("91212129"),
			},
			want:    "9",
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

func Test_solution_Part2(t *testing.T) {
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
			name: "`1212` produces `6`: the list contains `4` items, and all four digits match the digit `2` items ahead",
			fields: fields{
				name: "",
			},
			args: args{
				input: strings.NewReader("1212"),
			},
			want:    "6",
			wantErr: false,
		},
		{
			name: "`1221` produces `0`, because every comparison is between a `1` and a `2`",
			fields: fields{
				name: "",
			},
			args: args{
				input: strings.NewReader("1221"),
			},
			want:    "0",
			wantErr: false,
		},
		{
			name: "`123425` produces `4`, because both `2`s match each other, but no other digit has a match",
			fields: fields{
				name: "",
			},
			args: args{
				input: strings.NewReader("123425"),
			},
			want:    "4",
			wantErr: false,
		},
		{
			name: "`123123` produces `12`",
			fields: fields{
				name: "",
			},
			args: args{
				input: strings.NewReader("123123"),
			},
			want:    "12",
			wantErr: false,
		},
		{
			name: "`12131415` produces `4`",
			fields: fields{
				name: "",
			},
			args: args{
				input: strings.NewReader("12131415"),
			},
			want:    "4",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			s := solution{
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
