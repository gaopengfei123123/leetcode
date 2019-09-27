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

func TestThreeSum(t *testing.T) {
	example := []struct {
		Input  []int
		Expect [][]int
	}{
		{
			Input: []int{-1, 0, 1, 2, -1, -4},
			Expect: [][]int{
				[]int{-1, 0, 1},
				[]int{-1, -1, 2},
			},
		},
		{
			Input:  []int{1, 2, -2, -1},
			Expect: [][]int{},
		},
		{
			Input: []int{0, 0, 0},
			Expect: [][]int{
				[]int{0, 0, 0},
			},
		},
		{
			Input: []int{0, 0, 0, 0},
			Expect: [][]int{
				[]int{0, 0, 0},
			},
		},
		{
			Input: []int{-1, 0, 1, 2, -1, -4},
			Expect: [][]int{
				[]int{-1, -1, 2},
				[]int{-1, 0, 1},
			},
		},
		{
			Input: []int{2, -2, 9, -9, 7, -7, 2, -7, 0, 3, 8, -9, -3, -9, -3, -10, -5, -4, -3, -9, -9, -4, 0, 3, -10, -7},
			Expect: [][]int{
				[]int{-10, 2, 8},
				[]int{-10, 3, 7},
				[]int{-9, 0, 9},
				[]int{-9, 2, 7},
				[]int{-7, -2, 9},
				[]int{-7, 0, 7},
				[]int{-5, -4, 9},
				[]int{-5, -3, 8},
				[]int{-5, -2, 7},
				[]int{-5, 2, 3},
				[]int{-4, -4, 8},
				[]int{-4, -3, 7},
				[]int{-4, 2, 2},
				[]int{-3, 0, 3},
				[]int{-2, 0, 2},
			},
		},
	}

	for _, demo := range example {
		output := threeSum(demo.Input)
		t.Logf("input: %v \n expect: %v \n output:%v \n", demo.Input, demo.Expect, output)
	}
}

func TestThreeSumClosest(t *testing.T) {
	example := []struct {
		Input   []int
		Target  int
		Closest int
	}{
		{
			Input:   []int{-1, 2, 1, -4},
			Target:  1,
			Closest: 2,
		},
		{
			Input:   []int{1, 1, 1, 0},
			Target:  100,
			Closest: 3,
		},
		{
			Input:   []int{0, 2, 1, -3},
			Target:  1,
			Closest: 0,
		},
		{
			Input:   []int{1, 1, -1, -1, 3},
			Target:  3,
			Closest: 3,
		},
	}

	for _, demo := range example {
		output := threeSumClosest(demo.Input, demo.Target)
		t.Logf("input: %v \n expect: %v \n output:%v \n", demo.Input, demo.Closest, output)
	}
}

func TestLetterCombinations(t *testing.T) {
	example := []struct {
		Input  string
		Expect []string
	}{
		{
			Input:  "23",
			Expect: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			Input:  "",
			Expect: []string{},
		},
	}
	for _, demo := range example {
		output := letterCombinations(demo.Input)
		t.Logf("input: %v \n expect: %v \n output:%v \n", demo.Input, demo.Expect, output)
	}
}

func TestFourSum(t *testing.T) {
	example := []struct {
		Input  []int
		Target int
		Expect [][]int
	}{
		{
			Input:  []int{1, 0, -1, 0, -2, 2},
			Target: 0,
			Expect: [][]int{
				[]int{-1, 0, 0, 1},
				[]int{-2, -1, 1, 2},
				[]int{-2, 0, 0, 2},
			},
		},

		{
			Input:  []int{-3, -2, -1, 0, 0, 1, 2, 3},
			Target: 0,
			Expect: [][]int{
				[]int{-3, -2, 2, 3},
				[]int{-3, -1, 1, 3},
				[]int{-3, 0, 0, 3},
				[]int{-3, 0, 1, 2},
				[]int{-2, -1, 0, 3},
				[]int{-2, -1, 1, 2},
				[]int{-2, 0, 0, 2},
				[]int{-1, 0, 0, 1},
			},
		},
	}

	for _, demo := range example {
		output := fourSum(demo.Input, demo.Target)
		t.Logf("input: %v \n expect: %v \n output:%v \n", demo.Input, demo.Expect, output)
	}
}

func TestRemoveNthFromEnd(t *testing.T) {
	example := []struct {
		Input  *ListNode
		N      int
		Expect *ListNode
	}{
		{
			Input:  array2SingleList([]int{1, 2, 3, 4, 5}),
			N:      2,
			Expect: array2SingleList([]int{1, 2, 3, 5}),
		},
	}

	for _, demo := range example {
		output := removeNthFromEnd(demo.Input, demo.N)
		echoList(output)
		echoList(demo.Expect)
	}
}

func TestIsValid(t *testing.T) {
	example := []struct {
		Input string

		Expect bool
	}{
		{
			Input:  "]",
			Expect: false,
		},
		{
			Input:  "()",
			Expect: true,
		},
		{
			Input:  "({[",
			Expect: false,
		},
		{
			Input:  "",
			Expect: true,
		},

		{
			Input:  "()[]{}",
			Expect: true,
		},
		{
			Input:  "(]",
			Expect: false,
		},
		{
			Input:  "{[]}",
			Expect: true,
		},
		{
			Input:  "([)]",
			Expect: false,
		},
		{
			Input:  "(([]){})",
			Expect: true,
		},
	}

	for _, demo := range example {
		output := isValid(demo.Input)
		t.Logf("\n  input: %v \n expect: %v \n output: %v \n", demo.Input, demo.Expect, output)
	}
}

func TestMergeTwoLists(t *testing.T) {
	example := []struct {
		Node1 *ListNode
		Node2 *ListNode

		Expect *ListNode
	}{
		{
			Node1:  array2SingleList([]int{1, 2, 4}),
			Node2:  array2SingleList([]int{1, 3, 4}),
			Expect: array2SingleList([]int{1, 1, 2, 3, 4, 4}),
		},
	}

	for _, demo := range example {
		output := mergeTwoLists(demo.Node1, demo.Node2)
		echoList(output)
		// echoList(demo.Expect)
	}
}

func TestGenerateParenthesis(t *testing.T) {
	example := []struct {
		N      int
		Expect []string
	}{
		// {
		// 	N: 1,
		// 	Expect: []string{
		// 		"()", // 1
		// 	},
		// },
		// {
		// 	N: 2,
		// 	Expect: []string{
		// 		"()()", // 1+1
		// 		"(())", // 2
		// 	},
		// },
		{
			N: 3,
			Expect: []string{
				"((()))", // 3
				"(()())", // 1.5 + 1.5
				"(())()", // 2+1
				"()(())", // 1+2
				"()()()", // 1+1+1
			},
		},
	}

	for _, demo := range example {
		output := generateParenthesis(demo.N)
		t.Logf("\n  input: %v \n expect: %v \n output: %v \n", demo.N, demo.Expect, output)
	}
}
