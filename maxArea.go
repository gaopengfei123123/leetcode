package leetcode

// 从两边向中间收敛, 高度低的向内偏移一位 复杂度 O(n)
func maxArea(height []int) int {
	leftIndex := 0
	rightIndex := len(height) - 1
	width := len(height)
	result := 0

	for leftIndex < rightIndex {
		long := 0
		if height[leftIndex] > height[rightIndex] {
			long = height[rightIndex]
			rightIndex--
		} else {
			long = height[leftIndex]
			leftIndex++
		}
		width--
		tmp := width * long
		if tmp > result {
			result = tmp
		}
	}
	return result
}

// 硬解法  复杂度为 O(n^2)
func maxAreaBak(height []int) int {
	result := 0
	for i := 0; i < len(height); i++ {
		for j := i; j < len(height); j++ {
			tmp := area([]int{height[i], i}, []int{height[j], j})
			if tmp > result {
				result = tmp
			}
		}
	}
	return result
}

// 计算面积
func area(a []int, b []int) int {
	ah := a[0]
	aw := a[1]
	bh := b[0]
	bw := b[1]

	width := 0
	height := 0
	if aw > bw {
		width = aw - bw
	} else {
		width = bw - aw
	}

	if bh > ah {
		height = ah
	} else {
		height = bh
	}
	return width * height
}
