package sumofsubarrayminimums

import (
	"math"
)

var tenPowNinePlusSeven = math.Pow(10, 9) + 7

func sumSubarrayMins(arr []int) int {
	leftIndex := -1
	rightIndex := 0
	sum := 0
	minIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[minIndex] > arr[i] {
			sum += (i - leftIndex) * (rightIndex - minIndex + 1)
			leftIndex = minIndex
			minIndex = i
		} else {
			rightIndex = i
		}
	}
	return 1
}

func sumSubarrayMins1(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		minValue := arr[i]
		sum += arr[i]
		for j := i + 1; j < len(arr); j++ {
			if minValue <= arr[j] {
				sum += minValue
				continue
			}
			minValue = min(minValue, arr[j])
			sum += minValue
		}
	}
	return sum % int(tenPowNinePlusSeven)

}
