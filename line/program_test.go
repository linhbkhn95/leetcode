package line

import (
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// {
		// 	name: "Test1",
		// 	args: args{
		// 		S: "4 5 6 - 7 +",
		// 	},
		// 	want: 8,
		// },
		// {
		// 	name: "Test1",
		// 	args: args{
		// 		S: "13 DUP 4 POP 5 DUP + DUP + -",
		// 	},
		// 	want: 7,
		// },
		{
			name: "Test3",
			args: args{
				S: "1048575 DUP +",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.S); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution1(t *testing.T) {
	type args struct {
		S string
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
				S: "00:01:07,400-234-090\n00:05:01,701-080-080\n00:05:00,400-234-090",
			},
			want: 900,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution1(tt.args.S); got != tt.want {
				t.Errorf("Solution1() = %v, want %v", got, tt.want)
			}
		})
	}
}
