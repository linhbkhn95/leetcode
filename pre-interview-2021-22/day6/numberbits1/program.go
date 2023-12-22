package numberbits1

func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		if num&1 == 1 {
			count++
		}
		// move pointer to left 1 unit
		//example 0101 => 010
		num >>= 1
	}
	return count
}
