package lgpd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LongestPalindrome(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		data     string
		expected string
	}{
		{
			name:     "Should return it when having a char",
			data:     "a",
			expected: "a",
		},
		{
			name:     "input is dababae then returning ababa",
			data:     "dababae",
			expected: "ababa",
		},
		{
			name:     "input is aaaa then returning aaaa",
			data:     "aaaa",
			expected: "aaaa",
		},
		{
			name:     "input is babad then returning bab",
			data:     "babad",
			expected: "bab",
		},
		{
			name:     "input is abbcccbbbcaaccbababcbcabca then returning bbcccbb",
			data:     "abbcccbbbcaaccbababcbcabca",
			expected: "bbcccbb",
		},
		{
			name:     "input is tattarrattat then returning tattarrattat",
			data:     "tattarrattat",
			expected: "tattarrattat",
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actual := LongestPalindrome(test.data)
			assert.Equal(t, test.expected, actual)
		})
	}
}
