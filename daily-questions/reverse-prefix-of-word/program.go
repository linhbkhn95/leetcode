package reverseprefixofword

func reversePrefix(word string, ch byte) string {
	for i := range word {
		if word[i] == ch {
			w := word[i:]
			for j := 0; j <= i; j++ {
				w = string(word[j]) + w
			}
			return w
		}
	}
	return word
}

/*
(99-x)/100-x = 98/100
<=> 9900 -100x = 9800 - 98x
<=> 2x = 100
<=> x = 50

*/
