package main

import "sort"

type Number struct {
	Number int
	Freq   int
}

func findLeastNumOfUniqueInts(arr []int, k int) int {
	mapping := make(map[int]int, 0)
	for _, num := range arr {
		mapping[num]++
	}
	numCounts := make([]*Number, 0, len(mapping))
	for k, v := range mapping {
		numCounts = append(numCounts, &Number{
			Number: k,
			Freq:   v,
		})
	}
	sort.Slice(numCounts, func(i, j int) bool { return numCounts[i].Freq < numCounts[j].Freq })
	i := 0
	for k > 0 && i < len(numCounts) {
		k -= numCounts[i].Freq
		if k >= 0 {
			i++
		}
	}

	return len(numCounts) - i
}
