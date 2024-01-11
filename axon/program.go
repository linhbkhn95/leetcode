package axon

import (
	"sort"
)

type WordCount struct {
	Word  string
	Count int
}

func Solution(words []string, k int) []string {
	wordsMapping := make(map[string]int, 0)
	for _, word := range words {
		wordsMapping[word]++
	}
	wordCounts := make([]*WordCount, 0, len(wordsMapping))
	for word, count := range wordsMapping {
		wordCounts = append(wordCounts, &WordCount{
			Word:  word,
			Count: count,
		})
	}
	sort.Slice(wordCounts, func(i, j int) bool {
		return wordCounts[i].Count > wordCounts[j].Count
	})
	result := make([]string, 0, len(wordCounts))
	for i := 0; i < k; i++ {
		if i == len(wordCounts) {
			break
		}
		result = append(result, wordCounts[i].Word)
	}
	return result
}
