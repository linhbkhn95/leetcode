package program

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// {
		// 	name: "test 1",
		// 	args: args{
		// 		A: []int{4, 2, 2, 4, 2},
		// 	},
		// 	want: 5,
		// },
		// {
		// 	name: "test 2",
		// 	args: args{
		// 		A: []int{1, 2, 3, 2},
		// 	},
		// 	want: 3,
		// },
		{
			name: "test 3",
			args: args{
				A: []int{1, 1, 2},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
