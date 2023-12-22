package palindrome

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	rev := 0
	number := x
	var nums []int
	for x >= 10 {
		remain := x % 10
		x = x / 10
		nums = append(nums, remain)
	}
	nums = append(nums, x)
	lenNum := len(nums)
	for i, nb := range nums {
		rev += nb * pow(10, lenNum-1-i)
	}

	return rev == number
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
