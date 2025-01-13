package minimumlengthofstringafteroperations

func minimumLength(s string) int {
	mapS := make(map[rune]int)
	for _, c := range s {
		mapS[c]++
	}
	result := 0
	for _, count := range mapS {
		if count < 3 {
			result += count
		} else if (count-2)%2 == 0 {
			result += 2
		} else {
			result += 1
		}
	}
	return result
}

// 2x + 1 + 2 = a

// a - 2 - 2 - 2 +
