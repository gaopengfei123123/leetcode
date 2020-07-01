package leetcode

import (
	"bytes"
	"encoding/binary"
)

// // Max 取最大值
// func Max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// // Min 取最小值
// func Min(a, b int) int {
// 	if a > b {
// 		return b
// 	}
// 	return a
// }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// BytesToInt 字节转int
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}
