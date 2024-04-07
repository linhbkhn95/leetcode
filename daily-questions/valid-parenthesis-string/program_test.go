package validparenthesisstring

import "testing"

func Test_checkValidString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				s: "()",
			},
			want: true,
		},
		{
			args: args{
				s: "(*)",
			},
			want: true,
		},
		{
			args: args{
				s: "(*))",
			},
			want: true,
		},
		{
			args: args{
				s: "(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())",
			},
			want: false,
		},
		{
			args: args{
				s: "(()())*)))())*)*(*()*()))())())((*)((((((())))())*))**)))()*))()))))()()))*)()(*(())((()((()**()()",
			},
			want: false,
		},
		{
			args: args{
				s: "(*()))*(",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkValidString(tt.args.s); got != tt.want {
				t.Errorf("checkValidString() = %v, want %v", got, tt.want)
			}
		})
	}
}
