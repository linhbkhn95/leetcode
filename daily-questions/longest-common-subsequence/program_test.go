package longestcommonsubsequence

import "testing"

func Test_longestCommonSubsequence(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				text1: "abcde",
				text2: "ace",
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				text1: "abc",
				text2: "abc",
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				text1: "abc",
				text2: "def",
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				text1: "bl",
				text2: "yby",
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				text1: "oxcpqrsvwf",
				text2: "shmtulqrypy",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
