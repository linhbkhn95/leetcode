package courseschedule

func canFinish(numCourses int, prerequisites [][]int) bool {

	nodes := make(map[int][]int)
	for _, pre := range prerequisites {
		nodes[pre[0]] = append(nodes[pre[0]], pre[1])
	}
	markedVisted := make(map[int]bool)
	footprint := make(map[int]bool)
	for i := 0; i < numCourses; i++ {
		if isCycle(nodes, i, markedVisted, footprint) {
			return false
		}
	}
	return true
}

func isCycle(nodes map[int][]int, start int, markedVisted, footprint map[int]bool) bool {

	if footprint[start] {
		return true
	}
	if markedVisted[start] {
		return false
	}
	footprint[start] = true
	dependingCourses := nodes[start]
	for _, course := range dependingCourses {
		if isCycle(nodes, course, markedVisted, footprint) {
			return true
		}
	}
	delete(footprint, start)
	markedVisted[start] = true
	return false
}
