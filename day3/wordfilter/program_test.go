package wordfilter

import "testing"

func TestWordFilter_F(t *testing.T) {
	type args struct {
		prefix string
		suffix string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				prefix: "a",
				suffix: "aa",
			},
			want: 5,
		},
	}
	wordFilter := Constructor([]string{"cabaabaaaa","ccbcababac","bacaabccba","bcbbcbacaa","abcaccbcaa","accabaccaa","cabcbbbcca","ababccabcb","caccbbcbab","bccbacbcba"})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordFilter.F(tt.args.prefix, tt.args.suffix); got != tt.want {
				t.Errorf("WordFilter.F() = %v, want %v", got, tt.want)
			}
		})
	}
}
// ["WordFilter","f","f","f","f","f","f","f","f","f","f"]
// [[["cabaabaaaa","ccbcababac","bacaabccba","bcbbcbacaa","abcaccbcaa","accabaccaa","cabcbbbcca","ababccabcb","caccbbcbab","bccbacbcba"]],["bccbacbcba","a"],["ab","abcaccbcaa"],["a","aa"],["cabaaba","abaaaa"],["cacc","accbbcbab"],["ccbcab","bac"],["bac","cba"],["ac","accabaccaa"],["bcbb","aa"],["ccbca","cbcababac"]]