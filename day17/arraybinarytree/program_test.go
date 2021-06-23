package arraybinarytree

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		arr []int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				arr: []int64{3, 6, 2, 9, -1, 10},
			},
			want: "",
		},
		{
			name: "test2",
			args: args{
				arr: []int64{1, 10, 5, 1, 0, 6},
			},
			want: "",
		},
		{
			name: "test3",
			args: args{
				arr: []int64{1, 4, 10, 5},
			},
			want: "Left",
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
