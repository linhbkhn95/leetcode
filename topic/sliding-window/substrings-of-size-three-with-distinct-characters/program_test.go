package substringsofsizethreewithdistinctcharacters

import "testing"

func Test_countGoodSubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				s: "xyzzaz",
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				s: "aababcabc",
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				s: "ijedehnd",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodSubstrings(tt.args.s); got != tt.want {
				t.Errorf("countGoodSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
