package maximumscoreaftersplittingastring

func maxScore(s string) int {
	l := len(s)
	countOne := 0
	for i := 0; i < l; i++ {
		if s[i] == '1' {
			countOne++
		}
	}
	if l == 2 {
		if countOne == 1 {
			if s[0] == '0' {
				return 2
			}
			return 0
		}
		return 1
	}
	maxScore := 0
	currentZeroCount := 0
	for i := 0; i < l-1; i++ {
		if s[i] == '0' {
			currentZeroCount += 1
		}
		maxScore = max(currentZeroCount+countOne-(i+1-currentZeroCount), maxScore)
	}
	return maxScore
}
