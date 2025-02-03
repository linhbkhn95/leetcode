package main

import "testing"

func Test_longestMonotonicSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: 5,
		},
		{
			name: "Test case 2",
			args: args{nums: []int{5, 4, 3, 2, 1}},
			want: 5,
		},
		{
			name: "Test case 3",
			args: args{nums: []int{1, 3, 2, 4, 3}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestMonotonicSubarray(tt.args.nums); got != tt.want {
				t.Errorf("longestMonotonicSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
