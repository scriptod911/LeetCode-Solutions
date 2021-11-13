func twoSum(nums []int, target int) []int { // nums is the array, target is the sum
	var m = make(map[int]int) // m is the map
	for i, v := range nums { // i is the index, v is the value
		if j, ok := m[target - v]; ok { // if the value is in the map
			return []int{j, i} // return the index
		}
		m[v] = i // if the value is not in the map, add it
	}
	return nil // if the value is not in the map, return nil
}