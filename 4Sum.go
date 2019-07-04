package leetcode

import "sort"

// import "fmt"

func fourSum(nums []int, target int) [][]int {
	ln := len(nums)
	if ln < 4 {
		return nil
	}
	sort.Ints(nums)

	// fmt.Printf("nums: %v, traget: %d \n", nums, target)

	result := [][]int{}

	for f := 0; f < ln-2; f++ {
		// 去重
		if f > 0 && nums[f] == nums[f-1] {
			continue
		}

		for s := f + 1; s < ln-1; s++ {
			// 去重
			if s > f+1 && nums[s] == nums[s-1] {
				continue
			}

			tmpTarget := target - nums[f] - nums[s]
			// fmt.Printf("tmpTarget: %d , first: %d, second:%d \n", tmpTarget, nums[f], nums[s])

			l := s + 1
			r := ln - 1
			// fmt.Printf("f: %d, s: %d \n", f, s)
			for l < r {
				// fmt.Printf("l: %d, r: %d \n", l, r)
				tmpResult := nums[l] + nums[r]
				// fmt.Printf("right: %d, left: %d \n", nums[l], nums[r])
				if tmpResult > tmpTarget {
					r--
				} else if tmpResult < tmpTarget {
					l++
				} else {
					tmp := []int{nums[f], nums[s], nums[l], nums[r]}
					// fmt.Printf("tmpResult: %d, tempTarget: %d, tmp: %v \n", tmpResult, tmpTarget, tmp)
					result = append(result, tmp)
					r--
					l++
					for r-1 > s && nums[r] == nums[r+1] {
						r--
					}
					for l+1 < ln-1 && nums[l] == nums[l-1] {
						l++
					}
				}
			}
		}

	}
	return result
}
