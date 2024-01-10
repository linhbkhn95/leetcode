package substringsofsizethreewithdistinctcharacters

func countGoodSubstrings(s string) int {
	appeared := make(map[byte]int, 0)
	start, end := 0, 0
	result := 0
	k := 3
	for end < len(s) {
		index, ok := appeared[s[end]]
		if ok {
			if index == end-1 {
				appeared = map[byte]int{}
			} else {
				delete(appeared, s[end])
			}
			appeared[s[end]] = end
			start = index + 1

			end++
		} else {
			if end-start == k-1 {
				result++
				delete(appeared, s[start])
				start += 1
			}
			appeared[s[end]] = end
			end++
		}
	}
	return result
}
