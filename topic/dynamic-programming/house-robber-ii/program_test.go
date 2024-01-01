package houserobberii

import "testing"

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "",
		// 	args: args{
		// 		nums: []int{2, 3, 2},
		// 	},
		// 	want: 3,
		// },
		// {
		// 	name: "",
		// 	args: args{
		// 		nums: []int{1, 2, 3, 1},
		// 	},
		// 	want: 4,
		// },
		{
			name: "",
			args: args{
				nums: []int{1, 22, 32, 22, 33, 2, 1, 2, 8, 88, 9, 99, 999, 2, 3, 4, 2, 3, 2, 3, 2, 3, 2, 3, 2, 4, 2, 3, 2, 34, 2, 3, 43, 3, 3, 2, 34, 2, 4, 3, 4, 2, 43, 3, 4, 4, 4, 3, 3, 4},
			},
			want: 1354,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
