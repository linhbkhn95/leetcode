package lengthoflongestsubarraywithatmostkfrequency

import "testing"

func Test_maxSubarrayLength(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 2, 3, 1, 2, 3, 1, 2},
				k:    2,
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2, 1, 2, 1, 2, 1, 2},
				k:    1,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				nums: []int{5, 5, 5, 5, 5, 5, 5},
				k:    4,
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				nums: []int{2, 2, 3},
				k:    1,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2, 2, 1, 3},
				k:    1,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubarrayLength(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxSubarrayLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
