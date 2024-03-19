package longest_consecutive_sequence

func longestConsecutive(nums []int) int {
	numsMapping := make(map[int]interface{})
	for _, n := range nums {
		numsMapping[n] = struct{}{}
	}
	current := 1
	result := 1
	for _, n := range nums {
		_, ok := numsMapping[n]
		if ok {
			left := n - 1
			for {
				_, ok := numsMapping[left]
				if !ok {
					break

				}

				current++
				delete(numsMapping, left)
				left--
			}
			right := n + 1
			for {
				_, ok := numsMapping[right]
				if !ok {
					break
				}
				current++
				delete(numsMapping, right)
				right++
			}
			result = max(result, current)
			current = 1
		}
	}
	return result

}
