package reversewordsinastring

import "testing"

func Test_trimSpace(t *testing.T) {
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
				s: "  hello world ",
			},
			want: "hello world d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimSpace(tt.args.s); got != tt.want {
				t.Errorf("trimSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
