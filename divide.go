package leetcode

import (
	"math"
	"fmt"
)

/**
解题思路:
	除法可以换算成多次减法, 但是如果碰到数字很大,除数很小的时候就比较糟糕, 时间复杂度在 O(N), 比较极端的例子就是 2^32/1 就要循环2^32次
	但是采用位运算将减少很多循环次数, 时间复杂度在 O(logn), 那么  2^32/1 这种就只需要循环32次

	首先判定商的正负号, 然后判定超界以及商为0的情况, 然后[被除数]-[除数]x2^n , 每次成功减去时累加n
	以此循环直到[被除数] < [除数]

	举例如: 100 / 3
	输出:
>>>>>>>start loop, dividend: 100, divisor: 3 <<<<<<<

i: 0
dividend: 97 (100-3x2^0)
quotient: 1 (0+1)

i: 1
dividend: 91 (97-3x2^1)
quotient: 3 (1+2)

i: 2
dividend: 79 (91-3x2^2)
quotient: 7 (3+4)

i: 3
dividend: 55 (79-3x2^3)
quotient: 15 (7+8)

i: 4
dividend: 7 (55-3x2^4)
quotient: 31 (15+16)

>>>>>>> next loop <<<<<<<

>>>>>>>start loop, dividend: 7, divisor: 3 <<<<<<<

i: 0
dividend: 4 (7-3x2^0)
quotient: 32 (31+1)

>>>>>>> next loop <<<<<<<

>>>>>>>start loop, dividend: 4, divisor: 3 <<<<<<<

i: 0
dividend: 1 (4-3x2^0)
quotient: 33 (32+1)

>>>>>>> next loop <<<<<<<

	得出商为 33
*/
func divide(dividend int, divisor int) int {
	quotient := 0
	sign := 1

	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	if dividend < 0 && divisor < 0 {
		dividend = -dividend
		divisor = -divisor
	} else if dividend < 0 {
		dividend = -1
		sign = -1
	} else if divisor < 0 {
		divisor = -1 
		sign = -1
	}

	for dividend >= divisor {
		multiple := divisor
		fmt.Printf(">>>>>>>start loop, dividend: %d, divisor: %d <<<<<<< \n\n", dividend,divisor)
		for i:=0; dividend >= multiple;i++{
			fmt.Printf("i: %d \n", i)
			dividend -= multiple
			fmt.Printf("dividend: %d (%d-%dx%d^%d) \n", dividend, dividend + multiple, divisor,2, i)
			quotient += 1 << uint(i)
			fmt.Printf("quotient: %d (%d+%d) \n\n", quotient, quotient - 1 << uint(i), 1 << uint(i))
			multiple <<=1
		}
		fmt.Printf(">>>>>>> next loop <<<<<<< \n\n\n")
	}

	if sign == -1 {
		return -quotient
	}
	return quotient
}

// func divide(dividend int, divisor int) int {
// 	quotient := 0
// 	sign := 1

// 	// 防止越界
// 	if dividend == math.MinInt32 && divisor == -1 {
// 		return math.MaxInt32
// 	}

// 	// capture the sign
// 	if divisor < 0 && dividend < 0 {
// 		divisor = -divisor
// 		dividend = -dividend

// 	} else if divisor < 0 {
// 		divisor = -divisor
// 		sign = -1

// 	} else if dividend < 0 {
// 		dividend = -dividend
// 		sign = -1

// 	}

// 	// divide some part to approach the answer
// 	for dividend >= divisor {
// 		multiple := divisor
// 		for i := 0; dividend >= multiple; i++ {
// 			dividend -= multiple
// 			quotient += 1 << uint(i)
// 			multiple <<= 1
// 		}
// 	}

// 	if sign == -1 {
// 		return -quotient
// 	}

// 	return quotient
// }