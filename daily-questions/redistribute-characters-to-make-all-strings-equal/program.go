package redistributecharacterstomakeallstringsequal

func makeEqual(words []string) bool {
	countChar := make(map[byte]int, 0)
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			countChar[word[i]]++
		}
	}
	for _, count := range countChar {
		if count%len(words) != 0 {
			return false
		}
	}
	return true
}
