package wordsubsets

func wordSubsets(words1 []string, words2 []string) []string {
	word1Mapping := make([][26]int, len(words1))
	for i := range words1 {
		mapping := [26]int{}
		for _, c := range words1[i] {
			mapping[c-'a']++
		}
		word1Mapping[i] = mapping
	}
	result := make([]string, 0)
	aggregateWord := [26]int{}
	for j := range words2 {
		feq := [26]int{}
		for _, c := range words2[j] {
			feq[c-'a']++
		}
		for i := 0; i < 25; i++ {
			if feq[i] > aggregateWord[i] {
				aggregateWord[i] = feq[i]
			}
		}
	}
	for i := range words1 {
		if isSubset(word1Mapping[i], aggregateWord) {
			result = append(result, words1[i])
		}
	}
	return result
}

func isSubset(a [26]int, b [26]int) bool {
	for i := range a {
		if a[i] < b[i] {
			return false
		}

	}
	return true
}
