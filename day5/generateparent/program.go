package generateparent

func generateParenthesis(n int) []string {
	res := []string{}
	helper(n, 0, 0, "", &res)
	return res
}
func helper(n, left, right int, temp string, res *[]string) {
	if n == right {
		*res = append(*res, temp)
		return
	}

	if left != n {
		helper(n, left+1, right, temp+"(", res)
	}

	if left > right {
		helper(n, left, right+1, temp+")", res)
	}
}
