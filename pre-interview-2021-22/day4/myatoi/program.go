package myatoi

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
	for _, c := range s {

		//case <> space
		if c != 32 {
			//case  than + - eg: -+, -0-,  0 -,
			if flag > 1 || len(nums) > 0 && (c == 43 || c == 45) {
				break
			}

			//char == +
			if c == 43 {
				flag += 1
				nagative = true
				continue
			}
			// char == -
			if c == 45 {
				flag += 1
				nagative = false
				continue
			}
			number, ok := numberStringMapping[c]
			if !ok {
				break
			}
			nums = append(nums, number)
			// "34 3"
		} else if len(nums) > 0 {
			break
			// " + - 9"
		} else if flag > 0 {
			return 0
		}
	}
	result := 0
	l := len(nums)

	//[]
	if l == 0 {
		return 0
	}

	//clear 0 at head
	index := 0
	for i, n := range nums {
		if n > 0 {
			index = i
			break
		}
	}
	//handle case [00000]
	if index == 0 && l > 0 && nums[0] == 0 {
		return 0
	}
	//get numbers ><0
	nums = nums[index:l]
	l = len(nums)
	for i, num := range nums {
		result += num * pow(10, l-1-i)
	}
	//handle out of range
	if l > 10 {
		if nagative {
			return 2147483647
		}
		return -2147483648
	}
	// calc
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
