package largestsubstringbetweentwoequalcharacters

func maxLengthBetweenEqualCharacters(s string) int {
	l := len(s)
	mappingIndex := make(map[byte][]int, 0)
	for i := 0; i < l; i++ {
		mappingIndex[s[i]] = append(mappingIndex[s[i]], i)
	}
	maxLen := -1
	for _, indexes := range mappingIndex {
		if len(indexes) == 1 {
			continue
		}
		maxLen = max(maxLen, indexes[len(indexes)-1]-indexes[0]-1)

	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
