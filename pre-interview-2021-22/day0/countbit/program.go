// Write your code here
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	strs := strings.Split(line, " ")
	nums := make([]int, len(strs))
	for i, str := range strs {
		nums[i], _ = strconv.Atoi(str)
	}
	fmt.Println(min(nums, n))

}

func min(array []int, n int) int {
	if n == 0 {
		return 0
	}
	min := array[0]

	for _, nb := range array {
		if min > nb {
			min = nb
		}
	}
	return min
}

func countBits(num int) []int {
	return doCountBits(num)
}

func doCountBits(num int) []int {
	var result []int
	for i := 0; i <= num; i++ {
		result = append(result, count(i))
	}
	return result
}

func count(num int) int {
	result := 0

	for num > 0 {
		remain := num % 2
		num = num / 2
		if remain == 1 {
			result++
		}
	}
	return result
}

func gcdOfStrings(str1 string, str2 string) string {
	l1 := len(str1)
	l2 := len(str2)
	if l1 == 0 || l2 == 0 {
		return ""
	}

	var gcd, cd string
	i := 0
	min, max := str1, str2
	if l1 > l2 {
		min, max = str2, str1
		l1, l2 = l2, l1
	}
	for i < l1 && len(cd) <= l1 {
		cd = min[0 : i+1]
		if l1%len(cd) == 0 && l2%len(cd) == 0 && isCD(min, cd) && isCD(max, cd) {
			gcd = cd
		}
		i++

	}
	return gcd

}

func isCD(str, cd string) bool {
	k := cd
	lenStr := len(str)
	for len(k) <= lenStr {
		if k == str {
			return true
		}
		k += cd
	}
	return false
}

func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	result := 1
	i := 1
	k := 3
	mark := make(map[int]bool)
	for k < n {
		if _, ok := mark[k]; ok {
			i += 1
			k = 2*i + 1
			continue
		}
		if k*k < n {
			mark[k*k] = true
		}
		if isPrime(k) {
			result++
		}
		i += 1
		k = 2*i + 1
	}
	return result
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
