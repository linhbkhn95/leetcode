package program

func findLength(nums1 []int, nums2 []int) int {
	l1, l2 := len(nums1), len(nums2)

	result := 0
	for i := 0; i < l1; i++ {
		if l1-1-i <= result {
			return result
		}
		maxLen := extractMaxSize(nums1, i, nums2, l2)
		if result < maxLen {
			result = maxLen
		}
	}
	return result
}

func extractMaxSize(num1 []int, start int, num2 []int, l2 int) int {
	j := 0
	result := 0
	currentLen := 0
	k := start
	for j < l2 {
		if l2-1-j <= currentLen {
			return currentLen
		}
		k = start
		if num1[start] == num2[j] {
			for num1[k] == num2[j] {
				currentLen++
				k++
				j++
			}
			if currentLen > result {
				result = currentLen
			}
		}
		j++
	}
	return result
}
