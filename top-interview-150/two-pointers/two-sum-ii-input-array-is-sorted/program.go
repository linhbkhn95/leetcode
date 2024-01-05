package twosumiiinputarrayissorted

func twoSum(numbers []int, target int) []int {
	start, end := 0, len(numbers)-1
	var sum int
	for start < end {
		sum = numbers[start] + numbers[end]
		if sum == target {
			return []int{start + 1, end + 1}
		}
		if sum > target {
			end--
		} else {
			start++
		}
	}
	return nil
}
