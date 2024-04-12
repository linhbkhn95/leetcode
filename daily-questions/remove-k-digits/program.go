package removekdigits

func removeKdigits(num string, k int) string {
	res := make([]rune, 0)
    
	for _, c := range num {
		for len(res) > 0 && res[len(res) - 1] > c && k > 0 {
			res = res[:len(res) - 1]
			k--
		}

		if len(res) > 0 || c != '0' {
			res = append(res, c)
		}
	}

	for len(res) > 0 && k > 0 {
		res = res[:len(res) - 1]
		k--
	}

	if len(res) == 0 {
		return "0"
	}
	return string(res)
}
// type Number struct {
// 	Value  byte
// 	Offset int
// }

// func removeKdigits(num string, k int) string {
// 	numbers := make([]*Number, 0, len(num)-1)
// 	for i := 0; i < len(num)-1; i++ {
// 		n := Number{
// 			Value:  num[i],
// 			Offset: i,
// 		}
// 		numbers = append(numbers, &n)
// 	}
// 	sort.Slice(numbers, func(i, j int) bool {
// 		return numbers[i].Value > numbers[j].Value
// 	})
// 	result := []byte{}
// 	if len(numbers) < k-1 {
// 		return "0"
// 	} else {
// 		removeDigitsOffset := make(map[int]interface{}, 0)
// 		for i := 0; i < k; i++ {
// 			removeDigitsOffset[numbers[i].Offset] = struct{}{}
// 		}
// 		if numbers[k-1].Value > '1' {
// 			for i := 0; i < len(num); i++ {
// 				if _, ok := removeDigitsOffset[i]; ok {
// 					continue
// 				}
// 				result = append(result, num[i])
// 			}
// 		}
// 	}
// 	return string(result)
// }
