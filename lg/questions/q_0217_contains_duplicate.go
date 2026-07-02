package questions

func containsDuplicate(nums []int) bool {
	numsMap := make(map[int]bool)
	for _, num := range nums {
		if numsMap[num] {
			return true
		}
		numsMap[num] = true
	}
	return false
}
