package leetcode

// 思路: 从中间向两边开花(6)
func longestPalindrome(s string) string {
	strLen := len(s)
	if strLen <= 1 {
		return s
	}
	result := ""
	for i := 0; i < strLen; i++ {
		// 字符串单数的情况
		str1 := isPalindromex(i, i, s)
		// 字符串复数的情况
		str2 := isPalindromex(i, i+1, s)
		if len(str1) > len(result) {
			result = str1
		}
		if len(str2) > len(result) {
			result = str2
		}
	}
	return result
}

func isPalindromex(left, right int, s string) string {
	strLen := len(s)
	result := ""
	for left >= 0 && right < strLen {
		// fmt.Printf("currentStr: %s , index: %d %d\n", s[left:right], left, right)
		if s[left] == s[right] {
			result = s[left : right+1]
		} else {
			break
		}
		left--
		right++
	}
	return result
}

// //当使用 goroutine 时速度反而更慢了
// func longestPalindrome(s string) string {
// 	strLen := len(s)
// 	if strLen <= 1 {
// 		return s
// 	}
// 	result := ""
// 	resChan := make(chan string, 2)
// 	for i := 0; i < strLen; i++ {
// 		// 字符串单数的情况
// 		go isPalindrome(i, i, s, resChan)
// 		// 字符串复数的情况
// 		go isPalindrome(i, i+1, s, resChan)

// 		str1 := <-resChan
// 		str2 := <-resChan
// 		if len(str1) > len(result) {
// 			result = str1
// 		}
// 		if len(str2) > len(result) {
// 			result = str2
// 		}
// 	}
// 	return result
// }

// func isPalindrome(left, right int, s string, resChan chan string) string {
// 	strLen := len(s)
// 	result := ""
// 	for left >= 0 && right < strLen {
// 		if s[left] == s[right] {
// 			result = s[left : right+1]
// 		} else {
// 			break
// 		}
// 		left--
// 		right++
// 	}
// 	resChan <- result
// 	return result
// }
