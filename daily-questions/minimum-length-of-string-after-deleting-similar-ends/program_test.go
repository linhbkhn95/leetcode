package minimum_length_of_string_after_deleting_similar_ends

import "testing"

func Test_minimumLength(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				s: "ca",
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				s: "cabaabac",
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				s: "aabccabba",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumLength(tt.args.s); got != tt.want {
				t.Errorf("minimumLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
