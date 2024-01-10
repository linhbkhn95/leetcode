package repeateddnasequences

func findRepeatedDnaSequences(s string) []string {
	return nativeSolution(s)
}

func nativeSolution(s string) []string {
	if len(s) < 10 {
		return nil
	}
	k := 10
	countDNA := make(map[string]int, 0)
	for i := 0; i+k-1 < len(s); i++ {
		countDNA[s[i:i+k]]++
	}
	result := make([]string, 0)
	for k, c := range countDNA {
		if c > 1 {
			result = append(result, k)
		}
	}
	return result
}
