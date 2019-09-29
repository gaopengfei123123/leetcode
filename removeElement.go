package leetcode

// 注意由于golang切片未扩容时会共享内存, 因此也需要改变数组中的值
func removeElement(nums []int, val int) int {
	nlen := len(nums)
	reslen := nlen
	curIdx := 0
	for i := 0; i < nlen; i++ {
		if nums[i] == val {
			reslen--
			continue
		}
		nums[curIdx] = nums[i]
		curIdx++
	}
	return reslen
}
