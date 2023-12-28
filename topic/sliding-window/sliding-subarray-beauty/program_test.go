package slidingsubarraybeauty

import (
	"reflect"
	"testing"
)

func Test_getSubarrayBeauty(t *testing.T) {
	type args struct {
		nums []int
		k    int
		x    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, -1, -3, -2, 3},
				k:    3,
				x:    2,
			},
			want: []int{-1, -2, -2},
		},
		{
			name: "",
			args: args{
				nums: []int{-1, -2, -3, -4, -5},
				k:    2,
				x:    2,
			},
			want: []int{-1, -2, -3, -4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSubarrayBeauty(tt.args.nums, tt.args.k, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSubarrayBeauty() = %v, want %v", got, tt.want)
			}
		})
	}
}
