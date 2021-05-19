package myatoi

import "fmt"

func myAtoi(s string) int {
	var numberStringMapping = map[rune]int{
		48: 0,
		49: 1,
		50: 2,
		51: 3,
		52: 4,
		53: 5,
		54: 6,
		55: 7,
		56: 8,
		57: 9,
	}
	var nums []int
	nagative := true
	flag := 0
	offset := -1
	for i, c := range s {
		if c != 32 {
			if flag > 1 || len(nums) > 0 && (c == 43 || c == 45) {
				return 0
			}

			if c == 43 {
				flag += 1
				nagative = true
				offset = i
				continue
			}
			if c == 45 {
				flag += 1
				nagative = false
				offset = i
				continue
			}
			number, ok := numberStringMapping[c]
			if !ok {
				break
			}
			nums = append(nums, number)
		} else if len(nums) > 0 {
			break
		}
	}
	fmt.Println(offset)
	result := 0
	l := len(nums)
	if l == 0 {
		return 0
	}
	index := 0
	for i, n := range nums {
		if n > 0 {
			index = i
			break
		}
	}
	if index == 0 && l > 0 && nums[0] == 0 {
		return 0
	}
	nums = nums[index:l]
	l = len(nums)
	for i, num := range nums {
		result += num * pow(10, l-1-i)
	}
	if l > 10 {
		if nagative {
			return 2147483647
		}
		return -2147483648
	}
	if !nagative {
		result *= -1
	}
	if result < -2147483648 {
		return -2147483648
	}
	if result > 2147483647 {
		return 2147483647
	}
	return result
}

func pow(p, n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 0; i < n; i++ {
		result = result * p
	}
	return result
}
