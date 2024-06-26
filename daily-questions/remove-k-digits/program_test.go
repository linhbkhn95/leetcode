package removekdigits

import "testing"

func Test_removeKdigits(t *testing.T) {
	type args struct {
		num string
		k   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// {
		// 	name: "",
		// 	args: args{
		// 		num: "1432219",
		// 		k:   3,
		// 	},
		// 	want: "1219",
		// },
		{
			name: "",
			args: args{
				num: "10200",
				k:   1,
			},
			want: "200",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeKdigits(tt.args.num, tt.args.k); got != tt.want {
				t.Errorf("removeKdigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
