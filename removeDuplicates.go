package leetcode

// 注意由于golang切片未扩容时会共享内存, 因此也需要改变数组中的值
func removeDuplicates(nums []int) int {
	nlen := len(nums)
	if nlen == 0 {
		return 0
	}

	curIdx, prev := 0, nums[0]
	for i := 1; i < nlen; i++ {
		if nums[i] != prev {
			curIdx++
			nums[curIdx] = nums[i]
			prev = nums[i]
		}
	}
	return curIdx + 1
}
