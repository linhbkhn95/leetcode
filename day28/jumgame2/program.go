package jumpgame2

func jump(nums []int) int {
	l := len(nums)
	maxJump := make([]int, l-1)

	for i := 0; i < l-1; i++ {
		maxJump[i] = i + nums[i]
	}
	minOptionCanJump := extractMinOptionCanJump(maxJump, l-1)
	if minOptionCanJump == -1 {
		return 0
	}
	minStep := 0

	if minOptionCanJump == 0 {
		return 1
	}
	position := minOptionCanJump
	step := 1
	k := position

	for k >= 0 {
		minPosition := chooseMinPosition(k, maxJump)
		if minPosition == -1 {
			break
		}

		step += 1
		if minPosition == 0 {
			if minStep == 0 || step < minStep {
				minStep = step
			}
			break
		}
		if minStep != 0 && minStep <= step {
			break
		}
		k = minPosition
	}

	return minStep

}

func chooseMinPosition(position int, maxJumps []int) int {
	for k := 0; k < position; k++ {
		if maxJumps[k] >= position {
			return k
		}
	}
	return -1
}

func extractMinOptionCanJump(maxJumps []int, end int) int {

	for i, maxJump := range maxJumps {
		if maxJump >= end {
			return i
		}
	}
	return -1
}
