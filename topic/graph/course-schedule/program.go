package courseschedule

func canFinish(numCourses int, prerequisites [][]int) bool {

	prerequisitesMap := make(map[int][]int)
	for _, pre := range prerequisites {
		prerequisitesMap[pre[1]] = append(prerequisitesMap[pre[1]], pre[0])
	}
	candidates := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if _, ok := prerequisitesMap[i]; !ok {
			candidates = append(candidates, i)
		}
	}
	nodes := make(map[int][]int)
	for _, pre := range prerequisites {
		nodes[pre[0]] = append(nodes[pre[0]], pre[1])
	}
	footprint := make(map[int]bool)
	for _, can := range candidates {
		if graph(nodes, can, numCourses-len(candidates)+1, footprint) {
			return true
		}
	}
	return len(candidates) == 0
}

func graph(nodes map[int][]int, start, remain int, footprint map[int]bool) bool {
	if remain == 0 {
		return true
	}
	if footprint[start] {
		return false
	}
	dependingCourses := nodes[start]

	for _, course := range dependingCourses {
		if footprint[course] {
			continue
		}
		footprint[course] = true
		if !graph(nodes, course, remain-1, footprint) {
			return true
		}
		delete(footprint, course)
	}
	return false
}
