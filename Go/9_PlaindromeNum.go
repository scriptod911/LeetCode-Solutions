func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	var num int
	for x > num {
		num = num*10 + x%10
		x /= 10
	}
	return x == num || x == num/10
}