package minimumnumberofcoins

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		coins []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				coins: []int{1, 1, 0, 1, 0, 1},
			},
			want: 1,
		},
		{
			name: "test 2",
			args: args{
				coins: []int{1, 0, 1, 0, 1, 1},
			},
			want: 1,
		},
		{
			name: "test 3",
			args: args{
				coins: []int{1, 1, 0, 1, 1, 1},
			},
			want: 2,
		},
		{
			name: "test 4",
			args: args{
				coins: []int{1, 0, 1},
			},
			want: 0,
		},
		{
			name: "test 5",
			args: args{
				coins: []int{0, 1, 1, 0},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.coins); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
