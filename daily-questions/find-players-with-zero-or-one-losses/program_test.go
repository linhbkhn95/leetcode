package main

import (
	"reflect"
	"testing"
)

func Test_findWinners(t *testing.T) {
	type args struct {
		matches [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				matches: [][]int{
					{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9},
				},
			},
			want: [][]int{
				{1, 2, 10}, {4, 5, 7, 8},
			},
		},
		{
			name: "",
			args: args{
				matches: [][]int{
					{2, 3}, {1, 3}, {5, 4}, {6, 4},
				},
			},
			want: [][]int{
				{1, 2, 5, 6}, nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWinners(tt.args.matches); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWinners() = %v, want %v", got, tt.want)
			}
		})
	}
}
