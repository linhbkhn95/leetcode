package maximumstrongpairxori

import "testing"

func Test_maximumStrongPairXor(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumStrongPairXor(tt.args.nums); got != tt.want {
				t.Errorf("maximumStrongPairXor() = %v, want %v", got, tt.want)
			}
		})
	}
}
