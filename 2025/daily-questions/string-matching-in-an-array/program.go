package stringsmatching

import "strings"

func stringMatching(words []string) []string {
	result := []string{}
	resultMap := make(map[string]interface{})
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if _, ok := resultMap[words[i]]; !ok {
				if strings.Contains(words[j], words[i]) {
					result = append(result, words[i])
					resultMap[words[i]] = struct{}{}
				}
			}
			if _, ok := resultMap[words[j]]; !ok {
				if strings.Contains(words[i], words[j]) {
					result = append(result, words[j])
					resultMap[words[j]] = struct{}{}
				}
			}
		}
	}
	return result
}
