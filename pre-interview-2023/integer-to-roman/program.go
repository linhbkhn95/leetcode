package romantointeger

import "strings"

var valueSymbol = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

func intToRoman(num int) string {
	result := ""
	for num > 0 {
		switch {
		case num >= 1000:
			{
				remain := num / 1000
				result += strings.Repeat("M", remain)
				num = num % 1000
			}
		case num >= 900:
			{
				result += "CM"
				num -= 900
			}
		case num >= 500:
			{
				remain := num / 500
				result += strings.Repeat("D", remain)
				num = num % 500
			}
		case num >= 400:
			{
				result += "CD"
				num -= 400
			}
		case num >= 100:
			{
				remain := num / 100
				result += strings.Repeat("C", remain)
				num = num % 100
			}
		case num >= 90:
			{
				result += "XC"
				num -= 90
			}
		case num > 50:
			{
				remain := num / 50
				result += strings.Repeat("L", remain)
				num = num % 50
			}
		case num >= 40:
			{
				result += "XL"
				num -= 40
			}
		case num >= 10:
			{
				remain := num / 10
				result += strings.Repeat("X", remain)
				num = num % 10
			}
		case num == 9:
			{
				result += "IX"
				num = 0
			}

		case num >= 5:
			{
				remain := num / 5
				result += strings.Repeat("V", remain)
				num = num % 5
			}
		case num == 4:
			{
				result += "IV"
				num = 0
			}
		default:
			{
				result += strings.Repeat("I", num)
				num = 0
			}
		}
	}
	return result
}

func process(s string, num, t int) (string, int) {
	remain := num / t
	s += strings.Repeat(valueSymbol[t], remain)
	num = num % 1000
	return s, num
}
