package longestincreasingsubsequence

import "testing"

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{
					2, 5, 3, 4,
				},
			},
			want: 3,
		},
		// {
		// 	name: "",
		// 	args: args{
		// 		nums: []int{
		// 			10, 9, 2, 5, 3, 7, 101, 18,
		// 		},
		// 	},
		// 	want: 4,
		// },
		// {
		// 	name: "",
		// 	args: args{
		// 		nums: []int{
		// 			0, 1, 0, 3, 2, 3,
		// 		},
		// 	},
		// 	want: 4,
		// },
		// {
		// 	name: "xxx",
		// 	args: args{
		// 		nums: []int{
		// 			7, 7, 7, 7, 7, 7, 7,
		// 		},
		// 	},
		// 	want: 1,
		// },
		// {
		// 	name: "",
		// 	args: args{
		// 		nums: []int{
		// 			0, 1, 0, 3, 2, 3, 3, 5, 4, 3, 5, 3, 5, 4, 5, 3, 5, 67, 7, 7, 5, 4, 4, 42, 4, 6, 7, 8, 9, 0,
		// 		},
		// 	},
		// 	want: 10,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
