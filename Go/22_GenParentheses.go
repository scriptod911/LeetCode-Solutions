func generateParenthesis(n int) []string {
	res := []string{}
	backtrack(&res, "", n, n)
	return res
}


func backtrack(res *[]string, cur string, left, right int) {
	if left == 0 && right == 0 {
		*res = append(*res, cur)
		return
	}
	if left > 0 {
		backtrack(res, cur+"(", left-1, right)
	}
	if right > left {
		backtrack(res, cur+")", left, right-1)
	}
}