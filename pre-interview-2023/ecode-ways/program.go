package ecodeways

import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {
	if len(s) < 2 {
		if s[0] == '0' {
			return 0
		}
		return 1
	}
	footprint := make(map[string]int)
	c1 := tryFindingDecode(s, len(s), 0, string(s[0]), footprint)
	c2 := tryFindingDecode(s, len(s), 0, string(s[0:2]), footprint)
	return c1 + c2
}

func tryFindingDecode(s string, l, index int, letter string, footprint map[string]int) int {
	if index >= l-1 {
		if len(letter) == 1 {
			if letter[0] == '0' {
				footprint[fmt.Sprintf("%d_1", index)] = 0
				return 0
			}
			footprint[fmt.Sprintf("%d_1", index)] = 1
			return 1
		}
		digit, err := strconv.ParseInt(letter, 10, 64)
		if err != nil || digit > 26 || digit == 0 {
			footprint[fmt.Sprintf("%d_2", index)] = 0
			return 0
		}
		footprint[fmt.Sprintf("%d_2", index)] = 1
		return 1
	}
	if letter[0] == '0' {
		footprint[fmt.Sprintf("%d_1", index)] = 0
		return 0
	}

	digit, err := strconv.ParseInt(letter, 10, 64)
	if err != nil || digit > 26 || digit == 0 {
		return 0
	}
	key := fmt.Sprintf("%d_1", index)
	if len(letter) == 2 {
		key = fmt.Sprintf("%d_2", index)

	}
	c, ok := footprint[key]
	if ok {
		return c
	}

	total := 0
	if index < l-1 {
		c1 := tryFindingDecode(s, l, index+1, string(s[index+1]), footprint)
		total += c1
	}
	if index < l-2 {
		c2 := tryFindingDecode(s, l, index+1, string(s[index:index+2]), footprint)
		total += c2
	}

	footprint[key] = total
	return total
}
