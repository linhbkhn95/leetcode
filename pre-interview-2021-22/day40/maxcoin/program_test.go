package maxcoin

import "testing"

func Test_maxCoins(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// {
		// 	name: "test 1",
		// 	args: args{
		// 		nums: []int{3, 1, 5, 8},
		// 	},
		// 	want: 167,
		// },
		{
			name: "test 2",
			args: args{
				nums: []int{3, 1, 5, 8, 2, 4, 7, 9, 10, 12, 13, 14},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCoins(tt.args.nums); got != tt.want {
				t.Errorf("maxCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
