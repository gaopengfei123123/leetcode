package leetcode

func romanToInt(s string) int {
	symbol := map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}
	result := 0

	for len(s) > 0 {
		if len(s) >= 2 {
			head := s[0:2]
			value, ok := symbol[head]
			if ok {
				result += value
				s = s[2:]
			} else {
				head = s[0:1]
				value, _ := symbol[head]
				result += value
				s = s[1:]
			}
		} else {
			head := s[0:1]
			value, _ := symbol[head]
			result += value
			s = s[1:]
		}

	}
	return result
}
