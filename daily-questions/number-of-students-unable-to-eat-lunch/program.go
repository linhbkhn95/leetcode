package numberofstudentsunabletoeatlunch

import "container/list"

func countStudents(students []int, sandwiches []int) int {
	studentsQueue := list.New()
	sandwichesStack := list.New()
	for _, student := range students {
		studentsQueue.PushBack(student)
	}
	for _, s := range sandwiches {
		sandwichesStack.PushBack(s)
	}
	cnt := 0
	for sandwichesStack.Len() > 0 && cnt < studentsQueue.Len() {
		st := studentsQueue.Front()
		sw := sandwichesStack.Front()
		stValue, _ := st.Value.(int)
		swValue, _ := sw.Value.(int)
		if stValue != swValue {
			student := studentsQueue.Front()
			studentsQueue.Remove(student)
			studentsQueue.PushBack(student)
			cnt++
		} else {
			student := studentsQueue.Front()
			studentsQueue.Remove(student)
			sandwich := sandwichesStack.Front()
			sandwichesStack.Remove(sandwich)
			cnt = 0
		}
	}

	return studentsQueue.Len()

}

func countStudents1(students []int, sandwiches []int) int {
	var (
		counters [2]int
		n        int = len(students)
	)

	for i := 0; i < n; i++ {
		counters[students[i]]++
	}

	for i := 0; i < n; i++ {
		if counters[sandwiches[i]] == 0 {
			break
		}
		counters[sandwiches[i]]--
	}

	return counters[0] + counters[1]
}
