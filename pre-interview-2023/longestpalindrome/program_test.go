package longestpalindrome

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		// {
		// 	name: "",
		// 	args: args{
		// 		s: "babad",
		// 	},
		// 	want: "bab",
		// },
		// {
		// 	name: "",
		// 	args: args{
		// 		s: "cbbd",
		// 	},
		// 	want: "bb",
		// },
		{
			name: "",
			args: args{
				s: "aaaaabaaaaaacaaaaadccccccccccccfcccccccccccc",
			},
			want: "ccccccccccccfcccccccccccc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
