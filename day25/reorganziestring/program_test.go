package reorganizstring

import (
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// {
		// 	name: "test 1",
		// 	args: args{
		// 		s: "aaabc",
		// 		k: 1,
		// 	},
		// 	want: "abaca",
		// },
		// {
		// 	name: "test 2",
		// 	args: args{
		// 		s: "aaaabc",
		// 		k: 1,
		// 	},
		// 	want: "",
		// },
		// {
		// 	name: "test 3",
		// 	args: args{
		// 		s: "aaaadbc",
		// 		k: 1,
		// 	},
		// 	want: "adabaca",
		// },
		// {
		// 	name: "test 4",
		// 	args: args{
		// 		s: "aabbccdd",
		// 		k: 2,
		// 	},
		// 	want: "abacbcbdcd",
		// },
		{
			name: "test 5",
			args: args{
				s: "baabba",
				k: 2,
			},
			want: "xxx",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution2(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reorganizeString(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test 4",
			args: args{
				S: "aaabc",
			},
			want: "abacac",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorganizeString(tt.args.S); got != tt.want {
				t.Errorf("reorganizeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
