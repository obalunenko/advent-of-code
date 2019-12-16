package solutions

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeName(t *testing.T) {
	type args struct {
		year   string
		puzzle string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "valid name",
			args: args{
				year:   "2019",
				puzzle: "day01",
			},
			want:    "2019" + string(os.PathSeparator) + "day01",
			wantErr: false,
		},
		{
			name: "missed year - error",
			args: args{
				year:   "",
				puzzle: "day01",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "missed puzzle - error",
			args: args{
				year:   "2019",
				puzzle: "",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MakeName(tt.args.year, tt.args.puzzle)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
