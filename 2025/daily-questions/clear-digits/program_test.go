package cleardigits

import "testing"

func Test_clearDigits(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s: "cb34",
			},
			want: "",
		},
		{
			name: "2",
			args: args{
				s: "abc",
			},
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clearDigits(tt.args.s); got != tt.want {
				t.Errorf("clearDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
