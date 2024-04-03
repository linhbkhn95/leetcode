package wordsearch

type key struct {
	i     int
	j     int
	index int
}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != word[0] {
				continue
			}
			if move(board, word, len(board), len(board[0]), len(word), 0, i, j, map[key]bool{}) {
				return true
			}
		}
	}
	return false
}

func move(board [][]byte, word string, m, n, l, index, i, j int, footprint map[key]bool) bool {

	if i < 0 || i > m-1 {
		return false
	}
	if j < 0 || j > n-1 {
		return false
	}
	if board[i][j] == '*' {
		return false
	}
	if index == l-1 {
		if board[i][j] == word[index] {
			return true
		}
		return false
	}
	k := key{
		i:     i,
		j:     j,
		index: index,
	}
	if footprint[k] {
		return false
	}
	if board[i][j] != word[index] {
		footprint[k] = true
		return false
	}
	temp := board[i][j]
	board[i][j] = '*'
	result := move(board, word, m, n, l, index+1, i+1, j, footprint) || move(board, word, m, n, l, index+1, i-1, j, footprint) || move(board, word, m, n, l, index+1, i, j+1, footprint) || move(board, word, m, n, l, index+1, i, j-1, footprint)
	board[i][j] = temp
	return result
}
