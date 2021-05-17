package palindrome

import "fmt"

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	rev := 0
	i := 0
	number := x
	for remain > 10 {
		k := x / 10
		remain := x - k*10
		i++
		rev += remain*pow(10, i) + k
		fmt.Println("rev", rev)
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
