package lgpd

import "fmt"

func LongestPalindrome(s string) string {
	maxLen := 0
	lenString := len(s)
	var result string
	for i := range s {

		word, lenWord := getPalindromeString(i, lenString, s)
		if maxLen < lenWord {
			result = word
			maxLen = lenWord
		}
	}
	return result
}

func getPalindromeString(i, len int, s string) (string, int) {

	if i < 1 || i > len-1 {
		return string(s[i]), 1
	}
	lenWord := 1
	if s[i] == s[i+1] {

	}
	start, finish := i-1, i+1
	if finish+1 < len-1 && s[i] == s[finish] {
		finish++
	}
	flag := true
	for start > -1 && finish < len {
		if s[start] != s[finish] {
			break

		}
		if flag && s[start] != s[i] {
			flag = false
		}
		start--
		finish++
		lenWord += 2
	}
	// if flag && finish < len && s[finish] == s[i] {
	// 	lenWord++
	// 	for finish < len {
	// 		if s[finish] != s[i] {
	// 			break
	// 		}
	// 		lenWord++
	// 		finish++
	// 	}
	// }
	if lenWord == 1 && i > 0 && s[i-1] == s[i] {

		return s[i-1 : i+1], 2
	}
	if s[start+1:finish] == "cbababcb" {
		fmt.Println(i)
	}
	return s[start+1 : finish], lenWord
}
