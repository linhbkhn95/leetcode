package checkifonestringswapcanmakestringsequal

import "testing"

func Test_areAlmostEqual(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "equal strings",
			args: args{
				s1: "abcd",
				s2: "abcd",
			},
			want: true,
		},
		{
			name: "one swap makes equal",
			args: args{
				s1: "abcf",
				s2: "abfc",
			},
			want: true,
		},
		{
			name: "more than one swap needed",
			args: args{
				s1: "abcd",
				s2: "dcba",
			},
			want: false,
		},
		{
			name: "different lengths",
			args: args{
				s1: "abc",
				s2: "abcd",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areAlmostEqual(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("areAlmostEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
