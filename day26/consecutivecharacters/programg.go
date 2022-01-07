package consecutivechars

func maxPower(s string) int {
	result := 0
	l := len(s)
	if l < 2 {
		return l
	}
	currentCount := 1

	for i := 1; i < l; i++ {
		if s[i] == s[i-1] {
			currentCount++
		} else {
			if currentCount > result {
				result = currentCount
			}
			currentCount = 1
		}
	}
	if currentCount > result {
		return currentCount
	}
	return result
}
