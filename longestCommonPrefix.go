package leetcode

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	result := strs[0]
	for i := 1; i < len(strs); i++ {
		target := strs[i]
		currentIndex := len(result)
		targetIndex := len(target)
		if targetIndex < currentIndex {
			currentIndex = targetIndex
		}
	LOOP:
		for currentIndex > 0 {
			if result[:currentIndex] == target[:currentIndex] {
				result = result[:currentIndex]
				break LOOP
			}
			currentIndex--
			result = result[:currentIndex]
		}
		if currentIndex == 0 {
			return ""
		}
	}
	return result
}
