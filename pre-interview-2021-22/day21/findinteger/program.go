package findinteger

import (
	"fmt"
	"math"
)

func findIntegers(n int) int {

	return n + 1 - getNumberOfOne(n)
}

func getNumberOfOne(number int) int {
	count, i := 0, 2
	for {
		if calc(i) < number {
			count++
			i++
		} else {
			break
		}
	}
	return count
}

func calc(numberOfOne int) int {
	s := 0
	for i := 0; i < numberOfOne; i++ {
		s += int(math.Pow(2, float64(i)))
	}
	fmt.Println("s numberOfOne", s, numberOfOne)
	return s
}
