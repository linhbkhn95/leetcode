package minimumnumberofoperationstomakearrayempty

import "testing"

func Test_minOperations(t *testing.T) {
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
		// 		nums: []int{2, 3, 3, 2, 2, 4, 2, 3, 4},
		// 	},
		// 	want: 4,
		// },
		// {
		// 	name: "",
		// 	args: args{
		// 		nums: []int{2, 1, 2, 2, 3, 3},
		// 	},
		// 	want: -1,
		// },
		{
			name: "",
			args: args{
				nums: []int{19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.nums); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
