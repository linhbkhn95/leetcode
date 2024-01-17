package main

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	char2Count := make(map[rune]int, 0)
	char1Count := make(map[rune]int, 0)

	for i, c := range word2 {
		char2Count[c]++
		char1Count[rune(word1[i])]++
	}
	occurrenceChar1 := map[int]int{}
	for _, oc := range char1Count {
		occurrenceChar1[oc]++
	}
	for k := range char2Count {
		if _, ok := char1Count[k]; !ok {
			return false
		}
	}
	if len(char1Count) != len(char2Count) {
		return false
	}

	for _, count := range char2Count {
		_, ok := occurrenceChar1[count]
		if !ok {
			return false
		}
		if occurrenceChar1[count] == 0 {
			return false
		}
		occurrenceChar1[count] -= 1
	}
	return true
}
