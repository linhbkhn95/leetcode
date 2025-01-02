package countvowelstringsinranges

/*
Recap:
- words a1 ... an
- queries array [l->r] -> [[1,2],[3,5] ...
- find total strings that start and end with vowel
	[1, 2 ...]
-
Solution:
- use array allocation technique a[n]
- count vowel per index
- count of vowel in range [i, j] = a[j] - a[i]  + 1 (if word[i] has start and end of word that are vowel)

*/

func vowelStrings(words []string, queries [][]int) []int {
	countVowel := make([]int, len(words))
	count := 0
	for i, word := range words {
		if isVowel(word[0]) && isVowel(word[len(word)-1]) {
			count += 1
		}
		countVowel[i] = count
	}
	result := make([]int, len(queries))
	for i, query := range queries {
		startWord := words[query[0]]
		result[i] = countVowel[query[1]] - countVowel[query[0]]
		if isVowel(startWord[0]) && isVowel(startWord[len(startWord)-1]) {
			result[i] += 1
		}
	}
	return result
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'o' || c == 'u' || c == 'i'
}
