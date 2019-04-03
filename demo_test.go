package leetcode

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	example := map[string]int{
		"abcabcbb": 3,
		" ":        1,
		"dvdf":     3,
		"aab":      2,
		"ohvhjdml": 6,
		"anviaj":   5,
		"abba":     2,
	}

	for input, expect := range example {
		output := lengthOfLongestSubstring(input)

		if expect != output {
			t.Errorf("input: %v, expect: %d, output:%d \n", input, expect, output)
		} else {
			t.Logf("input: %v, expect: %d, output:%d \n", input, expect, output)
		}
	}
}

func TestLengthOfLongestSubstringPro(t *testing.T) {

	example := map[string]int{
		"abcabcbb": 3,
		" ":        1,
		"dvdf":     3,
		"aab":      2,
		"ohvhjdml": 6,
		"anviaj":   5,
		"abba":     2,
	}

	for input, expect := range example {
		output := lengthOfLongestSubstringPro(input)
		if expect != output {
			t.Errorf("input: %v, expect: %d, output:%d \n", input, expect, output)
		} else {
			t.Logf("input: %v, expect: %d, output:%d \n", input, expect, output)
		}
	}
}

// func TestFindMedianSortedArrays(t *testing.T) {
// 	example := []struct {
// 		first  []int
// 		second []int
// 		expect float64
// 	}{
// 		{first: []int{1, 3}, second: []int{4, 5}, expect: 3.5},
// 		// {first: []int{1, 2}, second: []int{3, 4}, expect: 2.5},
// 	}

// 	for i, demo := range example {
// 		output := findMedianSortedArrays(demo.first, demo.second)
// 		t.Logf("\ndemo%d: %v \nexpect: %v, result: %v\n==================\n", i+1, demo, demo.expect, output)
// 	}
// }

func TestLongestPalindrome(t *testing.T) {
	example := map[string]string{
		"babad": "bab or aba",
		"cbbd":  "bb",
		"a":     "a",
		"":      "",
		"ccc":   "ccc",
		"abcba": "abcba",
	}

	for input, expect := range example {
		output := longestPalindrome(input)
		t.Logf("input: %v, expect: %s, output:%s \n", input, expect, output)
	}
}

func TestReverse(t *testing.T) {
	example := map[int]int{
		123:  321,
		-123: -321,
		120:  21,
		-120: -21,
	}

	for input, expect := range example {
		output := reverse(input)
		// t.Logf("input: %v, expect: %d, output:%d \n", input, expect, output)
		if expect != output {
			t.Errorf("input: %v, expect: %d, output:%d \n", input, expect, output)
		} else {
			t.Logf("input: %v, expect: %d, output:%d \n", input, expect, output)
		}
	}
}

func TestMyAtoi(t *testing.T) {
	example := map[string]int{
		"42":                    42,
		"   -42":                -42,
		"4193 with words":       4193,
		"words and 987":         0,
		"-91283472332":          -2147483648,
		"+1":                    1,
		"  0000000000012345678": 12345678,
		"+-2":                   0,
		"9223372036854775808":   2147483647,
		"0-1":                   0,
	}

	for input, expect := range example {
		output := myAtoi(input)
		// t.Logf("input: %v, expect: %d, output:%d \n", input, expect, output)
		if expect != output {
			t.Errorf("input: %v, expect: %d, output:%d \n", input, expect, output)
		} else {
			t.Logf("input: %v, expect: %d, output:%d \n", input, expect, output)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	example := map[int]bool{
		-121: false,
		10:   false,
		121:  true,
	}

	for input, expect := range example {
		output := isPalindrome(input)
		if expect != output {
			t.Errorf("input: %v, expect: %v, output:%v \n", input, expect, output)
		} else {
			t.Logf("input: %v, expect: %v, output:%v \n", input, expect, output)
		}
	}
}

type arr []int

func TestMaxArea(t *testing.T) {

	example := map[int][]int{
		49: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
	}

	for expect, input := range example {
		output := maxArea(input)
		if expect != output {
			t.Errorf("input: %v, expect: %v, output:%v \n", input, expect, output)
		} else {
			t.Logf("input: %v, expect: %v, output:%v \n", input, expect, output)
		}

	}
}

func TestIntToRoman(t *testing.T) {
	example := map[string]int{
		"I":       1,
		"IV":      4,
		"IX":      9,
		"LVIII":   58,
		"MCMXCIV": 1994,
		"XX":      20,
	}

	for expect, input := range example {
		output := intToRoman(input)

		if expect != output {
			t.Errorf("input: %v, expect: %s, output:%s \n", input, expect, output)
		} else {
			t.Logf("input: %v, expect: %s, output:%s \n", input, expect, output)
		}
	}
}

func TestRomanToInt(t *testing.T) {
	example := map[string]int{
		// "I":       1,
		// "IV":      4,
		// "IX":      9,
		"LVIII": 58,
		// "IX": 9,
		"MCMXCIV": 1994,
		// "XX":      20,
	}

	for input, expect := range example {
		output := romanToInt(input)

		if expect != output {
			t.Errorf("input: %v, expect: %d, output:%d \n", input, expect, output)
		} else {
			t.Logf("input: %v, expect: %d, output:%d \n", input, expect, output)
		}
	}
}

// if expect != output {
// 	} else {
// 	}

func TestLongestCommonPrefix(t *testing.T) {
	example := []struct {
		Input  []string
		Expect string
	}{
		{Input: []string{"flower", "flow", "flight"}, Expect: "fl"},
		{Input: []string{"dog", "racecar", "car"}, Expect: ""},
		{Input: []string{}, Expect: ""},
		{Input: []string{"aa", "a"}, Expect: "a"},
		{Input: []string{"react", "res", "respect"}, Expect: "re"},
	}

	for _, demo := range example {
		output := longestCommonPrefix(demo.Input)
		if demo.Expect != output {
			t.Errorf("input: %v, expect: %s, output:%s \n", demo.Input, demo.Expect, output)
		} else {
			t.Logf("input: %v, expect: %s, output:%s \n", demo.Input, demo.Expect, output)
		}

	}
}
