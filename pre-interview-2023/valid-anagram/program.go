package validanagram

// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}
// 	count := make(map[byte]int)
// 	for i := 0; i < len(s); i++ {
// 		count[s[i]]++
// 		count[t[i]]--
// 	}

// 	for _, value := range count {
// 		if value > 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

func isAnagram(s string, t string) bool {
	chars := make([]int, 26)

	for _, v := range s {
		i := int(v - 'a')
		chars[i]++
	}

	for _, v := range t {
		i := int(v - 'a')
		chars[i]--
	}

	for _, v := range chars {
		if v != 0 {
			return false
		}
	}

	return true
}
