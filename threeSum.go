package leetcode

import "fmt"
import "sort"

func threeSum(nums []int) [][]int {
	var finall [][]int
	if len(nums) < 3 {
		return finall
	}
	existKey := map[string]int{}

	numsMap := map[int]int{}
	for _, value := range nums {
		time, ok := numsMap[value]
		if ok {
			numsMap[value] = time + 1
		} else {
			numsMap[value] = 1
		}
	}

	for i := 0; i < len(nums); i++ {
		tmpNumsMap := map[int]int{}
		for key, value := range numsMap {
			if value > 0 {
				tmpNumsMap[key] = value
			}
		}

		var first, second, third int
		if tmpNumsMap[nums[i]] > 0 {
			tmpNumsMap[nums[i]]--
			first = nums[i]
		} else {
			break
		}

	LOOP:
		for j := i + 1; j < len(nums); j++ {
			if tmpNumsMap[nums[j]] > 0 {
				tmpNumsMap[nums[j]]--
				second = nums[j]
			} else {
				break LOOP
			}
			third = 0 - first - second
			time, ok := tmpNumsMap[third]
			if ok && time > 0 {
				result := []int{first, second, third}
				sort.Ints(result)
				key := fmt.Sprintf("%v", result)
				_, ok := existKey[key]
				if !ok {
					existKey[key] = 1
					finall = append(finall, result)
				}
			}
		}
	}
	return finall
}
