package reverseprefixofword

import "testing"

func Test_reversePrefix(t *testing.T) {
	type args struct {
		word string
		ch   byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				word: "abcdefd",
				ch:   byte('d'),
			},
			want: "dcbaefd",
		},
		{
			name: "",
			args: args{
				word: "xyxzxe",
				ch:   byte('z'),
			},
			want: "zxyxxe",
		},
		{
			name: "",
			args: args{
				word: "abcd",
				ch:   byte('z'),
			},
			want: "abcd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrefix(tt.args.word, tt.args.ch); got != tt.want {
				t.Errorf("reversePrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
