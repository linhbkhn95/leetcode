package custom_sort_string

import "sort"

func customSortString(order string, s string) string {
	d := [26]int{}
	for i := range order {
		d[order[i]-'a'] = i
	}

	res := []byte(s)
	sort.Slice(res, func(i, j int) bool {
		return d[res[i]-'a'] < d[res[j]-'a']
	})
	return string(res)
}
