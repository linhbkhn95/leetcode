package maximumnestingdepthoftheparentheses

func maxDepth(s string) int {
	maxVPS := 0

	currentVPS := 0
	for i := range s {
		if s[i] == '(' {
			currentVPS++
		}
		if s[i] == ')' {
			maxVPS = max(maxVPS, currentVPS)
			currentVPS--
		}
	}
	return maxVPS
}
