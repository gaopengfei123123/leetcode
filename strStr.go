package leetcode

import "fmt"

func strStr(haystack string, needle string) int {
	result := -1
	nLen := len(needle)
	hLen := len(haystack)

	if hLen < nLen {
		return result
	}

	if haystack == needle {
		return 0
	}

	for i := 0; i < hLen-nLen+1; i++ {
		fmt.Println(haystack[i:i+nLen], needle)
		if haystack[i:i+nLen] == needle {
			return i
		}
	}

	return result
}
