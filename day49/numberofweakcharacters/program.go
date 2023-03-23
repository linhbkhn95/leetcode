package numberofweakcharacters

func numberOfWeakCharacters(properties [][]int) int {
	if len(properties) < 2 {
		return 0
	}

	mapping := make(map[int]interface{}, 0)
	l := len(properties)
	result := 0
	for i := 0; i < l; i++ {
		for j := i + 1; j < l-1; j++ {
			if _, ok := mapping[i]; ok {
				continue
			}
			if properties[i][0] < properties[j][0] && properties[i][1] < properties[j][1] {
				result++
				mapping[i] = nil
			}
		}
	}
	return result
}
