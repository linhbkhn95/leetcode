package main

import "testing"

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s: "ADOBECODEBANC",
				t: "ABC",
			},
			want: "BANC",
		},
		{
			name: "",
			args: args{
				s: "a",
				t: "a",
			},
			want: "a",
		},
		{
			name: "",
			args: args{
				s: "a",
				t: "aa",
			},
			want: "",
		},
		{
			name: "",
			args: args{
				s: "cabwefgewcwaefgcf",
				t: "cae",
			},
			want: "cwae",
		},
		{
			name: "",
			args: args{
				s: "aaaaaaaaaaaabbbbbcdd",
				t: "abcdd",
			},
			want: "abbbbbcdd",
		},
		{
			name: "",
			args: args{
				s: "aa",
				t: "aa",
			},
			want: "aa",
		},
		{
			name: "",
			args: args{
				s: "cgklivwehljxrdzpfdqsapogwvjtvbzahjnsejwnuhmomlfsrvmrnczjzjevkdvroiluthhpqtffhlzyglrvorgnalk",
				t: "mqfff",
			},
			want: "fsrvmrnczjzjevkdvroiluthhpqtff",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
