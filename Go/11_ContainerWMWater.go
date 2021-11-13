func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxArea := 0
	for l < r {
		area := (r - l) * min(height[l], height[r])
		if area > maxArea {
			maxArea = area
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxArea
}


func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
