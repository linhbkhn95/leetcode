package main

func uniqueOccurrences(arr []int) bool {
	occurrences := make(map[int]int)
	for _, num := range arr {
		occurrences[num]++
	}
	count := make(map[int]struct{})
	for _, v := range occurrences {
		if _, ok := count[v]; ok {
			return false
		}
		count[v] = struct{}{}
	}
	return true
}
