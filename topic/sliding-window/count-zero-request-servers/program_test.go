package countzerorequestservers

import (
	"reflect"
	"testing"
)

func Test_countServers(t *testing.T) {
	type args struct {
		n       int
		logs    [][]int
		x       int
		queries []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// {
		// 	name: "",
		// 	args: args{
		// 		n: 3,
		// 		logs: [][]int{
		// 			{1, 3},
		// 			{2, 6},
		// 			{1, 5},
		// 		},
		// 		x:       5,
		// 		queries: []int{10, 11},
		// 	},
		// 	want: []int{1, 2},
		// },
		{
			name: "",
			args: args{
				n: 3,
				logs: [][]int{
					{2, 4}, {2, 1}, {1, 2}, {3, 1},
				},
				x:       2,
				queries: []int{3, 4},
			},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countServers(tt.args.n, tt.args.logs, tt.args.x, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countServers() = %v, want %v", got, tt.want)
			}
		})
	}
}
