package program

import "strings"

func wordPattern(pattern string, s string) bool {
	b := strings.Split(s, " ")
	if len(pattern) != len(b) {
		return false
	}
	mapping := make(map[rune]string)
	mapString := make(map[string]rune)

	for i, r := range pattern {
		w, ok := mapping[r]
		if !ok {
			rc, ok := mapString[b[i]]
			if ok && rc != r {
				return false
			}
			mapString[b[i]] = r
			mapping[r] = b[i]
		} else if w != b[i] {
			return false
		}
	}
	return true
}
