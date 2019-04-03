package leetcode

import (
// "fmt"
)

// 思路: 基于字符串进行截取并对最终字符串进行遍历验证, 时间复杂度 <= O(2n)
func lengthOfLongestSubstring(s string) int {
	maxLen := len(s)
	tmpStr := ""
	tmpArr := make(map[string]int, 1)
	result := 0

	for i := 0; i < maxLen; i++ {
		current := s[i : i+1]
		index, ok := tmpArr[current]
		if !ok {
			tmpStr += current
			tmpArr[current] = len(tmpArr)
		} else {
			tmpStr = tmpStr[index:] + current
			tmpStr = tmpStr[1:]
			tmpArr = map[string]int{}
			for j := 0; j < len(tmpStr); j++ {
				key := tmpStr[j : j+1]
				tmpArr[key] = j
			}

		}
		if len(tmpStr) > result {
			result = len(tmpStr)
		}
	}

	return result
}

// 思路: 验证同一个字符重复出现时最长的间隔, O(n)
func lengthOfLongestSubstringPro(s string) int {
	tmpArr := make(map[string]int, 1)
	maxLen := len(s)
	result := 0
	lastPos := 0

	for i := 0; i < maxLen; i++ {
		current := s[i : i+1]
		index, exist := tmpArr[current]
		if exist {
			lastPos = max(index, lastPos)
		}
		result = max(result, i-lastPos+1)
		tmpArr[current] = i + 1
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
