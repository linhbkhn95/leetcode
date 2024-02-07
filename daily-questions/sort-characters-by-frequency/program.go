package main

import (
	"sort"
	"strings"
)

type CharFrequency struct {
	c byte
	f int
}

func frequencySort(s string) string {
	frequencies := make(map[byte]int)
	for i := range s {
		frequencies[s[i]]++
	}
	charFrequencies := make([]*CharFrequency, 0, len(frequencies))
	for c, count := range frequencies {
		charFrequencies = append(charFrequencies, &CharFrequency{
			c: c,
			f: count,
		})
	}
	sort.Slice(charFrequencies, func(i, j int) bool {
		return charFrequencies[j].f < charFrequencies[i].f
	})
	result := ""
	for _, f := range charFrequencies {
		result += strings.Repeat(string(f.c), f.f)
	}
	return result
}
