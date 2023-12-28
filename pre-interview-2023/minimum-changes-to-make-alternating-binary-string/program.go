package minimumchangestomakealternatingbinarystring

func minOperations(s string) int {
	return min(calc(s, '0'), calc(s, '1'))
}

func calc(s string, startC byte) int {
	var next byte
	if startC == '0' {
		next = '1'
	} else {
		next = '0'
	}
	result := 0
	if s[0] != startC {
		result += 1
	}
	for i := 1; i < len(s); i++ {
		if next != s[i] {
			result++
		}
		if next == '0' {
			next = '1'
		} else {
			next = '0'
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
