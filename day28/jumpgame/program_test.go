package jumpgame

import "testing"

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// {
		// 	name: "test 1",
		// 	args: args{
		// 		nums: []int{2, 3, 1, 1, 4},
		// 	},
		// 	want: true,
		// },
		{
			name: "test 2",
			args: args{
				nums: []int{0, 2, 3},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump1(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
