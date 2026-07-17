package questions

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	strMaps := make(map[[26]int][]string, 0)
	for _, str := range strs {
		var key [26]int
		for _, c := range str {
			key[c-'a'] += 1
		}
		if _, ok := strMaps[key]; ok {
			strMaps[key] = append(strMaps[key], str)
		} else {
			strMaps[key] = make([]string, 0)
			strMaps[key] = append(strMaps[key], str)
		}
	}
	for _, strs := range strMaps {
		result = append(result, strs)
	}
	return result
}
