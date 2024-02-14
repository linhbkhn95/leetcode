package findfirstpalindromicstringinthearray

import "testing"

func Test_firstPalindrome(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				words: []string{"abc", "car", "ada", "racecar", "cool"},
			},
			want: "ada",
		},
		{
			name: "",
			args: args{
				words: []string{"notapalindrome", "racecar"},
			},
			want: "racecar",
		},
		{
			name: "",
			args: args{
				words: []string{"def", "ghi"},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstPalindrome(tt.args.words); got != tt.want {
				t.Errorf("firstPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
