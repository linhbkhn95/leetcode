package main

import "testing"

func Test_longestNiceSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// {
		// 	name: "",
		// 	args: args{
		// 		s: "YazaAay",
		// 	},
		// 	want: "aAa",
		// },
		{
			name: "",
			args: args{
				s: "aAztTkZKv",
			},
			want: "bB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestNiceSubstring(tt.args.s); got != tt.want {
				t.Errorf("longestNiceSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
