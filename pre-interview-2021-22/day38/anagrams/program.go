package program

// func findAnagrams(s string, p string) []int {
// 	lS, lP := len(s), len(p)
// 	if lP > lS {
// 		return nil
// 	}
// 	charCount := [26]byte{}
// 	pMapping := make(map[rune]interface{}, lP)
// 	for _, r := range p {
// 		if _, ok := pMapping[r]; !ok {
// 			charCount[r-'a'] = 1
// 			pMapping[r] = nil
// 		} else {
// 			charCount[r-'a'] += 1
// 		}
// 	}
// 	i := 0
// 	result := make([]int, 0)
// 	for i < lS && i+lP <= lS {
// 		if _, ok := pMapping[rune(s[i])]; !ok {
// 			i++
// 			continue
// 		}
// 		j := i
// 		check := [26]byte{}
// 		isBreak := false
// 		for j < lS && j < i+lP {
// 			if _, ok := pMapping[rune(s[j])]; !ok {
// 				i = j
// 				isBreak = true
// 				break
// 			}
// 			check[s[j]-'a'] += 1
// 			j++
// 		}
// 		if isBreak {
// 			i++
// 			continue
// 		}
// 		if check == charCount {
// 			for j < lS {
// 				result = append(result, i)
// 				if s[i] != s[j] {
// 					i++
// 					break
// 				}
// 				j++
// 				i++
// 			}
// 		} else {
// 			i++
// 		}
// 	}
// 	return result
// }

func findAnagrams(s string, p string) []int {
	lS, lP := len(s), len(p)
	if lP > lS {
		return nil
	}
	charCount, dictCount := [26]byte{}, [26]byte{}
	pMapping := make(map[rune]interface{}, lP)
	for _, r := range p {
		if _, ok := pMapping[r]; !ok {
			charCount[r-'a'] = 1
			pMapping[r] = nil
		} else {
			charCount[r-'a'] += 1
		}
	}
	result := make([]int, 0)
	for i, c := range s {
		if i >= lP {
			dictCount[s[i-lP]-'a']--
		}
		dictCount[c-'a']++
		if dictCount == charCount {
			result = append(result, i-lP+1)
		}
	}
	return result
}
