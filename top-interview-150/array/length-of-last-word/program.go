package main

func lengthOfLastWord(s string) int {
	i := len(s) - 1
	for s[i] == ' ' {
		i--
	}
	result := 0
	for i >= 0 && s[i] != ' ' {
		result++
		i--
	}
	return result
}
