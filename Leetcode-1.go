func twoSum(nums []int, target int) ( output []int ){

	// value, index
	m := make(map[int]int)

	for i, v := range nums{
		tmp := target-v
		tmpi, ok := m[tmp]
		fmt.Printf("v: %v, i %v, tmp: %v, ok: %v\n", v, i, tmp, ok)
		if ok && tmpi != i {
			return []int{tmpi, i}
		}
		m[v] = i
	}
	return output
}
