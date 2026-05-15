func isAnagram(input string, compare string) bool {
	inputMap := make(map[rune]int)
	//compareMap := make(map[bytes]int)

	//make sure they are the same length
	if len(input) != len(compare) {
		return false
	}

	//have a map of the value we have
	for _, v := range input {
		inputMap[v]++
	}

	// map of what we want
	for _, v := range compare {
		num, ok := inputMap[v]
		if !ok || num <= 0 {
			return false
		}
		inputMap[v]--
	}

	return true
}
