package fullfillbracket

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Test 1",
			args: args{
				s: "><",
			},
			want: "<><>",
		},
		{
			name: "Test 2",
			args: args{
				s: ">><<>><<>><<",
			},
			want: "<<>><<>><<>><<>>",
		},
		{
			name: "Test 3",
			args: args{
				s: "><>>>><",
			},
			want: "<<<<<>><<>><<>><<>>>>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.s); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
