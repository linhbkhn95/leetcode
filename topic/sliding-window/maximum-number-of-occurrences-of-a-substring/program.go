package maximumnumberofoccurrencesofasubstring

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {

	counter := make(map[string]int)
	lenOfUniqueChar := 0
	countC := map[byte]int{}
	for i := 0; i < minSize; i++ {
		countC[s[i]]++
		if countC[s[i]] == 1 {
			lenOfUniqueChar++
		}
	}
	if lenOfUniqueChar <= maxLetters {
		counter[s[0:minSize]]++
	}
	for i := minSize; i < len(s); i++ {
		countC[s[i]]++
		if countC[s[i]] == 1 {
			lenOfUniqueChar++
		}
		if i-minSize == 0 {
			countC[s[i-minSize]]--
		}
		if countC[s[i-minSize]] == 0 {
			lenOfUniqueChar--
		}
		if lenOfUniqueChar <= maxLetters {
			counter[s[i-minSize:i]]++
		}
	}
	maxCount := 0
	for _, v := range counter {
		maxCount = max(maxCount, v)
	}
	return maxCount
}
