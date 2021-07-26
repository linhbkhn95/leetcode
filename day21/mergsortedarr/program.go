package mergsortedarr

import "fmt"

func MergSortedArray(firstArr, secondArr []int) []int {

	l1, l2 := len(firstArr), len(secondArr)

	result := make([]int, l1+l2)
	fmt.Println("l1, l2", l1, l2)

	i, j, k := 0, 0, 0
	for i < l1 && j < l2 {
		fmt.Println("i,j,k", i, j, k)
		if firstArr[i] < secondArr[j] {
			result[k] = firstArr[i]
			k++
			i++
		} else {
			result[k] = secondArr[j]
			k++
			j++
		}
	}
	fmt.Println("i, j", i, j)

	if i < l1 {
		for i < l1 {
			result[k] = firstArr[i]
			i++
			k++
		}
	} else {
		if j < l2 {
			for j < l2 {
				result[k] = secondArr[j]
				j++
				k++
			}
		}
	}
	return result

}
