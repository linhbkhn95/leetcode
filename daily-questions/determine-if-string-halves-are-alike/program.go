package main

var vowels = map[byte]bool{
	97: true, 101: true, 105: true, 111: true, 117: true,
}

func halvesAreAlike(s string) bool {
	ls := len(s)
	numOfVowels := 0
	for i := 0; i < ls/2; i++ {
		if vowels[s[i]] || vowels[s[i]+32] {
			numOfVowels++
		}
	}
	for i := ls - 1; i >= ls/2; i-- {
		if vowels[s[i]] || vowels[s[i]+32] {
			numOfVowels--
		}
	}
	return numOfVowels == 0
}
