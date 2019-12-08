package fuel

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_module_fuel(t *testing.T) {
	type fields struct {
		mass int
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "mass 12",
			fields: fields{
				mass: 12,
			},
			want: 2,
		},
		{
			name: "mass 14",
			fields: fields{
				mass: 14,
			},
			want: 2,
		},
		{
			name: "mass 1969",
			fields: fields{
				mass: 1969,
			},
			want: 654,
		},
		{
			name: "mass 100756",
			fields: fields{
				mass: 100756,
			},
			want: 33583,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			m := module{
				mass: tt.fields.mass,
			}

			got := m.fuel()
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_calculate(t *testing.T) {
	type args struct {
		filepath string
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "",
			args: args{
				filepath: filepath.Join("testdata", "input.txt"),
			},
			want:    5193796,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			got, err := calculate(tt.args.filepath)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)

			t.Logf("fuel sum: %d \n", got)
		})
	}
}
