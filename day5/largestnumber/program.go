package largestnumber

func largestNumber(nums []int) string {
	s :=0
	for _,n:= range nums{
		s+=n
	}
	if s==0{
		return "0"
	}
	l := len(nums)
	var arrs [][]int
	mapping := map[int]string{
		0: "0",
		1: "1",
		2: "2",
		3: "3",
		4: "4",
		5: "5",
		6: "6",
		7: "7",
		8: "8",
		9: "9",
	}
	for _, n := range nums {
		arrs = append(arrs, toArray(n))
	}
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if !max(arrs[i], arrs[j]) {
				swap(arrs, i, j)
			}
		}
	}
	result := ""
	for _, arr := range arrs {
		for _, n := range arr {
			s := mapping[n]
			result += s
		}
	}
	return result
}

func toArray(n int) []int {
	arr := []int{}

	for n >= 10 {
		remain := n % 10
		n = n / 10
		arr = append(arr, remain)
	}
	arr = append(arr, n)
	l := len(arr)
	if l == 1 {
		return arr
	}
	i, j := 0, l-1
	for i < j {
		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp
		i++
		j--
	}
	return arr
}

func max(arr1 []int, arr2 []int) bool {
	l1 := len(arr1)
	l2 := len(arr2)
	a12, a21 := append(arr1, arr2...), append(arr2, arr1...)
	for i := 0; i < l1+l2; i++ {
		if a12[i] > a21[i] {
			return true
		}
		if a21[i] > a12[i] {
			return false
		}
	}
	return false
}

func swap(arr [][]int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
