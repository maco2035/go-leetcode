func groupAnagrams(input []string) (output [][]string) {
	m := map[[26]int][]string{}

	//make everything into a map
	for _, v := range input{
		k := [26]int{}
		for i := range v {
			k[v[i] - 'a'] +=1
		}
		m[k] = append(m[k], v)
	}

	for _, v := range m {
		output = append(output, v)
	}

	return
}

