package parentheses

import "testing"

func Test_longestValidParentheses(t *testing.T) {
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
			name: "Test 1",
			args: args{
				s: "()",
			},
			want: 2,
		},
		{
			name: "Test 2",
			args: args{
				s: "(()",
			},
			want: 2,
		},
		{
			name: "Test 3",
			args: args{
				s: "(()()))(()()())",
			},
			want: 8,
		},
		{
			name: "Test 4",
			args: args{
				s: "(()()))(()()())(((())((((((()))))))))))))))))))))",
			},
			want: 30,
		},
		{
			name: "Test 5",
			args: args{
				s: "()(()",
			},
			want: 2,
		},
		{
			name: "Test 6",
			args: args{
				s: "(())(",
			},
			want: 4,
		},
		{
			name: "Test 7",
			args: args{
				s: "(()(((()",
			},
			want: 2,
		},
		{
			name: "Test 8",
			args: args{
				s: "(())(())",
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses1(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
