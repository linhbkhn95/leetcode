package assigncookies

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	lg, ls := len(g), len(s)
	result, i, j := 0, 0, 0
	for i < ls && j < lg {
		for i < ls && s[i] < g[j] {
			i++
		}
		if i == ls {
			break
		}
		result++
		j++
		i++
	}
	return result

}
