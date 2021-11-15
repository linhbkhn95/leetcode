package evenoldarray

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			args: args{
				arr: []int{1, 2, 3, 4},
			},
			want: []int{2, 1, 4, 3},
		},
		{
			args: args{
				arr: []int{2, 4, 6, 8, 10, 1, 3, 5, 7, 9},
			},
			want: []int{2, 1, 6, 3, 10, 1, 4, 5, 6, 9},
		},
		{
			args: args{
				arr: []int{2, 4, 6, 8, 10, 1, 3, 5, 7, 9},
			},
			want: []int{2, 1, 6, 3, 10, 1, 4, 5, 6, 9},
		},
		{
			args: args{
				arr: []int{1, 2, 4, 6, 7, 9},
			},
			want: []int{2, 1, 4, 7, 6, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
