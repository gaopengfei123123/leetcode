package leetcode

// 贪心算法
func intToRoman(num int) string {
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	intArr := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	result := ""
	for num > 0 {
	LOOP:
		for i := 0; i < len(intArr); i++ {
			if num >= intArr[i] {
				num = num - intArr[i]
				result += roman[i]
				break LOOP
			}
		}
	}
	return result
}
