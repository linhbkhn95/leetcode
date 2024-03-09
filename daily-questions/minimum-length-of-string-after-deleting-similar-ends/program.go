package minimum_length_of_string_after_deleting_similar_ends

func minimumLength(s string) int {
	start, end := 0, len(s)-1
	for start < end {
		if s[start] != s[end] {
			return end - start + 1
		}
		for s[start] == s[start+1] {
			start++
		}

		for s[end] == s[end-1] {
			end--
		}
		if s[start] == s[end] {
			end--
			start++
		}
	}
	return 0
}
