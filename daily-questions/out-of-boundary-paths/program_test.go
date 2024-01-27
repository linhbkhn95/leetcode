package outofboundarypaths

import "testing"

func Test_findPaths(t *testing.T) {
	type args struct {
		m           int
		n           int
		maxMove     int
		startRow    int
		startColumn int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "",
		// 	args: args{
		// 		m:           2,
		// 		n:           2,
		// 		maxMove:     2,
		// 		startRow:    0,
		// 		startColumn: 0,
		// 	},
		// 	want: 6,
		// },
		// {
		// 	name: "",
		// 	args: args{
		// 		m:           1,
		// 		n:           3,
		// 		maxMove:     3,
		// 		startRow:    0,
		// 		startColumn: 1,
		// 	},
		// 	want: 12,
		// },
		{
			name: "",
			args: args{
				m:           36,
				n:           5,
				maxMove:     50,
				startRow:    15,
				startColumn: 3,
			},
			want: 390153306,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPaths(tt.args.m, tt.args.n, tt.args.maxMove, tt.args.startRow, tt.args.startColumn); got != tt.want {
				t.Errorf("findPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
