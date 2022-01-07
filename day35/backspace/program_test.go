package program

import (
	"testing"
)

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				s: "ad#",
				t: "ac#",
			},
			want: true,
		},
		{
			name: "test 2",
			args: args{
				s: "ad##",
				t: "a#c#",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestEvenWord(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Test 1",
			args: args{
				sentence: "Time to write great code",
			},
			want: "Time",
		},
		{
			name: "Test 2",
			args: args{
				sentence: "It is a pleasant day today",
			},
			want: "pleasant",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestEvenWord(tt.args.sentence); got != tt.want {
				t.Errorf("longestEvenWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
