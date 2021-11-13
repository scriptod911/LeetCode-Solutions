func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			}
			if abs(sum-target) < abs(res-target) {
				res = sum
			}
			if sum < target {
				j++
			} else {
				k--
			}
		}
	}
	return res
}


func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
