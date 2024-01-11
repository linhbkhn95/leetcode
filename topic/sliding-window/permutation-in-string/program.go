package permutationinstring

func checkInclusion(s1 string, s2 string) bool {
	l1, l2 := len(s1), len(s2)
	if l1 > l2 {
		return false
	}
	start, end := 0, 0
	countChar := make(map[rune]int)
	for _, c := range s1 {
		countChar[c]++
	}

	currentCountChar := copyCountWord(countChar)
	for end < l2 {
		_, ok := countChar[rune(s2[end])]
		if !ok {
			currentCountChar = copyCountWord(countChar)
			end = end + 1
			start = end
			continue
		}

		currentCountChar[rune(s2[end])]--
		if currentCountChar[rune(s2[end])] == -1 {
			i := start
			for i < end {
				if s2[i] == s2[end] {
					break
				}
				currentCountChar[rune(s2[i])]++
				i++
			}
			currentCountChar[rune(s2[end])]++
			start = i + 1
		}
		if currentCountChar[rune(s2[end])] == 0 {
			delete(currentCountChar, rune(s2[end]))
		}
		if len(currentCountChar) == 0 {
			return true
		}
		end++

	}
	return len(currentCountChar) == 0

}

func copyCountWord(map1 map[rune]int) map[rune]int {
	result := make(map[rune]int, 0)
	for k, v := range map1 {
		result[k] = v
	}
	return result
}
