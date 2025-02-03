package gridgame

import "testing"

func Test_gridGame(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "1",
			args: args{
				grid: [][]int{
					{3, 3, 1}, {8, 5, 2},
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gridGame(tt.args.grid); got != tt.want {
				t.Errorf("gridGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
