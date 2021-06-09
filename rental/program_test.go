package rental

import (
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		parent []string
		ch     []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				parent: []string{"a", "b", "c", "a", "c", "b", "c", "d"},
				ch:     []string{"b", "c"},
			},
			want: 2,
		},
		{
			name: "test1",
			args: args{
				parent: []string{"a", "a", "a", "a", "a", "a"},
				ch:     []string{"a", "a"},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.parent, tt.args.ch); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecondSolution(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				arr: []int{
					1, 2, 3, 4, 5, 7, 8, 9, 10,
				},
			},
			want: 6,
		},
		{
			name: "test2",
			args: args{
				arr: []int{
					1, 4, 7, 13, 16,
				},
			},
			want: 10,
		},
		{
			name: "test3",
			args: args{
				arr: []int{
					-10, -4, -1, 2, 5,
				},
			},
			want: -7,
		},
		{
			name: "test4",
			args: args{
				arr: []int{
					-20, -18, -16, -12,
				},
			},
			want: -14,
		},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SecondSolution(tt.args.arr); got != tt.want {
				t.Errorf("SecondSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
