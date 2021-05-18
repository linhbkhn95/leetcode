package integertoman

func IntToRoman(num int) string {
	romanMapping := map[int]string{
		1:   "I",
		5:   "V",
		10:  "X",
		50:  "L",
		100: "C",
		500: "D",
	}
	nums := []int{500, 100, 50, 10, 5, 1}
	result := ""
	for num > 0 {
		if num > 1000 {
			result += "M"
			num -= 1000
		} else if num/900 == 9 {
			result += "CM"
			num -= 900
		} else if num/100 == 4 {
			result += "CD"
			num -= 400
		} else if num/10 == 9 {
			result += "XC"
			num -= 90
		} else if num/10 == 4 {
			result += "XL"
			num -= 40
		} else if num == 9 {
			result += "IX"
			num -= 9
		} else if num == 4 {
			result += "IV"
			num -= 4
		}
		for _, n := range nums {
			if num/n > 1 {
				result += romanMapping[n]
				num -= n
			}
		}
	}
	return result
}
