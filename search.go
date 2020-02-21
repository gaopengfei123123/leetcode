package leetcode

/**
解题思路:
	没啥好说的, 就二分法查找
*/
func search(nums []int, target int) int {

	left := 0
	right := len(nums) - 1
	mid := left + (right-left)/2

	for left <= right {
		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] { //左边升序
			if target >= nums[left] && target <= nums[mid] { //在左边范围内
				right = mid - 1
			} else { //只能从右边找
				left = mid + 1
			}
		} else { //右边升序
			if target >= nums[mid] && target <= nums[right] { //在右边范围内
				left = mid + 1
			} else { //只能从左边找
				right = mid - 1
			}
		}
		mid = left + (right-left)/2
	}

	return -1
}
