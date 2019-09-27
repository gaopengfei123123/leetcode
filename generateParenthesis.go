package leetcode

// 解题思路: 1. 递归; 2. 不做单条件返回, 而是多条件触发递归
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}

	leftRemain, rightRemain := n-1, n
	cur := []byte{'('}
	var res []string
	doGenerateParenthesis(leftRemain, rightRemain, cur, &res)
	return res
}

func doGenerateParenthesis(leftRemain, rightRemain int, cur []byte, res *[]string) {
	if leftRemain == 0 && rightRemain == 0 {
		*res = append(*res, string(cur))
		return
	}

	if leftRemain > 0 {
		doGenerateParenthesis(leftRemain-1, rightRemain, append(cur, '('), res)
	}
	if rightRemain > leftRemain {
		doGenerateParenthesis(leftRemain, rightRemain-1, append(cur, ')'), res)
	}
}
