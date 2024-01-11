package countnumberofnicesubarrays

func numberOfSubarrays(nums []int, k int) int {
	l := len(nums)
	count := 0
	start, end := 0, 0
	result := 0
	fistOddIndex := -1
	for end < l {
		if nums[end]%2 == 0 {
			if count == k {
				result += fistOddIndex - start + 1
			}
			end++
			continue
		}

		if fistOddIndex == -1 {
			fistOddIndex = end
		}
		count += 1
		switch {
		case count == k:
			{
				result += fistOddIndex - start + 1
			}
		case count > k:
			{
				start = fistOddIndex + 1
				fistOddIndex = start
				for fistOddIndex < end {
					if nums[fistOddIndex]%2 == 1 {
						break
					}
					fistOddIndex++
				}
				result += fistOddIndex - start + 1
				count--
			}
		}
		end++
	}
	return result
}
