package leetcode

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	phoneMap := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
	result := phoneMap[string(digits[0])]
	for i := 1; i < len(digits); i++ {
		item := phoneMap[string(digits[i])]
		tmp := []string{}
		for j := 0; j < len(item); j++ {
			for l := 0; l < len(result); l++ {
				value := result[l] + item[j]
				tmp = append(tmp, value)
			}
		}
		result = tmp
	}
	return result
}
