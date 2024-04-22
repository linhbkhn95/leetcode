package findifpathexistsingraph

import "testing"

func Test_validPath(t *testing.T) {
	type args struct {
		n           int
		edges       [][]int
		source      int
		destination int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				n: 0,
				edges: [][]int{
					{0, 1}, {1, 2}, {2, 0},
				},
				source:      0,
				destination: 2,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				n: 0,
				edges: [][]int{
					{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3},
				},
				source:      0,
				destination: 5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPath(tt.args.n, tt.args.edges, tt.args.source, tt.args.destination); got != tt.want {
				t.Errorf("validPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
