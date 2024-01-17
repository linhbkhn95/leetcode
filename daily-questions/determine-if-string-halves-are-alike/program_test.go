package main

import "testing"

func Test_halvesAreAlike(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				s: "book",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				s: "textbook",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				s: "TEextbook",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := halvesAreAlike(tt.args.s); got != tt.want {
				t.Errorf("halvesAreAlike() = %v, want %v", got, tt.want)
			}
		})
	}
}
