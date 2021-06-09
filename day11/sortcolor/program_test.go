package sortcolor

import (
	"reflect"
	"testing"
)

func Test_sortColors(t *testing.T) {
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
				nums: []int{2,1,2,0,1,1,1,1},
			},
			want: []int{0,1,1,1,1,1,2,2},
		},
		// {
		// 	name: "test1",
		// 	args: args{
		// 		nums: []int{0,1,2,1,1,1,1,2,2},
		// 	},
		// 	want: []int{0,1,1,1,1,1,2,2,2},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("sortColor() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
