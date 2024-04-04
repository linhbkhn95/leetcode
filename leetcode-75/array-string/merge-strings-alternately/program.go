package mergestringsalternately

func mergeAlternately(word1 string, word2 string) string {
	i, j, l1, l2 := 0, 0, len(word1), len(word2)
	result := ""
	for i < l1 && j < l2 {
		result += string(word1[i])
		result += string(word2[j])
		i++
		j++
	}
	if i < l1 {
		result += word1[i:]
	}
	if j < l2 {
		result += word2[j:]
	}
	return result
}
