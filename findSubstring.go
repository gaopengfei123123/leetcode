package leetcode

import "fmt"

/**
解题思路:
这一题的关键点在于 "words 的长度相同"
这意味着采用移动窗口的方法对字符串检索的时候, 是以单个word的长度为单位跳跃的

因此, 如下的检索是以 `wordLen * wordNums` 为一整块, 从左到右以 wordLen 为 step 滑动
同时因为是单词, 所以起始点需要以 单个 wordLen 循环一遍, 这样能覆盖到全部字段
在函数中就体现在两个 for 循环上, 第一层循环单个 wordLen, 以对齐各个word, 第二层循环进行分段检测

*/
func findSubstring(s string, words []string) []int {
	if len(words) == 0 || len(s) == 0 || len(s) < len(words[0])*len(words) {
		return []int{}
	}
	wordLen := len(words[0])
	wordNums := len(words)
	sLen := len(s)
	res := make([]int, 0)
	l := wordLen * wordNums
	var left, cursor int
	right := l
	reset := true
	var m map[string]int8
	for i := 0; i < wordLen; i++ {
		left = i
		cursor = left
		right = left + l
		fmt.Printf(">>>>>>>>> LOOP: left: %d, right: %d, string: %s \n", left, right, s[left:right])
		if reset {
			reset = false
			m = map[string]int8{}
			for _, w := range words {
				m[w]++
			}
		}

		for cursor < right && right <= sLen {
			index := s[cursor : cursor+wordLen]
			how, ok := m[index]

			fmt.Printf("left: %d, cursor: %d, right: %d, string: %s  \n", left, cursor, right, s[left:right])
			fmt.Printf("index: %#+v, how: %#+v, ok: %#+v \n\n", index, how, ok)

			if !ok {
				if reset {
					reset = false
					m = map[string]int8{}
					for _, w := range words {
						m[w]++
					}
				}
				left = cursor + wordLen
				cursor = left
				right = left + l
				continue
			}

			if how <= 0 {
				fmt.Printf("how < 0 , %s \n", s[left:left+wordLen])
				m[s[left:left+wordLen]]++
				left += wordLen
				right += wordLen
				continue
			}

			reset = true
			m[index]--
			cursor += wordLen
			if cursor == right {
				res = append(res, left)
				m[s[left:left+wordLen]]++
				left += wordLen
				right += wordLen
			}

		}
	}
	return res
}

// func findSubstring(s string, words []string) []int {
// 	if len(words) == 0 || len(s) == 0 || len(s) < len(words[0])*len(words) {
// 		return []int{}
// 	}

// 	sLen := len(words[0])
// 	aLen := len(words[0]) * len(words)
// 	res := []int{}

// 	for i := 0; i < len(s)-aLen+1; i++ {
// 		tmpS := s[i : aLen+i]
// 		// fmt.Printf("%#+v \n", tmpS)
// 		mp := make(map[string]int, len(words))
// 		for _, key := range words {
// 			if _, ok := mp[key]; ok {
// 				mp[key]++
// 			} else {
// 				mp[key] = 1
// 			}
// 		}
// 		verify := false
// 	LOOP:
// 		for j := 0; j < len(tmpS); j += sLen {
// 			chunk := tmpS[j : j+sLen]
// 			// fmt.Printf("%#+v \n", chunk)

// 			v, ok := mp[chunk]
// 			if !ok || v <= 0 {
// 				break LOOP
// 			}
// 			mp[chunk]--

// 			if j == len(tmpS)-sLen {
// 				verify = true
// 			}
// 			// fmt.Printf("i: %d, j:%d, tmpMap: %#+v \n", i, j, mp)
// 		}
// 		if verify {
// 			res = append(res, i)
// 		}
// 	}

// 	return res
// }
