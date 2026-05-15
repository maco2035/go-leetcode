func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, n := range nums {
		v := m[n]
		m[n] = v + 1
		if v >= 1 {
			return true
		}
	}
	return false
}