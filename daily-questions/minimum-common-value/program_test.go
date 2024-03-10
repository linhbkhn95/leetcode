package minimum_common_value

import "testing"

func Test_getCommon(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{2, 4},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				nums1: []int{1, 2, 3, 6},
				nums2: []int{2, 3, 4, 5},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				nums1: []int{2, 4},
				nums2: []int{1, 2},
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCommon(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("getCommon() = %v, want %v", got, tt.want)
			}
		})
	}
}
