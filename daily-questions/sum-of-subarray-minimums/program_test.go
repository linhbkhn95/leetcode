package sumofsubarrayminimums

import "testing"

func Test_sumSubarrayMins(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				arr: []int{3, 1, 2, 4},
			},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumSubarrayMins(tt.args.arr); got != tt.want {
				t.Errorf("sumSubarrayMins() = %v, want %v", got, tt.want)
			}
		})
	}
}
