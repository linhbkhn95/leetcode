package programa

func numPairsDivisibleBy60(time []int) int {
	for i := range time {
		if time[i] >= 60 {
			time[i] = time[i] % 60
		}
	}
	mapping := make(map[int]int)
	for _, remain := range time {
		_, ok := mapping[remain]
		if !ok {
			mapping[remain] = 1
		} else {
			mapping[remain]++
		}
	}
	result := 0
	for remain, c := range mapping {
		if remain != 0 && remain != 30 {
			if count, ok := mapping[60-remain]; ok {
				result += count * c
				delete(mapping, remain)
				delete(mapping, 60-remain)
			}
		} else if c > 1 {
			c := 60 - remain
			if remain == 0 {
				c = 0
			}
			if count, ok := mapping[c]; ok {
				if count == 2 {
					result += 1
				} else {
					result += process(2, count)
				}
				delete(mapping, c)
			}
		}
	}
	return result
}

func process(c, k int) int {
	if k > c {
		c, k = k, c
	}
	result := 1
	for i := c; i > c-k; i-- {
		result *= i
	}
	for j := 2; j <= k; j++ {
		result /= j
	}
	return result
}
