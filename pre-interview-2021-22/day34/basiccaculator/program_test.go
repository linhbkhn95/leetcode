package program

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "test 1",
		// 	args: args{
		// 		s: "3+2*22",
		// 	},
		// 	want: 47,
		// },
		// {
		// 	name: "test 2",
		// 	args: args{
		// 		s: " 3+ 5/2 ",
		// 	},
		// 	want: 5,
		// },
		// {
		// 	name: "test 3",
		// 	args: args{
		// 		s: " 0 - 22 ",
		// 	},
		// 	want: -22,
		// },
		{
			name: "test 4",
			args: args{
				s: "5 - 2 / 2*2 + 4 / 2",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
