package program

import "testing"

func Test_longestIncreasingPath(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				matrix: [][]int{
					{3, 4, 5}, {3, 2, 6}, {2, 2, 1},
				},
			},
			want: 4,
		},

		{
			name: "test 1",
			args: args{
				matrix: [][]int{
					{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
					{19, 18, 17, 16, 15, 14, 13, 12, 11, 10},
					{20, 21, 22, 23, 24, 25, 26, 27, 28, 29},
					{39, 38, 37, 36, 35, 34, 33, 32, 31, 30},
					{40, 41, 42, 43, 44, 45, 46, 47, 48, 49},
					{59, 58, 57, 56, 55, 54, 53, 52, 51, 50},
					{60, 61, 62, 63, 64, 65, 66, 67, 68, 69},
					{79, 78, 77, 76, 75, 74, 73, 72, 71, 70},
					{80, 81, 82, 83, 84, 85, 86, 87, 88, 89},
					{99, 98, 97, 96, 95, 94, 93, 92, 91, 90},
					{100, 101, 102, 103, 104, 105, 106, 107, 108, 109},
					{119, 118, 117, 116, 115, 114, 113, 112, 111, 110},
					{120, 121, 122, 123, 124, 125, 126, 127, 128, 129},
					{139, 138, 137, 136, 135, 134, 133, 132, 131, 130},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				},
			},
			want: 140,
		},
		{
			name: "test 1",
			args: args{
				matrix: [][]int{
					{3, 15, 8, 4, 10},
					{14, 9, 14, 2, 7},
					{15, 5, 12, 10, 0},
					{13, 7, 12, 12, 18},
					{2, 3, 1, 6, 18},
				},
			},
			want: 5,
		},
		{
			name: "test 6",
			args: args{
				matrix: [][]int{
					{13, 6, 16, 6, 16, 4},
					{9, 13, 5, 13, 7, 11},
					{11, 7, 9, 17, 0, 7},
					{7, 8, 5, 14, 11, 8},
					{14, 2, 8, 7, 9, 5},
					{1, 15, 3, 11, 11, 6},
				},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestIncreasingPath(tt.args.matrix); got != tt.want {
				t.Errorf("longestIncreasingPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
