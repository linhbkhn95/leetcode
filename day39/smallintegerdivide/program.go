package program

func smallestRepunitDivByK(k int) int {
	remain, lenN := 1, 1

	mark := make(map[int]interface{}, k)
	for remain%k != 0 {
		N := remain*10 + 1
		remain = N % k
		if _, ok := mark[remain]; ok {
			return -1
		}
		mark[remain] = nil
		lenN++
	}
	return lenN
}
