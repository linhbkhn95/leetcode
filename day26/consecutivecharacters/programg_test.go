package consecutivechars

import "testing"

func Test_maxPower(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				s: "abbcccddddeeeeedcba",
			},
			want: 5,
		},
		{
			name: "test 2",
			args: args{
				s: "triplepillooooow",
			},
			want: 5,
		},
		{
			name: "test 3",
			args: args{
				s: "hooraaaaaaaaaaay",
			},
			want: 11,
		},
		{
			name: "test 4",
			args: args{
				s: "tourist",
			},
			want: 1,
		},
		{
			name: "test 5",
			args: args{
				s: "cc",
			},
			want: 2,
		},
		{
			name: "test 6",
			args: args{
				s: "aaccc",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPower(tt.args.s); got != tt.want {
				t.Errorf("maxPower() = %v, want %v", got, tt.want)
			}
		})
	}
}
