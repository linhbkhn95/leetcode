package programa

import "testing"

func Test_numPairsDivisibleBy60(t *testing.T) {
	type args struct {
		time []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Test 1",
			args: args{
				time: []int{30, 20, 150, 100, 40, 20, 20, 20, 20, 40, 30, 30, 30, 30, 30, 30, 30, 20, 150, 100, 40, 20, 20, 20, 20, 40, 30, 30, 30, 30, 30, 30, 30, 20, 150, 100, 40, 20, 20, 20, 20, 40, 30, 30, 30, 30, 30, 30, 30, 20, 150, 100, 40, 20, 20, 20, 20, 40, 30, 30, 30, 30, 30, 30, 30, 20, 150, 100, 40, 20, 20, 20, 20, 40, 30, 30, 30, 30, 30, 30},
			},
			want: 1155,
		},
		{
			name: "Test 2",
			args: args{
				time: []int{439, 407, 197, 191, 291, 486, 30, 307, 11},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numPairsDivisibleBy60(tt.args.time); got != tt.want {
				t.Errorf("numPairsDivisibleBy60() = %v, want %v", got, tt.want)
			}
		})
	}
}
