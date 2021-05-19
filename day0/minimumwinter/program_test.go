package minimumwinter

import (
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				arr: []int{1, 3, 6, 4, 1, 2},
			},
			want: 5,
		},
		{
			name: "test1",
			args: args{
				arr: []int{1, 2, 3},
			},
			want: 4,
		},
		{
			name: "test1",
			args: args{
				arr: []int{-1, -3},
			},
			want: 1,
		},
		{
			name: "test1",
			args: args{
				arr: []int{1},
			},
			want: 2,
		},
		{
			name: "test1",
			args: args{
				arr: []int{3},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.arr); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRevesed(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test 120",
			args: args{number: 120},
			want: 21,
		},
		{
			name: "test -123",
			args: args{number: 0},
			want: 0,
		},
		{
			name: "test 1534236469",
			args: args{number: 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Revesed(tt.args.number); got != tt.want {
				t.Errorf("Revesed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			name: "test1",
			args: args{s: "123"},
			want: 123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				A: []int{4, 9, 8, 2, 6},
				K: 3,
			},
			want: 18,
		},
		{
			name: "test1",
			args: args{
				A: []int{5, 6, 3, 4, 2},
				K: 5,
			},
			want: 20,
		},
		{
			name: "test1",
			args: args{
				A: []int{7, 7, 7, 7, 7},
				K: 1,
			},
			want: -1,
		},
		{
			name: "test1",
			args: args{
				A: []int{2, 3, 5, 5},
				K: 3,
			},
			want: 12,
		},
		{
			name: "test1",
			args: args{
				A: []int{2, 3, 5, 5},
				K: 3,
			},
			want: 12,
		},
		{
			name: "test1",
			args: args{
				A: []int{2, 3, 3, 5, 5},
				K: 8,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeperate(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				S: "babaa",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Seperate(tt.args.S); got != tt.want {
				t.Errorf("Seperate() = %v, want %v", got, tt.want)
			}
		})
	}
}
