package primeindiagonal

import "testing"

func Test_diagonalPrime(t *testing.T) {
	type args struct {
		nums [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				nums: [][]int{
					{1, 2, 3},
					{5, 6, 7},
					{9, 10, 11},
				},
			},
			want: 11,
		},
		{
			name: "",
			args: args{
				nums: [][]int{
					{1, 2, 3},
					{5, 17, 7},
					{9, 10, 11},
				},
			},
			want: 17,
		},
		{
			name: "",
			args: args{
				nums: [][]int{
					{1, 2},
					{5, 17},
				},
			},
			want: 17,
		},
		{
			name: "",
			args: args{
				nums: [][]int{
					{6},
				},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diagonalPrime(tt.args.nums); got != tt.want {
				t.Errorf("diagonalPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
