package countprefixandsuffixpairsi

import "testing"

func Test_countPrefixSuffixPairs(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				words: []string{"a", "abb"},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPrefixSuffixPairs(tt.args.words); got != tt.want {
				t.Errorf("countPrefixSuffixPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
