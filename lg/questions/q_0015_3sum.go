package questions

import "sort"

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)

	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			continue
		}
		twoSumRes := twoSum(nums, -num, i+1)
		if len(twoSumRes) > 0 {
			for _, twr := range twoSumRes {
				result = append(result, []int{num, twr[0], twr[1]})
			}
		}
	}
	return result
}

func twoSum(nums []int, sum, index int) [][]int {
	result := make([][]int, 0)
	i := index
	j := len(nums) - 1
	for i < j {
		currentSum := nums[i] + nums[j]
		if currentSum > sum {
			j -= 1
		} else if currentSum < sum {
			i += 1
		} else {
			result = append(result, []int{nums[i], nums[j]})
			i += 1
			j -= 1
			for i < j && nums[i] == nums[i-1] {
				i++
			}
			for i < j && nums[j] == nums[j+1] {
				j--
			}
		}
	}
	return result
}
