package destinationcity

func destCity(paths [][]string) string {
	bags := make(map[string]int, 0)
	for _, path := range paths {
		_, ok := bags[path[0]]
		if !ok {
			bags[path[0]] = 0
		} else {
			delete(bags, path[0])
		}
		_, ok = bags[path[1]]
		if !ok {
			bags[path[1]] = 1
		} else {
			delete(bags, path[1])
		}
	}
	for k, v := range bags {
		if v == 1 {
			return k
		}
	}
	return ""
}
