package countprefixandsuffixpairsi

func countPrefixSuffixPairs(words []string) int {
	result := 0
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if isPrefixAndSuffix(words[i], words[j]) {
				result++
			}
		}
	}
	return result
}

func isPrefixAndSuffix(a, b string) bool {
	if len(a) > len(b) {
		return false
	}
	la, lb := len(a)-1, len(b)-1
	for i := 0; i <= la; i++ {
		if a[i] != b[i] || a[la-i] != b[lb-i] {
			return false
		}
	}
	return true
}
