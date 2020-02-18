package leetcode

// import "fmt"

func reverseString(s []byte) {
	ln := len(s)
	if ln <= 1 {
		return
	}
	s[0], s[ln-1] = s[ln-1], s[0]
	reverseString(s[1 : ln-1])
}
