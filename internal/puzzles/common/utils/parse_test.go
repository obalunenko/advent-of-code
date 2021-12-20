package utils

import (
	"errors"
	"io"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
)

func TestParseInts(t *testing.T) {
	type args struct {
		in  io.Reader
		sep string
	}

	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "",
			args: args{
				in:  strings.NewReader("1,2,3,4,5"),
				sep: ",",
			},
			want:    []int{1, 2, 3, 4, 5},
			wantErr: assert.NoError,
		},
		{
			name: "",
			args: args{
				in:  strings.NewReader("1\n2\n3\n4\n5"),
				sep: "",
			},
			want:    []int{1, 2, 3, 4, 5},
			wantErr: assert.NoError,
		},
		{
			name: "",
			args: args{
				in:  strings.NewReader("1,2\n3,4\n5"),
				sep: ",",
			},
			want:    []int{1, 2, 3, 4, 5},
			wantErr: assert.NoError,
		},
		{
			name: "",
			args: args{
				in:  iotest.ErrReader(errors.New("custom error")),
				sep: ",",
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "broken int",
			args: args{
				in:  strings.NewReader("1s,2\n3,4\n5"),
				sep: ",",
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseInts(tt.args.in, tt.args.sep)
			if !tt.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
