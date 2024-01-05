package ransomnote

func canConstruct(ransomNote string, magazine string) bool {
	magazineMapping := make(map[byte]int)
	for i := 0; i < len(magazine); i++ {
		magazineMapping[magazine[i]]++
	}
	for i := 0; i < len(ransomNote); i++ {
		magazineMapping[ransomNote[i]]--
	}

	for _, count := range magazineMapping {
		if count < 0 {
			return false
		}
	}
	return true
}
