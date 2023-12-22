package regularexpression

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// {
		// 	name: "test 1",
		// 	args: args{
		// 		s: "aabc",
		// 		p: "aabc",
		// 	},
		// 	want: true,
		// },
		// {
		// 	name: "test 2",
		// 	args: args{
		// 		s: "aabc",
		// 		p: "a*bc",
		// 	},
		// 	want: true,
		// },
		// {
		// 	name: "test 3",
		// 	args: args{
		// 		s: "a",
		// 		p: "a*bc",
		// 	},
		// 	want: false,
		// },
		{
			name: "test 4",
			args: args{
				s: "ab",
				p: ".*",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
