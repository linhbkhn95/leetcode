package rental

func Solution(parent []string, ch []string) int {

	lenParent, lenCh := len(parent), len(ch)
	if len(parent) == 0 {
		return 0
	}
	if lenCh > lenParent {
		return 0
	}
	result := 0
	for i := 0; i < lenParent; i++ {
		if parent[i] == ch[0] && (i+lenCh) <= lenParent && parent[i+lenCh-1] == ch[lenCh-1] {
			check := true
			for j := 0; j < lenCh; j++ {

				if parent[i+j] != ch[j] {
					check = false
					break
				}
			}
			if check {
				result++
			}
		}
	}
	return result
}

func SecondSolution(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	i := 0
	d := findD(arr)

	for i < len(arr)-1 {
		if arr[i]+d != arr[i+1] {
			return arr[i] + d
		}
		i++
	}
	return -1
}

func findD(arr []int) int {
	l := len(arr)
	d := arr[1] - arr[0]
	if (arr[l-1] - arr[l-2]) < d {
		return arr[l-1] - arr[l-2]
	}
	return d
}
