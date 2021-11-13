func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	if len(p) == 1 {
		return len(s) == 1 && (s[0] == p[0] || p[0] == '.')
	}
	if p[1] != '*' {
		if len(s) == 0 || (s[0] != p[0] && p[0] != '.') {
			return false
		}
		return isMatch(s[1:], p[1:])
	}
	if isMatch(s, p[2:]) {
		return true
	}
	if len(s) == 0 || (s[0] != p[0] && p[0] != '.') {
		return false
	}
	return isMatch(s[1:], p)
}
