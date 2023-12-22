package romaninteger

func romanToInt(s string) int {
	romanMapping := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	l := len(s)

	result := 0
	pre := 0
	for i := l - 1; i > -1; i-- {
		value, ok := romanMapping[string(s[i])]
		if !ok {
			return 0
		}

		if pre > value {
			result -= value
		} else {
			result += value
		}
		pre = value
	}
	return result
}
