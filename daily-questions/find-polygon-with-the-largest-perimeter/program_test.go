package findpolygonwiththelargestperimeter

import "testing"

func Test_largestPerimeter(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "",
			args: args{
				nums: []int{5, 5, 5},
			},
			want: 15,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 12, 1, 2, 5, 50, 3},
			},
			want: 12,
		},
		{
			name: "",
			args: args{
				nums: []int{5, 5, 50},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestPerimeter(tt.args.nums); got != tt.want {
				t.Errorf("largestPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
