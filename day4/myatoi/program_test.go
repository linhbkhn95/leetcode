package myatoi

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				s: "-42",
			},
			want: -42,
		},
		{
			name: "test2",
			args: args{
				s: "   42",
			},
			want: 42,
		},
		{
			name: "test3",
			args: args{
				s: "words and 987",
			},
			want: 0,
		},
		{
			name: "test4",
			args: args{
				s: "4193 with words",
			},
			want: 4193,
		},
		{
			name: "test5",
			args: args{
				s: "-91283472332",
			},
			want: -2147483648,
		},
		{
			name: "test6",
			args: args{
				s: "2147483648",
			},
			want: 2147483647,
		},
		{
			name: "test7",
			args: args{
				s: "+-12",
			},
			want: 0,
		},
		{
			name: "test8",
			args: args{
				s: "00000-42a1234",
			},
			want: 0,
		},
		{
			name: "test9",
			args: args{
				s: "   +0 123",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
