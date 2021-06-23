package schedulerttask

import "testing"

func Test_leastInterval(t *testing.T) {
	type args struct {
		tasks []byte
		n     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// {
		// 	name: "test1",
		// 	args: args{
		// 		tasks: []byte("AAABBB"),
		// 		n:     2,
		// 	},
		// 	want: 8,
		// },
		{
			name: "test2",
			args: args{
				tasks: []byte("AAAAAABCDEFG"),
				n:     2,
			},
			want: 16,
		},
		// {
		// 	name: "test3",
		// 	args: args{
		// 		tasks: []byte("AAABBB"),
		// 		n:     0,
		// 	},
		// 	want: 6,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leastInterval(tt.args.tasks, tt.args.n); got != tt.want {
				t.Errorf("leastInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
