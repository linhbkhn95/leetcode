package main

import (
	"reflect"
	"testing"
)

func Test_findClosestElements(t *testing.T) {
	type args struct {
		arr []int
		k   int
		x   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   7,
			},
			want: []int{2, 3, 4, 5},
		},
		{
			name: "",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   3,
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   -1,
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "",
			args: args{
				arr: []int{1},
				k:   1,
				x:   1,
			},
			want: []int{1},
		},
		{
			name: "",
			args: args{
				arr: []int{1},
				k:   1,
				x:   1,
			},
			want: []int{1},
		},
		{
			name: "",
			args: args{
				arr: []int{0, 1, 1, 1, 2, 3, 6, 7, 8, 9},
				k:   9,
				x:   4,
			},
			want: []int{0, 1, 1, 1, 2, 3, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findClosestElements(tt.args.arr, tt.args.k, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findClosestElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
