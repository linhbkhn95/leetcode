package main

import "fmt"

func main() {
	fmt.Println('a', 'A')
	fmt.Println('z', 'Z')

}

func longestNiceSubstring(s string) string {
	return dp(s)
}

func dp(s string) string {
	if len(s) < 2 {
		return ""
	}
	footprint := make(map[rune]bool)
	for _, c := range s {
		footprint[c] = true
	}
	result := ""
	for i, c := range s {
		if c > 90 && footprint[c-32] {
			continue
		}
		if c <= 90 && footprint[c+32] {
			continue
		}
		str1, str2 := dp(s[:i]), dp(s[i+1:])
		if len(str1) >= len(str2) {
			result = str1
		} else {
			result = str2
		}
	}
	if len(result) > 0 {
		return result
	}
	return s
}

func longestNiceSubstring1(s string) string {

	maxLen := 1
	result := ""
	i, j := 0, 0
	for i = 0; i < len(s)-1; i++ {
		for j = i + 1; j < len(s); j++ {
			if i == 3 && j == len(s)-1 {
				fmt.Println("xx")
			}
			ok := isNiceSubstring(s[i:j])
			if ok {
				if j-i > maxLen {
					maxLen = j - i
					result = s[i:j]
				}
			}
		}
		ok := isNiceSubstring(s[i:j])
		if ok {
			if j-i > maxLen {
				maxLen = j - i
				result = s[i:j]
			}
		}
	}

	return result

}
func isNiceSubstring(s string) bool {
	footprint := make(map[rune]int)
	if len(s) < 2 {
		return false
	}
	for _, c := range s {
		if _, ok := footprint[c]; !ok {
			if c > 90 {
				_, ok := footprint[c-32]
				if ok {
					footprint[c-32]++
				} else {
					footprint[c] = 1

				}
			} else {
				_, ok := footprint[c+32]
				if ok {
					footprint[c+32]++
				} else {
					footprint[c] = 1
				}
			}
		}
	}
	for _, count := range footprint {
		if count == 1 {
			return false
		}
	}
	return true
}
