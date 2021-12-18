package jumpgame

// func canJump(nums []int) bool {
// 	return Solution(nums, 0, len(nums))
// }

// func Solution(nums []int, current, l int) bool {
// 	if current > l {
// 		return false
// 	}
// 	jump := nums[current]
// 	if jump == 0 {
// 		if l == current+1 {
// 			return true
// 		}
// 		return false
// 	}
// 	if jump >= l-current {
// 		return true
// 	}

// 	for i := 1; i <= jump; i++ {
// 		if current+i < l && Solution(nums, current+i, l) {
// 			return true
// 		}
// 		continue
// 	}
// 	return false
// }

func canJump1(nums []int) bool {
	l := len(nums)
	maxJumps := make([]int, l-1)
	for i := 0; i < l-1; i++ {
		maxJumps[i] = i + nums[i]
	}
	positionsCanJumpToEnd := extractPositionCanJumpToEnd(maxJumps, l)
	if len(positionsCanJumpToEnd) == 0 {
		return false
	}

	for i := len(positionsCanJumpToEnd) - 1; i >= 0; i-- {
		position := positionsCanJumpToEnd[i]
		tartget := l - 1
		for position >= 0 {

			if maxJumps[position] < tartget {
				position -= 1
			} else {
				if position == 0 {
					return true
				}
				tartget = position
				position -= 1
			}

		}

	}
	return false
}

func canJump(nums []int) bool {
	l := len(nums)
	maxJumps := make([]int, l-1)
	for i := 0; i < l-1; i++ {
		maxJumps[i] = i + nums[i]
	}
	positionsCanJumpToEnd := extractPositionCanJumpToEnd(maxJumps, l)
	if len(positionsCanJumpToEnd) == 0 {
		return false
	}
	mapping := make(map[int]bool, l-1)

	for i := len(positionsCanJumpToEnd) - 1; i >= 0; i-- {
		position := positionsCanJumpToEnd[i]
		ok := findWayGoBackStart(position, l-1, mapping, maxJumps)
		if ok {
			return true
		}
		mapping[position] = false
	}
	return false
}

func findWayGoBackStart(position, targetPosition int, mapping map[int]bool, maxJumps []int) bool {
	if position < 0 {
		return false
	}
	if value, ok := mapping[position]; ok {
		return value
	}

	if maxJumps[position] < targetPosition {
		return findWayGoBackStart(position-1, targetPosition, mapping, maxJumps)
	}
	if position == 0 {
		return true
	}

	return findWayGoBackStart(position-1, position, mapping, maxJumps)

}

func extractPositionCanJumpToEnd(maxJumps []int, l int) []int {
	var result []int
	for i, maxJump := range maxJumps {
		if maxJump >= l-1 {
			result = append(result, i)
		}
	}
	return result
}
