package findthetownjudge

func findJudge(n int, trust [][]int) int {

	trusted := make(map[int]int)
	for _, v := range trust {
		if v[0] == v[1] {
			trusted[v[1]] -= n
		} else {
			trusted[v[1]]++
		}
	}
	result := -1
	for k, v := range trusted {
		if v == n-1 {
			if result == -1 {
				result = k
			} else {
				return -1
			}
		}
	}
	return result
}
