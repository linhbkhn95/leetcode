package program

func isRobotBounded(instructions string) bool {
	// l := len(instructions)
	// k := 0
	// for k :=0; k{
	// 	if instruction == 'G' {

	// 	} else if instruction == 'L' {

	// 	}
	// }

	return true
}

func canMoveInCircle(instructions string) bool {
	coodinator := [2]int{}
	degree := 90
	for _, instruction := range instructions {
		degree = covertDegree(degree)
		if instruction == 'G' {
			move(coodinator, degree)
		} else if instruction == 'L' {
			degree -= 90
		} else {
			degree += 90
		}
	}
	degree = covertDegree(degree)
	if degree == 90 && coodinator[0] != 0 && coodinator[1] != 0 {
		return false
	}
	return true
}

func move(coodinator [2]int, degree int) {
	if degree == 0 {
		coodinator[0] += 1
	} else if degree == 90 || degree == -270 {
		coodinator[1] += 1
	} else if degree == 180 || degree == -180 {
		coodinator[0] -= 1
	} else {
		coodinator[1] -= 1
	}
}

func covertDegree(degree int) int {
	if degree > 360 || degree < -360 {
		return degree % 360
	}
	return degree
}
