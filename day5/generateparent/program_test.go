package generateparent

import (
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				n: 3,
			},
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := generateParenthesis(tt.args.n)
			mapping := make(map[string]bool)
			for _, s := range tt.want {
				mapping[s] = true
			}
			for _, s := range got {
				_, ok := mapping[s]
				if !ok {
					t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
				}

			}
		})
	}
}
