package main

func minSteps(s string, t string) int {
	charCount := make(map[rune]int, 0)
	for _, c := range s {
		charCount[c]++
	}
	for _, c := range t {
		charCount[c]--
	}
	steps := 0

	for _, count := range charCount {
		if count > 0 {
			steps += count
		}
	}
	return steps
}
