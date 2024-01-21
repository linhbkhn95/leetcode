package main

import "testing"

func Test_minSumOfLengths(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "",
		// 	args: args{
		// 		arr:    []int{3, 2, 2, 4, 3},
		// 		target: 3,
		// 	},
		// 	want: 2,
		// },
		// {
		// 	name: "",
		// 	args: args{
		// 		arr:    []int{7, 3, 4, 7},
		// 		target: 7,
		// 	},
		// 	want: 2,
		// },
		// {
		// 	name: "",
		// 	args: args{
		// 		arr:    []int{4, 3, 2, 6, 2, 3, 4},
		// 		target: 6,
		// 	},
		// 	want: -1,
		// },
		{
			name: "",
			args: args{
				arr:    []int{1, 6, 1},
				target: 7,
			},
			want: -1,
		},
		// {
		// 	name: "",
		// 	args: args{
		// 		arr:    []int{2, 1, 3, 3, 2, 3, 1},
		// 		target: 6,
		// 	},
		// 	want: 5,
		// },
		// {
		// 	name: "",
		// 	args: args{
		// 		arr:    []int{1, 1, 1, 2, 2, 2, 4, 4},
		// 		target: 6,
		// 	},
		// 	want: 6,
		// },
		// {
		// 	name: "",
		// 	args: args{
		// 		arr:    []int{2, 2, 4, 4, 4, 4, 4, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		// 		target: 20,
		// 	},
		// 	want: 23,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSumOfLengths(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("minSumOfLengths() = %v, want %v", got, tt.want)
			}
		})
	}
}
