package main

import (
	"reflect"
	"testing"
)

func Test_queryResults(t *testing.T) {
	type args struct {
		limit   int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				limit: 4,
				queries: [][]int{
					{1, 4},
					{2, 5},
					{1, 3},
					{3, 4},
				},
			},
			want: []int{1, 2, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := queryResults(tt.args.limit, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queryResults() = %v, want %v", got, tt.want)
			}
		})
	}
}
