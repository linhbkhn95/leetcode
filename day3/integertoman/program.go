package integertoman

func IntToRoman(num int) string {
	romanMapping := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}
	nums := []int{500, 100, 50, 10, 5, 1}
	result := ""
	for num > 0 {
		switch {

		case num >= 1000:
			{
				result += "M"
				num -= 1000
				break
			}
		case num >= 900:
			{
				result += "CM"
				num -= 900
				break
			}
		case 400 <= num && num < 500:
			{
				result += "CD"
				num -= 400
				break
			}
		case num >= 90 && num < 100:
			{
				result += "XC"
				num -= 90
				break
			}
		case num >= 40 && num < 50:
			{
				result += "XL"
				num -= 40
				break
			}
		case num == 9:
			{
				result += "IX"
				num -= 9
				break
			}
		case num == 4:
			{
				result += "IV"
				num -= 4
				break
			}

		default:
			{
				for _, n := range nums {
					if num/n >= 1 {
						result += romanMapping[n]
						num -= n
						break
					}
				}
			}
		}
	}
	return result
}
