func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	start := 0
	maxLen := 0
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			start = max(start, m[s[i]]+1)
		}
		m[s[i]] = i
		maxLen = max(maxLen, i-start+1)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}