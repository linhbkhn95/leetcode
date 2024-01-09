package eh

import (
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		x string
		y string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				x: "ABCD",
				y: "AC",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				x: "ABCD",
				y: "CA",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				x: "ABCDA",
				y: "BA",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution1(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution2(t *testing.T) {
	type args struct {
		connections []string
		name1       string
		name2       string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				connections: []string{"fred:joe", "joe:mary", "mary:fred", "mary:bill"},
				name1:       "fred",
				name2:       "bill",
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				connections: []string{"fred:joe", "joe:mary", "mary:fred", "aaa:bill"},
				name1:       "fred",
				name2:       "bill",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution2(tt.args.connections, tt.args.name1, tt.args.name2); got != tt.want {
				t.Errorf("Solution2() = %v, want %v", got, tt.want)
			}
		})
	}
}
