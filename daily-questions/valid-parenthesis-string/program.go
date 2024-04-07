package validparenthesisstring

func checkValidString(s string) bool {
	countStart, count := 0, 0
	leftOffsets := []int{}
	starOffsets := []int{}
	for i := range s {
		switch s[i] {
		case '(':
			leftOffsets = append(leftOffsets, i)
			count++
		case ')':
			if count+countStart <= 0 {
				return false
			}
			count--
			if len(leftOffsets) > 0 {
				leftOffsets = leftOffsets[:len(leftOffsets)-1]
			}
		case '*':
			countStart++
			starOffsets = append(starOffsets, i)
		}
	}

	if count < 0 {
		return countStart+count >= 0
	}
	i, j := len(leftOffsets)-1, len(starOffsets)-1
	for i >= 0 && j >= 0 {
		if leftOffsets[i] > starOffsets[j] {
			return false
		}
		i--
		j--
	}
	if i >= 0 {
		return false
	}
	return true
}

// (
// *(
