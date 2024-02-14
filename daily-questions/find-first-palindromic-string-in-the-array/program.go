package findfirstpalindromicstringinthearray

func firstPalindrome(words []string) string {
	for _, word := range words {
		if isPalindrome(word) {
			return word
		}
	}
	return ""
}
func isPalindrome(word string) bool {
	start, end := 0, len(word)-1
	for start < end {
		if word[start] != word[end] {
			return false
		}
		start++
		end--
	}
	return true
}
