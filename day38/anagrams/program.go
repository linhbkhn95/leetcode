package program

func findAnagrams(s string, p string) []int {
	lS, lP := len(s), len(p)
	if lP > lS {
		return nil
	}
	charCount := [26]byte{}
	pMapping := make(map[rune]interface{}, lP)
	for _, r := range p {
		if _, ok := pMapping[r]; !ok {
			charCount[r-'a'] = 1
			pMapping[r] = nil
		} else {
			charCount[r-'a'] += 1
		}
	}
	i := 0
	result := make([]int, 0)
	for i < lS && i+lP <= lS {
		if _, ok := pMapping[rune(s[i])]; !ok {
			i++
			continue
		}
		j := i
		check := [26]byte{}
		for j < lS && j < i+lP {
			if _, ok := pMapping[rune(s[j])]; !ok {
				i = j
				break
			}
			check[s[j]-'a'] += 1
			j++
		}
		if check == charCount {
			result = append(result, i)
		}
		i++

	}
	return result
}
