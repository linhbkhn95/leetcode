package foursumcount

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	var result int

	sum12mapping := make(map[int]int)
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			count, ok := sum12mapping[num2+num1]
			if ok {
				sum12mapping[num2+num1] = count + 1
			} else {
				sum12mapping[num2+num1] = 1
			}
		}
	}
	sum34mapping := make(map[int]int)

	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			count, ok := sum34mapping[num3+num4]
			if ok {
				sum34mapping[num3+num4] = count + 1
			} else {
				sum34mapping[num3+num4] = 1
			}
		}
	}

	for sum, count := range sum34mapping {
		if c, ok := sum12mapping[0-sum]; ok {

			result += (count * c)
		}
	}

	return result
}
