package leetcode

/**
解题思路:
	题中要求时间复杂度必须 O(logN) 而且这是有序数组,果断使用二分法去搞
	有一种思路是先利用二分法找到target的位置, 然后左右扩位置找到最大值最小值
	这种可以做, 不过有条件, 就是重复数字不能太多, 比如 [0,0,0,0,0,0,0] target:0, 这种时间复杂度直接O(N)

	出于以上考虑, 我还是选择执行两次二分查找去分别查最大值和最小值, 而且即使找到target也要继续找
	二分查找难点就在+1-1的临界值上, 这种还是要仔细考虑一下

*/
func searchRange(nums []int, target int) []int {
	lo := searchOneSide(nums, target, true)  // 获取最小值
	hi := searchOneSide(nums, target, false) // 获取最大值
	return []int{lo, hi}
}

func searchOneSide(nums []int, target int, isLeft bool) int {
	res := -1
	lo := 0
	hi := len(nums) - 1
	mid := lo + (hi-lo)/2
	for lo <= hi {
		// fmt.Printf("nums[%d]: %d, nums[%d]: %d, nums[%d]: %d \n", lo, nums[lo], mid, nums[mid], hi, nums[hi])
		if target == nums[mid] {
			// fmt.Printf("target: %d, mid: %d \n", target, mid)
			res = mid
			if isLeft {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
			mid = lo + (hi-lo)/2
			continue
		}

		if target < nums[mid] {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
		mid = lo + (hi-lo)/2
		// fmt.Printf("lo: %d, mid: %d hi: %d\n", lo, mid, hi)
	}
	return res
}
