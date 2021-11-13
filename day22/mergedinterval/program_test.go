package mergedinterval

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			},
			want: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name: "test2",
			args: args{
				intervals: [][]int{{1, 3}},
			},
			want: [][]int{{1, 3}},
		},
		{
			name: "test3",
			args: args{
				intervals: [][]int{{1, 3}, {2, 3}, {4, 6}, {2, 5}},
			},
			want: [][]int{{1,6}},
		},
		{
			name: "test4",
			args: args{
				intervals: [][]int{{1, 4}, {2, 3}},
			},
			want: [][]int{{1,4}},
		},
		{
			name: "test5",
			args: args{
				intervals: [][]int{{1, 4}, {0, 5}},
			},
			want: [][]int{{0,5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
