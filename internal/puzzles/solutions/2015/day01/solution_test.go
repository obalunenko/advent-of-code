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
			name: "",
			fields: fields{
				name: "",
			},
			args: args{
				input: strings.NewReader("(())"),
			},
			want:    "0",
			wantErr: false,
		},
		{
			name: "",
			fields: fields{
				name: "",
			},
			args: args{
				input: strings.NewReader("()()"),
			},
			want:    "0",
			wantErr: false,
		},
		{
			name: "",
			fields: fields{
				name: "",
			},
			args: args{
				input: strings.NewReader("((("),
			},
			want:    "3",
			wantErr: false,
		},
		{
			name: "",
			fields: fields{
				name: "",
			},
			args: args{
				input: strings.NewReader("(()(()("),
			},
			want:    "3",
			wantErr: false,
		},
		{
			name: "",
			fields: fields{
				name: "",
			},
			args: args{
				input: strings.NewReader("))((((("),
			},
			want:    "3",
			wantErr: false,
		},
		{
			name: "",
			fields: fields{
				name: "",
			},
			args: args{
				input: strings.NewReader("())"),
			},
			want:    "-1",
			wantErr: false,
		},
		{
			name: "",
			fields: fields{
				name: "",
			},
			args: args{
				input: strings.NewReader("))("),
			},
			want:    "-1",
			wantErr: false,
		},
		{
			name: "",
			fields: fields{
				name: "",
			},
			args: args{
				input: strings.NewReader(")))"),
			},
			want:    "-3",
			wantErr: false,
		},
		{
			name: "",
			fields: fields{
				name: "",
			},
			args: args{
				input: strings.NewReader(")())())"),
			},
			want:    "-3",
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
			name: "",
			fields: fields{
				name: "",
			},
			args: args{
				input: strings.NewReader(")"),
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "",
			fields: fields{
				name: "",
			},
			args: args{
				input: strings.NewReader("()())"),
			},
			want:    "5",
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
