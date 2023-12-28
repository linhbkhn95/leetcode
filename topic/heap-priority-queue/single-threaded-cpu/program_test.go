package singlethreadedcpu

import (
	"reflect"
	"testing"
)

func Test_getOrder(t *testing.T) {
	type args struct {
		tasks [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// {
		// 	name: "",
		// 	args: args{
		// 		tasks: [][]int{
		// 			{19, 13}, {16, 9}, {21, 10}, {32, 25}, {37, 4}, {49, 24}, {2, 15}, {38, 41}, {37, 34}, {33, 6}, {45, 4}, {18, 18}, {46, 39}, {12, 24},
		// 		},
		// 	},
		// 	want: []int{6, 1, 2, 9, 4, 10, 0, 11, 5, 13, 3, 8, 12, 7},
		// },
		{
			name: "",
			args: args{
				tasks: [][]int{
					{5, 2}, {7, 2}, {9, 4}, {6, 3}, {5, 10}, {1, 1},
				},
			},
			want: []int{5, 0, 1, 3, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getOrder(tt.args.tasks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
