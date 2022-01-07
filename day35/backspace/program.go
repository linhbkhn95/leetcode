package program

func backspaceCompare(s string, t string) bool {
	return doBackspace(s) == doBackspace(t)
}

func doBackspace(s string) string {
	result := ""
	for _, c := range s {
		if c == '#' {
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
		} else {
			result += string(c)
		}
	}
	return result
}

func longestEvenWord(sentence string) string {
	// Write your code here
	maxLen := 0
	current := 0
	result, currentString := "", ""
	l, i := len(sentence), 0
	for i < l {
		c := sentence[i]
		if c == ' ' {
			k := i
			for sentence[k] == ' ' {
				k++
			}
			i = k
			if current%2 == 0 {
				if current > maxLen {
					result = currentString
					maxLen = current
				}
			}
			current = 0
			currentString = ""
		} else {
			currentString += string(c)
			current++
		}
		i++
	}
	if current%2 == 0 {
		if current > maxLen {
			return currentString
		}
	}
	return result

}

func max(first, second int) int {
	if first > second {
		return first
	}
	return second
}
