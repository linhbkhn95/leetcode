package palindrome

import (
	"fmt"
	"strings"
)

func partition(s string) [][]string {
	result := [][]string{}
	palindromeMapping := make(map[string]bool)
	Solution(s, 0, 1, len(s), "", &result, palindromeMapping)
	return result
}

func Solution(s string, start, end, lastIndex int, item string, result *[][]string, palindromeMapping map[string]bool) {
	if end > lastIndex+1 {
		return
	}
	if end == lastIndex && end != start {
		sampler := s[start:end]
		if extractPalindrome(s, start, end, palindromeMapping) {
			item += sampler
			*result = append(*result, strings.Split(item, " "))
		}
		return
	}
	sampler := s[start:end]
	if extractPalindrome(s, start, end, palindromeMapping) {
		k := item + sampler + " "
		Solution(s, end, end+1, lastIndex, k, result, palindromeMapping)
	}
	Solution(s, start, end+1, lastIndex, item, result, palindromeMapping)

}

func extractPalindrome(s string, start, end int, palindromeMapping map[string]bool) bool {
	key := fmt.Sprintf("%d_%d", start, end)

	value, ok := palindromeMapping[key]
	var oKay bool
	if ok {
		oKay = value
	} else {
		oKay = isPalindrome(s[start:end])
		palindromeMapping[key] = oKay
	}
	return oKay
}

func isPalindrome(s string) bool {
	if s == "" {
		return false
	}
	l := len(s)
	if l == 1 {
		return true
	}
	i, j := 0, l-1
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
