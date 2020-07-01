package leetcode

import "fmt"

/*
动态规划
从左往右扫一遍, 每个下标都会保存一个来自左方的最大值
从右往左扫一遍, 保存来自右方的最大值
此时每个下标保存了左右两个最大值, 取其中最小的值与当前下标值相减,
就能得到这一格能存储雨水的最大值
时间复杂度O(n)
空间复杂度O(n)
*/
func trap(height []int) int {
	fmt.Printf("%#+v \n", height)
	size := len(height)

	if size < 2 {
		return 0
	}

	leftMax := make([]int, size)
	rightMax := make([]int, size)

	leftMax[0] = height[0]
	for i := 1; i < size; i++ {
		leftMax[i] = max(height[i], leftMax[i-1])
	}
	rightMax[size-1] = height[size-1]
	for i := size - 2; i >= 0; i-- {
		rightMax[i] = max(height[i], rightMax[i+1])
	}
	result := 0
	for i := 0; i < size-1; i++ {
		result += min(leftMax[i], rightMax[i]) - height[i]
	}
	return result
}

/**
使用栈的方式
设置一个只存下标的栈(这里用数组表示)
把盛雨水的图形看做一个碗, 栈只存碗左边的下标, 栈中对应的值从下向上呈下降趋势
从左向右遍历时一旦出现上升趋势的下标, 说明能构成一个"碗", 就要开始干活了
这个思路和 Longest Valid Parentheses 的类似, 一层层的计算出每个满足条件的结果, 最后相加
时间复杂度O(n)
空间复杂度O(n)
*/
func trapStack(height []int) int {
	result := 0
	if len(height) < 2 {
		return result
	}

	st := []int{}

	current := 0
	for current < len(height) {
		// 调试用
		// fmt.Printf("all stack %v \n", st)

		for len(st) > 0 && height[current] > height[st[len(st)-1]] {
			// 调试用
			// fmt.Printf("work stack %v \n", st)

			// 最低点的下标
			top := st[len(st)-1]
			st = st[0 : len(st)-1]
			if len(st) == 0 {
				break
			}

			// st[len(st)-1]是次低点的下标
			// current - st[len(st)-1] 计算碗的宽度, 因为有"边界", 所以要减一
			// 而且这个宽度下"碗"的高度是一致的
			dst := current - st[len(st)-1] - 1
			// 计算碗的高度
			h := min(height[current], height[st[len(st)-1]]) - height[top]

			// 调试用
			// fmt.Printf("top: %d, st.Top: %d, current: %d \ndst %d, h: %d \n", top, st[len(st)-1], current, dst, h)
			result += dst * h
		}
		st = append(st, current)
		current++
	}

	return result
}

/**
双指针解法
和第一个方法一样, 不过更简单直接一些,
从左边往右边遍历时只要最右边比左边高, 就认为能构成"碗" 就开始计算盛放的体积
同理从左往右遍历也是一样

时间复杂度: O(n)
空间复杂度: O(1)
*/
func trapPointer(height []int) int {
	result := 0
	left := 0
	right := len(height) - 1
	leftMax := 0
	rightMax := 0
	for left < right {
		if height[left] < height[right] {
			if height[left] < leftMax {
				result += leftMax - height[left]
			} else {
				leftMax = height[left]
			}
			left++
		} else {
			if height[right] < rightMax {
				result += rightMax - height[right]
			} else {
				rightMax = height[right]
			}
			right--
		}
	}
	return result
}
