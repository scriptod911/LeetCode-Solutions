func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		} else {
			return s[:1]
		}
	}
	var maxStr string
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if isPalindrome(s[i:j]) {
				if len(s[i:j]) > len(maxStr) {
					maxStr = s[i:j]
				}
			}
		}
	}
	return maxStr
}


func isPalindrome(s string) bool {
	if len(s) == 0 {
		return false
	}
	if len(s) == 1 {
		return true
	}
	if len(s) == 2 {
		if s[0] == s[1] {
			return true
		} else {
			return false
		}
	}
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}