package main

import (
	"strconv"
)

func sequentialDigits(low int, high int) []int {
	result := make([]int, 0)

	s := strconv.Itoa(low)

	lLow := len(s)

	fistDigital := getFirstDigital(rune(s[0]))

	startLen := lLow
	for {
		if 10-fistDigital < startLen {
			startLen++
			fistDigital = 1
		}
		sqNumber := generateSequentialNumber(fistDigital, startLen)
		if sqNumber < low {
			fistDigital += 1
			continue
		}
		if sqNumber > high {
			break
		}

		result = append(result, sqNumber)
		fistDigital++
	}

	return result

}

func generateSequentialNumber(firstDigital int, lenN int) int {
	expected := strconv.Itoa(firstDigital)
	for lenN-1 > 0 {
		firstDigital += 1
		expected += strconv.Itoa(firstDigital)
		lenN--
	}

	n, _ := strconv.Atoi(expected)
	return n
}

func getFirstDigital(c rune) int {
	switch c {
	case '1':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	}
	return 0
}
