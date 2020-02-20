package leetcode

/**
解题思路:
1. 利用括号成对的特性, 左右括号相等
	有一个 `(` 那必然会有一个 `)` , 计左右括号出现的次数, 次数相等的算作一次合并
	这个方法有局限性在如果出现 `(()` 这种情况单次遍历不能统计到,
	作为弥补就需要正反两次遍历, 取最大值
	时间复杂度 O(n), 空间复杂度 O(1)

2. 采取动态规划的求解方式
	因为括号和套娃似的, 外层括号可以直接采用内层的处理结果, 不过要另建一个数组来缓存结果
	创建缓存数组dp,  比如:
	s = []string{"(", "(", ")", "(", ")", ")"}
	dp = []int{0, 0, 2, 0, 4, 6}
	在合法括号中左括号的下标和dp映射, 合法括号的长度为对应下标的元素值
	那么上述例子中 下标3存在一个合法的括号, 长度为2, 这是 () 这种组合
	下标为4存在一个合法括号, 这也是一个()组合, 但是不会计算第一个括号了, 直接到下标3拿计算结果

	到下标5时, 存在 )) 这种组合, 那么已知下标4是有值的(代表是一个合法括号) 那么到下标5就直接往前推 4 个长度
	得到下标 0 是 (, 下标5是 ), 括号成立,  因此内层括号+2 就是这个的括号长度
	时间复杂度 O(n), 空间复杂度 O(n)

3. 使用栈
	栈的特性是后进先出, 括号也有类似的特性, 因此可以这么搞
	我们首先将 −1 放入栈顶。
	对于遇到的每个 ( ，我们将它的下标放入栈中。
	对于遇到的每个 ) ，我们弹出栈顶的元素并将当前元素的下标与弹出元素下标作差，得出当前有效括号字符串的长度。通过这种方法，我们继续计算有效子字符串的长度，并最终返回最长有效子字符串的长度。
	时间复杂度 O(n), 空间复杂度 O(n)
	作者：LeetCode
	链接：https://leetcode-cn.com/problems/longest-valid-parentheses/solution/zui-chang-you-xiao-gua-hao-by-leetcode/
	来源：力扣（LeetCode）
	著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}

	left := byte('(')

	var rCnt, lCnt, length int

	for i := 0; i < len(s); i++ {
		if s[i] == left {
			lCnt++
		} else {
			rCnt++
		}
		if rCnt > lCnt { // 从左往右必死的情况
			lCnt, rCnt = 0, 0
		}
		if rCnt == lCnt && rCnt*2 > length {
			length = rCnt * 2
		}
	}

	lCnt, rCnt = 0, 0
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == left {
			lCnt++
		} else {
			rCnt++
		}

		if rCnt < lCnt { // 从右往左必死的情况
			lCnt, rCnt = 0, 0
		}
		if rCnt == lCnt && rCnt*2 > length {
			length = rCnt * 2
		}
	}

	return length
}

func longestValidParenthesesV2(s string) int {
	if len(s) < 2 {
		return 0
	}
	dp := make([]int, len(s))
	left := byte('(')
	right := byte(')')

	for i := 1; i < len(s); i++ {
		if s[i] == right {
			if s[i-1] == left { // 如 () 这样的情况
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
				continue
			}

			pre := dp[i-1]
			if (i-pre-1) >= 0 && s[i-1] == right { // 如 )) 这样的情况
				if s[i-pre-1] == left {
					dp[i] = dp[i-1] + 2
					if i-dp[i] > 0 { // 如果存在紧挨着的两个合法组合, 就拼起来
						dp[i] += dp[i-dp[i]]
					}
				}
			}
		}
	}

	// fmt.Printf("\n %#+v \n", dp)
	var maxLen int
	for i := 0; i < len(s); i++ {
		if maxLen < dp[i] {
			maxLen = dp[i]
		}
	}
	return maxLen
}
