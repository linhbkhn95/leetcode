package livecoding

type Employee struct {
	ID         int
	Importance int
	Subs       []int
}

func Solution(emps []*Employee, id int) int {
	empIndexes := make(map[int]*Employee)
	for _, e := range emps {
		empIndexes[e.ID] = e
	}
	footprint := make(map[int]struct{})
	return dp(empIndexes, footprint, id)
}

func dp(indexes map[int]*Employee, footprint map[int]struct{}, empID int) int {
	employ, ok := indexes[empID]
	if !ok {
		return 0
	}
	_, ok = footprint[empID]
	if ok {
		return 0
	}
	footprint[empID] = struct{}{}
	total := employ.Importance
	for _, cod := range employ.Subs {
		total += dp(indexes, footprint, cod)
	}
	return total
}
