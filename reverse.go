package leetcode

import (
	"math"
	"strconv"
)

// 略
func reverse(x int) int {
	iString := strconv.Itoa(x)
	iArr := []rune(iString)

	resultStr := ""
	minus := false
	for i := 0; i < len(iArr); i++ {
		if i == 0 {
			if string(iArr[0]) == "-" {
				minus = true
			} else {
				resultStr = string(iArr[i]) + resultStr
			}
		} else {
			resultStr = string(iArr[i]) + resultStr
		}
	}
	result, _ := strconv.Atoi(resultStr)
	if result > 2147483647 {
		result = 0
	}
	if minus {
		result = result * -1
	}

	return result
}

func reverseRro(x int) int {
	xabs := int(math.Abs(float64(x)))
	xx := ""
	if x < 0 {
		xx += "-"
	}
	for {
		if xabs >= 10 {
			xx += strconv.Itoa(xabs % 10)
			xabs /= 10
			continue
		}
		xx += strconv.Itoa(xabs)
		break
	}
	xxx, _ := strconv.Atoi(xx)
	if float64(xxx) > math.Pow(float64(2), float64(31))-1 || float64(xxx) < math.Pow(float64(-2), float64(31)) {
		return 0
	}
	return xxx
}
