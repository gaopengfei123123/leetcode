package leetcode

import (
	// "fmt"
	"math"
	"sort"
)

// Runtime: 4 ms, faster than 99.11% of Go online submissions for 3Sum Closest.
// Memory Usage: 2.7 MB, less than 100.00% of Go online submissions for 3Sum Closest.
func threeSumClosest(nums []int, target int) int {
	var result int
	ln := len(nums)
	sort.Ints(nums)

	// taget在中间的情况
	result = nums[0] + nums[1] + nums[2]
	for i := 0; i < ln-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		l := ln - 1
		for j < l {
			tmp := nums[i] + nums[j] + nums[l]
			if math.Abs(float64(tmp-target)) < math.Abs(float64(result-target)) {
				result = tmp
			}
			if tmp == target {
				return tmp
			} else if tmp < target {
				j++
			} else {
				l--
			}
		}
	}
	return result
}
