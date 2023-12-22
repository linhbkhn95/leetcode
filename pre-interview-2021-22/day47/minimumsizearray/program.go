package minimumsizearray

func minSubArrayLen(target int, nums []int) int {
	l := len(nums)
	sum := 0
	minSize := l + 1
	currentSize := 0
	start, end := 0, 0
	for end < l {
		sum += nums[end]
		for sum >= target {

			currentSize = end - start + 1
			if currentSize == 1 {
				return 1
			}
			if currentSize < minSize {
				minSize = currentSize
			}
			sum -= nums[start]
			start++
		}
		end++

	}
	if minSize == l+1 {
		return 0
	}
	return minSize
}
