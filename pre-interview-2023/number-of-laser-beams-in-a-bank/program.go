package numberoflaserbeamsinabank

func numberOfBeams(bank []string) int {
	nums := make([]int, len(bank))
	result := 0
	for i := range bank {
		num := 0
		for j := 0; j < len(bank[i]); j++ {
			if bank[i][j] == '1' {
				num++
			}
		}
		nums[i] = num
		for k := i - 1; k >= 0; k-- {
			if nums[k] > 0 {
				result += nums[i] * nums[k]
				break
			}
		}
	}
	return result
}
