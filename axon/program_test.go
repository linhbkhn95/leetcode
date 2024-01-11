package axon

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		words []string
		k     int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// {
		// 	name: "",
		// 	args: args{
		// 		words: []string{"one", "two", "one", "three", "four", "five", "five", "six", "seven", "nine", "eight", "three"},
		// 		k:     3,
		// 	},
		// 	want: []string{"one", "three", "five"},
		// },
		{
			name: "",
			args: args{
				words: []string{"one", "two", "one", "three"},
				k:     6,
			},
			want: []string{"one", "two", "three"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.words, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
