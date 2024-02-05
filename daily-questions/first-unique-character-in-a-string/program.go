package main

func firstUniqChar(s string) int {
	count := [26]int{}
	for i := range s {
		count[s[i]-97]++
	}
	for i := range s {
		if count[s[i]-97] == 1 {
			return i
		}
	}
	return -1
}
