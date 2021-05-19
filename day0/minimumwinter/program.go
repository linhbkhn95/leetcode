package minimumwinter

import "fmt"

func Solution(arr []int) int {
	if len(arr) == 0 {
		return 1
	}
	return doSolution(arr)
}

func doSolution(arr []int) int {
	numberMapping := make(map[int]bool)
	min := 1
	for _, number := range arr {
		if number > 0 {
			numberMapping[number] = true
			if min > number {
				min = number
			}
		}
	}
	if len(numberMapping) == 0 {
		return 1
	}
	if min > 1 {
		return 1
	}
	result := 1
	for {
		if _, ok := numberMapping[result]; !ok {
			break
		}
		result += 1
	}
	return result
}

func Revesed(number int) int {
	rev := 0

	for number != 0 {
		remain := number % 10
		number = number / 10
		if rev > 2147483647/10 || (rev == 2147483647/10 && remain > 7) {
			return 0
		}
		if rev < -2147483647/10 || (rev == -2147483648/10 && remain < -8) {
			return 0
		}
		rev = rev*10 + remain
	}
	return rev
}

var numberStringMap = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

func myAtoi(s string) int {

	s = getNumberString(s)
	lenString := len(s)
	result := 0
	if lenString > 0 {

		lenArr := len(s)
		for i := (lenArr - 1); i > 0; i-- {
			value, ok := numberStringMap[string(s[i])]
			if !ok {
				return 0
			}

			result = result + value*pow(10, lenArr-1-i)
		}
		if string(string(s[0])) == "-" {
			result = result * -1
			if result < -2147483648 {
				result = -2147483648
			}
		} else {
			if string(string(s[0])) != "+" {
				value, ok := numberStringMap[string(s[0])]
				if !ok {
					return 0
				}
				result = result + value*pow(10, lenArr-1)
			}
			if result >= 2147483648 {
				result = 2147483647
			}
		}
		return result
	}
	return 0
}

func pow(number int, p int) int {
	if p == 0 {
		return 1
	}
	s := 1
	for i := 0; i < p; i++ {
		s = s * number
	}
	return s
}

func getNumberString(s string) string {

	var arrS []string
	s = trimSpace(s)
	for i := 0; i < len(s); i++ {

		if string(s[i]) == "-" || string(s[i]) == "+" {
			arrS = append(arrS, string(s[i]))
			continue
		}
		_, ok := numberStringMap[string(s[i])]

		if !ok {
			if len(arrS) == 0 {
				return ""
			}
			break
		}
		arrS = append(arrS, string(s[i]))

	}
	return passArrayToString(arrS)
}
func trimSpace(s string) string {
	if len(s) == 0 {
		return s
	}
	start := 0
	finish := len(s) - 1
	for string(s[start]) == " " || string(s[finish]) == " " {
		if string(s[start]) == " " {
			start += 1
		}
		if start > finish-1 {
			break
		}
		if string(s[finish]) == " " {
			finish -= 1
		}

	}
	return s[start : finish+1]
}
func passArrayToString(s []string) string {

	result := ""
	for _, c := range s {
		result += c
	}
	return result

}

func Sum(A []int, K int) int {
	if len(A) == 0 {
		return -1
	}
	result := doSum(A, len(A), K, 0, 0, 0)
	if result == 0 {
		return -1
	}
	return result
}

func doSum(A []int, len, K, offset, currentSum, largetSum int) int {
	fmt.Println(K, offset, currentSum, largetSum)
	if offset == len || K == 0 {
		if K == 0 {
			if currentSum%2 == 0 && currentSum > largetSum {
				return currentSum
			}
		}
		return 0
	}
	case1 := doSum(A, len, K-1, offset+1, currentSum+A[offset], largetSum)
	case2 := doSum(A, len, K, offset+1, currentSum, largetSum)
	if case1 > case2 {
		return case1
	}
	return case2
}

func Seperate(S string) int {
	numberA := 0
	for i := 0; i < len(S); i++ {
		if string(S[i]) == "a" {
			numberA++
		}
	}
	result := 0
	for i := 0; i < numberA; i++ {
		doSeperate(S, len(S), 0, 0, i, result)
	}
	return result
}

func doSeperate(s string, len, offset, currentA, sameNumberA, allOfWay int) {
	if offset == len-1 {
		if currentA == sameNumberA {
			allOfWay += 1
		}
		return
	}
	if currentA < sameNumberA {
		if string(s[offset]) == "a" {
			doSeperate(s, len, offset+1, currentA+1, sameNumberA, allOfWay)
		} else {
			doSeperate(s, len, offset+1, currentA, sameNumberA, allOfWay)
		}
	} else {
		doSeperate(s, len, offset+1, 0, sameNumberA, allOfWay)
	}
}
