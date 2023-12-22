package nextpermution

import (
	"reflect"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: []int{1, 3, 2},
		},
		{
			name: "test2",
			args: args{
				nums: []int{1, 1, 5},
			},
			want: []int{1, 5, 1},
		},
		{
			name: "test3",
			args: args{
				nums: []int{3, 2, 1},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "test4",
			args: args{
				nums: []int{1, 9, 1, 3, 2},
			},
			want: []int{1, 9, 2, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("Test_nextPermutation() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
