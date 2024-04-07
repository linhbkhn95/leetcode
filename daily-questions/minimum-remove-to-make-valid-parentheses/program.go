package minimumremovetomakevalidparentheses

func minRemoveToMakeValid(s string) string {
	result := make([]byte, 0)
	leftParenthesesOffset := make([]int, 0)
	count := 0
	index := 0
	for i := range s {
		switch s[i] {
		case '(':
			leftParenthesesOffset = append(leftParenthesesOffset, index)
			count++
		case ')':
			if count == 0 {
				continue
			}
			leftParenthesesOffset = leftParenthesesOffset[:len(leftParenthesesOffset)-1]
			count--
		}
		result = append(result, s[i])
		index++
	}
	for i := len(leftParenthesesOffset) - 1; i >= 0; i-- {
		if len(result) == 1 {
			return ""
		}
		result = append(result[:leftParenthesesOffset[i]], result[leftParenthesesOffset[i]+1:]...)
	}
	return string(result)
}
