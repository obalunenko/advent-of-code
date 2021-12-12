package day05

import (
	"fmt"
	"io"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/obalunenko/advent-of-code/internal/puzzles/common/utils"
)

func Test_solution_Year(t *testing.T) {
	var s solution

	want := "2021"
	got := s.Year()

	assert.Equal(t, want, got)
}

func Test_solution_Day(t *testing.T) {
	var s solution

	want := "5"
	got := s.Day()

	assert.Equal(t, want, got)
}

func Test_solution_Part1(t *testing.T) {
	var s solution

	type args struct {
		input io.Reader
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "test example from description",
			args: args{
				input: utils.ReaderFromFile(t, filepath.Join("testdata", "input.txt")),
			},
			want:    "5",
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Part1(tt.args.input)
			if !tt.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_solution_Part2(t *testing.T) {
	var s solution

	type args struct {
		input io.Reader
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "test example from description",
			args: args{
				input: utils.ReaderFromFile(t, filepath.Join("testdata", "input.txt")),
			},
			want:    "12",
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Part2(tt.args.input)
			if !tt.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_parseLine(t *testing.T) {
	type args struct {
		line string
	}

	tests := []struct {
		name    string
		args    args
		want    []position
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "",
			args: args{
				line: "0,9 -> 5,9",
			},
			want: []position{
				{
					x: 0,
					y: 9,
				},
				{
					x: 5,
					y: 9,
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "",
			args: args{
				line: "105,697 -> 287,697",
			},
			want: []position{
				{
					x: 105,
					y: 697,
				},
				{
					x: 287,
					y: 697,
				},
			},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseLine(tt.args.line)
			if !tt.wantErr(t, err, fmt.Sprintf("parseLine(%v)", tt.args.line)) {
				return
			}

			assert.Equalf(t, tt.want, got, "parseLine(%v)", tt.args.line)
		})
	}
}

func Test_parseCoordinates(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name    string
		args    args
		want    position
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "",
			args: args{
				s: "0,9",
			},
			want: position{
				x: 0,
				y: 9,
			},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseCoordinates(tt.args.s)
			if !tt.wantErr(t, err, fmt.Sprintf("parseCoordinates(%v)", tt.args.s)) {
				return
			}
			assert.Equalf(t, tt.want, got, "parseCoordinates(%v)", tt.args.s)
		})
	}
}

func Test_getLines(t *testing.T) {
	type args struct {
		input io.Reader
	}

	var tests = []struct {
		name    string
		args    args
		want    []line
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "",
			args: args{
				input: utils.ReaderFromFile(t, filepath.Join("testdata", "input.txt")),
			},
			want:    getTestLines(t),
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getLines(tt.args.input)
			if !tt.wantErr(t, err, fmt.Sprintf("getLines(%v)", tt.args.input)) {
				return
			}
			assert.Equalf(t, tt.want, got, "getLines(%v)", tt.args.input)
		})
	}
}

func Test_newDiagram(t *testing.T) {
	type args struct {
		maxX int
		maxY int
	}

	tests := []struct {
		name string
		args args
		want diagram
	}{
		{
			name: "",
			args: args{
				maxX: 3,
				maxY: 4,
			},
			want: diagram{
				data: [][]int{
					{0, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
				},
			},
		},
		{
			name: "",
			args: args{
				maxX: 9,
				maxY: 9,
			},
			want: diagram{
				data: [][]int{
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, newDiagram(tt.args.maxX, tt.args.maxY), "newDiagram(%v, %v)", tt.args.maxX, tt.args.maxY)
		})
	}
}

func getTestLines(tb testing.TB) []line {
	tb.Helper()

	return []line{
		{
			start: position{x: 0, y: 9},
			end:   position{x: 5, y: 9},
		},
		{
			start: position{x: 8, y: 0},
			end:   position{x: 0, y: 8},
		},
		{
			start: position{x: 9, y: 4},
			end:   position{x: 3, y: 4},
		},
		{
			start: position{x: 2, y: 2},
			end:   position{x: 2, y: 1},
		},
		{
			start: position{x: 7, y: 0},
			end:   position{x: 7, y: 4},
		},
		{
			start: position{x: 6, y: 4},
			end:   position{x: 2, y: 0},
		},
		{
			start: position{x: 0, y: 9},
			end:   position{x: 2, y: 9},
		},
		{
			start: position{x: 3, y: 4},
			end:   position{x: 1, y: 4},
		},
		{
			start: position{x: 0, y: 0},
			end:   position{x: 8, y: 8},
		},
		{
			start: position{x: 5, y: 5},
			end:   position{x: 8, y: 2},
		},
	}
}
func Test_filterLines(t *testing.T) {
	type args struct {
		lines  []line
		filter filterFunc
	}

	tests := []struct {
		name string
		args args
		want []line
	}{
		{
			name: "",
			args: args{
				lines:  getTestLines(t),
				filter: part1Filter,
			},
			want: []line{
				{
					start: position{x: 0, y: 9},
					end:   position{x: 5, y: 9},
				},
				{
					start: position{x: 9, y: 4},
					end:   position{x: 3, y: 4},
				},
				{
					start: position{x: 2, y: 2},
					end:   position{x: 2, y: 1},
				},
				{
					start: position{x: 7, y: 0},
					end:   position{x: 7, y: 4},
				},
				{
					start: position{x: 0, y: 9},
					end:   position{x: 2, y: 9},
				},
				{
					start: position{x: 3, y: 4},
					end:   position{x: 1, y: 4},
				},
			},
		},
		{
			name: "",
			args: args{
				lines:  getTestLines(t),
				filter: part2Filter,
			},
			want: []line{
				{
					start: position{x: 0, y: 9},
					end:   position{x: 5, y: 9},
				},
				{
					start: position{x: 8, y: 0},
					end:   position{x: 0, y: 8},
				},
				{
					start: position{x: 9, y: 4},
					end:   position{x: 3, y: 4},
				},
				{
					start: position{x: 2, y: 2},
					end:   position{x: 2, y: 1},
				},
				{
					start: position{x: 7, y: 0},
					end:   position{x: 7, y: 4},
				},
				{
					start: position{x: 6, y: 4},
					end:   position{x: 2, y: 0},
				},
				{
					start: position{x: 0, y: 9},
					end:   position{x: 2, y: 9},
				},
				{
					start: position{x: 3, y: 4},
					end:   position{x: 1, y: 4},
				},
				{
					start: position{x: 0, y: 0},
					end:   position{x: 8, y: 8},
				},
				{
					start: position{x: 5, y: 5},
					end:   position{x: 8, y: 2},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, filterLines(tt.args.lines, tt.args.filter), "filterLines(%v, %v)", tt.args.lines, tt.args.filter)
		})
	}
}

func Test_diagram_String(t *testing.T) {
	tests := []struct {
		name string
		d    diagram
		want string
	}{
		{
			name: "",
			d: diagram{
				data: [][]int{
					{0, 1, 2, 0, 0},
					{1, 0, 0, 0, 0},
					{0, 1, 0, 2, 0},
				},
			},
			want: ".12..\n" +
				"1....\n" +
				".1.2.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.d.String(), "String()")
		})
	}
}

func Test_getBounds(t *testing.T) {
	type args struct {
		lines []line
	}

	tests := []struct {
		name string
		args args
		want position
	}{
		{
			name: "",
			args: args{
				lines: []line{
					{
						start: position{x: 0, y: 9},
						end:   position{x: 5, y: 9},
					},
					{
						start: position{x: 9, y: 4},
						end:   position{x: 3, y: 4},
					},
					{
						start: position{x: 2, y: 2},
						end:   position{x: 2, y: 1},
					},
					{
						start: position{x: 7, y: 0},
						end:   position{x: 7, y: 4},
					},
					{
						start: position{x: 0, y: 9},
						end:   position{x: 2, y: 9},
					},
					{
						start: position{x: 3, y: 4},
						end:   position{x: 1, y: 4},
					},
				},
			},
			want: position{
				x: 9,
				y: 9,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getBounds(tt.args.lines), "getBounds(%v)", tt.args.lines)
		})
	}
}

func Test_drawDiagram(t *testing.T) {
	type args struct {
		lines []line
	}

	tests := []struct {
		name            string
		args            args
		wantDiagramPath string
	}{
		{
			name: "part 1 diagram",
			args: args{
				lines: []line{
					{
						start: position{x: 0, y: 9},
						end:   position{x: 5, y: 9},
					},
					{
						start: position{x: 9, y: 4},
						end:   position{x: 3, y: 4},
					},
					{
						start: position{x: 2, y: 2},
						end:   position{x: 2, y: 1},
					},
					{
						start: position{x: 7, y: 0},
						end:   position{x: 7, y: 4},
					},
					{
						start: position{x: 0, y: 9},
						end:   position{x: 2, y: 9},
					},
					{
						start: position{x: 3, y: 4},
						end:   position{x: 1, y: 4},
					},
				},
			},
			wantDiagramPath: filepath.Join("testdata", "diagram_part1.txt"),
		},
		{
			name: "part 2 diagram",
			args: args{
				lines: []line{
					{
						start: position{x: 0, y: 9},
						end:   position{x: 5, y: 9},
					},
					{
						start: position{x: 8, y: 0},
						end:   position{x: 0, y: 8},
					},
					{
						start: position{x: 9, y: 4},
						end:   position{x: 3, y: 4},
					},
					{
						start: position{x: 2, y: 2},
						end:   position{x: 2, y: 1},
					},
					{
						start: position{x: 7, y: 0},
						end:   position{x: 7, y: 4},
					},
					{
						start: position{x: 6, y: 4},
						end:   position{x: 2, y: 0},
					},
					{
						start: position{x: 0, y: 9},
						end:   position{x: 2, y: 9},
					},
					{
						start: position{x: 3, y: 4},
						end:   position{x: 1, y: 4},
					},
					{
						start: position{x: 0, y: 0},
						end:   position{x: 8, y: 8},
					},
					{
						start: position{x: 5, y: 5},
						end:   position{x: 8, y: 2},
					},
				},
			},
			wantDiagramPath: filepath.Join("testdata", "diagram_part2.txt"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := utils.ReaderFromFile(t, tt.wantDiagramPath)

			content, err := io.ReadAll(file)
			require.NoError(t, err)

			assert.Equalf(t, string(content), drawDiagram(tt.args.lines).String(),
				"drawDiagram(%v)", tt.args.lines)
		})
	}
}

func Test_diagram_dangerZones(t *testing.T) {
	type fields struct {
		data [][]int
	}

	type args struct {
		f dangerFunc
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "",
			fields: fields{
				data: [][]int{
					{0, 1, 2, 0, 0},
					{1, 0, 0, 0, 0},
					{0, 1, 0, 2, 0},
				},
			},
			args: args{
				f: isDangerZone,
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := diagram{
				data: tt.fields.data,
			}
			assert.Equalf(t, tt.want, d.dangerZones(tt.args.f), "dangerZones(%v)", tt.args.f)
		})
	}
}

func Test_diagram_drawDiagonal(t *testing.T) {
	type fields struct {
		diagram diagram
	}

	type args struct {
		l line
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		expected diagram
	}{
		{
			name: "8,0 -> 0,8",
			fields: fields{
				diagram: newDiagram(9, 9),
			},
			args: args{
				l: line{
					start: position{
						x: 8,
						y: 0,
					},
					end: position{
						x: 0,
						y: 8,
					},
				},
			},
			expected: diagram{
				data: [][]int{
					{0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
					{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
					{0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
					{0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
					{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
					{0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
					{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
					{0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
					{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				},
			},
		},
		{
			name: "1,1 -> 3,3",
			fields: fields{
				diagram: newDiagram(9, 9),
			},
			args: args{
				l: line{
					start: position{
						x: 1,
						y: 1,
					},
					end: position{
						x: 3,
						y: 3,
					},
				},
			},
			expected: diagram{
				data: [][]int{
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				},
			},
		},
		{
			name: "9,7 -> 7,9",
			fields: fields{
				diagram: newDiagram(9, 9),
			},
			args: args{
				l: line{
					start: position{
						x: 9,
						y: 7,
					},
					end: position{
						x: 7,
						y: 9,
					},
				},
			},
			expected: diagram{
				data: [][]int{
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, // 7
					{0, 0, 0, 0, 0, 0, 0, 0, 1, 0}, // 8
					{0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, // 9
				},
			},
		},
		{
			name: "8,0 -> 0,8",
			fields: fields{
				diagram: newDiagram(9, 9),
			},
			args: args{
				l: line{
					start: position{
						x: 8,
						y: 0,
					},
					end: position{
						x: 0,
						y: 8,
					},
				},
			},
			expected: diagram{
				data: [][]int{
					{0, 0, 0, 0, 0, 0, 0, 0, 1, 0}, // 0
					{0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, // 1
					{0, 0, 0, 0, 0, 0, 1, 0, 0, 0}, // 2
					{0, 0, 0, 0, 0, 1, 0, 0, 0, 0}, // 3
					{0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, // 4
					{0, 0, 0, 1, 0, 0, 0, 0, 0, 0}, // 5
					{0, 0, 1, 0, 0, 0, 0, 0, 0, 0}, // 6
					{0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, // 7
					{1, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // 8
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // 9
					// 0 1 2  3  4  5  6  7  8  9
				},
			},
		},
		{
			name: "6,4 -> 2,0",
			fields: fields{
				diagram: newDiagram(9, 9),
			},
			args: args{
				l: line{
					start: position{
						x: 6,
						y: 4,
					},
					end: position{
						x: 2,
						y: 0,
					},
				},
			},
			expected: diagram{
				data: [][]int{
					{0, 0, 1, 0, 0, 0, 0, 0, 0, 0}, // 0
					{0, 0, 0, 1, 0, 0, 0, 0, 0, 0}, // 1
					{0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, // 2
					{0, 0, 0, 0, 0, 1, 0, 0, 0, 0}, // 3
					{0, 0, 0, 0, 0, 0, 1, 0, 0, 0}, // 4
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // 5
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // 6
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // 7
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // 8
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // 9
					// 0 1 2  3  4  5  6  7  8  9
				},
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.diagram.drawDiagonal(tt.args.l)

			assert.Equal(t, tt.expected, tt.fields.diagram)
		})
	}
}
