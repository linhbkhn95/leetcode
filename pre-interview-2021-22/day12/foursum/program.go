package foursum

type key struct {
	index1 int
	index2 int
}

type markKey struct {
	sumSquare int
	sum       int
}

func fourSum(nums []int, target int) [][]int {
	l := len(nums)
	if l < 4 {
		return [][]int{}
	}
	mapping := make(map[int][]key)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			mapping[nums[i]+nums[j]] = append(mapping[nums[i]+nums[j]], key{index1: i, index2: j})
		}
	}

	var result [][]int

	mark := make(map[markKey]bool)
	for sum, arr1 := range mapping {
		arr, ok := mapping[target-sum]
		if ok {
			for _, num1 := range arr1 {
				for _, num2 := range arr {
					if num1.index1 == num2.index1 || num1.index1 == num2.index2 || num1.index2 == num2.index1 || num1.index2 == num2.index2 {
						continue
					}
					h := getKey(nums[num1.index1], nums[num1.index2], nums[num2.index1], nums[num2.index2])
					if _, ok := mark[h]; !ok {
						mark[h] = true
						result = append(result, []int{nums[num1.index1], nums[num1.index2], nums[num2.index1], nums[num2.index2]})
					}

				}
			}
		}
	}
	return result
}

func getKey(num1, num2, num3, num4 int) markKey {
	return markKey{sumSquare: num1*num1 + num2*num2 + num3*num3 + num4*num4, sum: num1 + num2 + num3 + num4}
}
