func letterCombinations(digits string) []string {
		if len(digits) == 0 {
		return []string{}
	}
	var res []string
	nums := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var dfs func(string, string)
	dfs = func(digits, cur string) {
		if len(digits) == 0 {
			res = append(res, cur)
			return
		}
		for _, v := range nums[digits[0]] {
			dfs(digits[1:], cur + string(v))
		}
	}
	dfs(digits, "")
	return res    
}
