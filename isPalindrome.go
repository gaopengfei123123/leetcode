package leetcode

// 略
func isPalindrome(x int) bool {
	bak := x
	if x < 0 {
		return false
	}
	reverse := 0
	for x > 0 {
		current := x % 10
		reverse = reverse*10 + current
		x = (x - current) / 10
	}
	return bak == reverse
}
