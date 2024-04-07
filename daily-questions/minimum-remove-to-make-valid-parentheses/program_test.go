package minimumremovetomakevalidparentheses

import "testing"

func Test_minRemoveToMakeValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: "lee(t(c)o)de)",
			},
			want: "lee(t(c)o)de",
		},
		{
			args: args{
				s: "a)b(c)d",
			},
			want: "ab(c)d",
		},
		{
			args: args{
				s: "))((",
			},
			want: "",
		},
		{
			args: args{
				s: "))a((a)",
			},
			want: "a(a)",
		},
		{
			args: args{
				s: "))a((a))((a())",
			},
			want: "a((a))(a())",
		},
		{
			args: args{
				s: "(((((",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minRemoveToMakeValid(tt.args.s); got != tt.want {
				t.Errorf("minRemoveToMakeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
