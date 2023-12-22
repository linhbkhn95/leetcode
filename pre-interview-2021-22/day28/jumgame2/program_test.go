package jumpgame2

import "testing"

func Test_jump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "test 1",
		// 	args: args{
		// 		nums: []int{2, 3, 1, 1, 4},
		// 	},
		// 	want: 2,
		// },
		// {
		// 	name: "test 2",
		// 	args: args{
		// 		nums: []int{2, 3, 0, 1, 4},
		// 	},
		// 	want: 2,
		// },
		{
			name: "test 3",
			args: args{
				nums: []int{2, 3, 1, 1, 4, 3, 4, 5, 5},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.args.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
