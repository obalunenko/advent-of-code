package day02

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_nounVerb(t *testing.T) {
	type args struct {
		noun int
		verb int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				noun: 12,
				verb: 2,
			},
			want: 1202,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			got := nounVerb(tt.args.noun, tt.args.verb)
			assert.Equal(t, tt.want, got)
		})
	}
}

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
				input: strings.NewReader("1,9,10,3,2,3,11,0,99,30,40,50,2,3"),
			},
			want:    "200",
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
