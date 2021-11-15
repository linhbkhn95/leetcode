package threesum

import (
	"sort"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			args: args{
				arr: []int{1, 2, 3, -1, 3},
				k:   2,
			},
			want: [][]int{{1, 2, -1}},
		},
		{
			args: args{
				arr: []int{-1, 0, 1, 2, -1, -4},
				k:   0,
			},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			args: args{
				arr: []int{0, 0, 0},
				k:   0,
			},
			want: [][]int{{0, 0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			type result struct {
				s1 int
				s2 int
				s3 int
			}
			got := threeSum(tt.args.arr, tt.args.k)
			gotMapping := make(map[result]interface{}, len(tt.want))
			for _, dt := range got {
				sort.Ints(dt)
				r := result{s1: dt[0], s2: dt[1], s3: dt[2]}
				gotMapping[r] = nil
			}
			for _, dt := range tt.want {
				sort.Ints(dt)
				r := result{s1: dt[0], s2: dt[1], s3: dt[2]}
				if _, ok := gotMapping[r]; !ok {
					t.Errorf("threeSum() = %v, want %v", got, tt.want)
				}
			}
			// if got := threeSum(tt.args.arr, tt.args.k); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("threeSum() = %v, want %v", got, tt.want)
			// }
		})
	}
}
