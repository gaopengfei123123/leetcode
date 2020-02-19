package leetcode

/**
这一道题的目的是给定一组数, 换算出一个比当前数字排列 "稍微大一点" 的数字排列
那么 123 "稍微大一点"的就是 132, 而不是 312, 213 这类的
解题思路:

正常情况下, 让数变大一点, 应该从右向左看, 遇到左边的数比右边小的情况, 那么这个节点就是要转换的节点

拿 124653 来举例, 从右向左发现 4 < 6, 那么4就应该和右边的一个数互换, 和谁换?
那么 4 就和 653 对比, 取其中 "刚好比它大" 的换, 那么5就比6合适, 互换后得到 125643
这个时候还没完, 因为我们只换了能转换的最高位中的最小值, 但是 643 还是没有 346小
因此我们需要的最终结果是 125346

根据以上的描述整理成编写步骤:
1. 自右向左应该是升序排列, 找到突然打破升序排列的数字k, 此时下标为t
2. 拿着数字k向右找, 找到刚好比k大一点的数, 两个数字位置互换
3. 然后下标 t 以右的位置的数字位置翻转, 从最开始自右向左升序变成自右向左降序
*/
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	// 1. 找到数字k的下标t
	t, i := len(nums)-2, len(nums)-1
	for t >= 0 && nums[t] >= nums[i] {
		t--
		i--
	}

	// k := nums[t]
	// 2. 找比k刚好大一点的数
	if t >= 0 {
		n := len(nums) - 1
		for nums[t] >= nums[n] { // 这里排除重复的数字
			n--
		}
		nums[t], nums[n] = nums[n], nums[t]
	}

	for l, r := t+1, len(nums)-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}
