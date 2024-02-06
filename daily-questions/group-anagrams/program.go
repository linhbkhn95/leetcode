package main

func groupAnagrams(strs []string) [][]string {
	mapResult := make(map[[26]int][]string)
	for _, s := range strs {
		count := [26]int{}
		for i := range s {
			count[s[i]-'a']++
		}
		mapResult[count] = append(mapResult[count], s)
	}
	result := make([][]string, 0)
	for _, group := range mapResult {
		result = append(result, group)
	}
	return result
}
