package maximumscoreaftersplittingastring

func maxScore(s string) int {
	l := len(s)
	// [2]int[0] -> count for 0
	// [2]int[1] -> count for 1
	countZero, countOne := 0, 0
	for i := 0; i < l; i++ {
		if s[i] == '0' {
			countZero++
		} else {
			countOne++
		}
	}

	max := 0
	arrScore := make([]int, l-1)
	if s[0] == '0' {
		arrScore[0] = 1 + countOne
	} else {
		arrScore[0] = 0 + countOne - 1

	}
	for i := 1; i < l-1; i++ {
		arrScore[i] = calcScore(countZero, countOne, i, s, arrScore)
	}
	for _, score := range arrScore {
		if max < score {
			max = score
		}
	}
	return max
}

func calcScore(countZero, countOne, idx int, s string, arrScore []int) int {
	if s[idx] == '0' {
		return arrScore[idx-1] + 1
	}
	return arrScore[idx-1] - 1

}
