package leetcode

// import "fmt"
import "sort"

// import "os"

func threeSum(nums []int) [][]int {
	var result [][]int

	n := len(nums)

	sort.Ints(nums)

	i := 0
	for i < n {
		if nums[i] > 0 {
			break
		}

		j := i + 1
		for j < n {
			if nums[i]+nums[j] > 0 {
				break
			}

			l := j + 1 // left range
			r := n - 1 // right range
			// binary search
			target := -nums[i] - nums[j]
			for l <= r {
				m := l + (r-l)/2
				if target == nums[m] {
					result = append(result, []int{nums[i], nums[j], nums[m]})
					break
				} else if target < nums[m] {
					r = m - 1
				} else {
					l = m + 1
				}
			}
			// next j
			j++
			// skip same element
			for j < n && nums[j-1] == nums[j] {
				j++
			}
		}
		// next i
		i++
		// skip same element
		for i < n && nums[i-1] == nums[i] {
			i++
		}
	}

	return result
}

// func threeSum(nums []int) [][]int {
// 	result := [][]int{}
// 	sort.Ints(nums)
// 	len := len(nums)
// 	fmt.Printf("%v \n", nums)

// 	var i, j int
// 	for i < len {
// 		first := nums[i]
// 		if first > 0 {
// 			break
// 		}

// 		j = i + 1
// 		for j < len {
// 			second := nums[j]
// 			if first+second > 0 {
// 				break
// 			}

// 			target := -first - second
// 			fmt.Printf("first: %v, second: %v, target: %v \n", first, second, target)

// 			left := j + 1
// 			right := len - 1

// 			for left <= right {

// 				mid := left + (right-1)/2
// 				middle := nums[mid]
// 				if target == middle {
// 					result = append(result, []int{first, second, middle})
// 					break
// 				} else if target < middle {
// 					right = mid - 1
// 				} else {
// 					left = mid + 1
// 				}
// 				fmt.Printf("left: %d, right: %v \n", left, right)
// 				fmt.Printf("mid: %v, middle: %v,  target: %v \n", mid, middle, target)
// 				os.Exit(0)
// 			}

// 			j++
// 			for j < len && nums[j-1] == nums[j] {
// 				j++
// 			}
// 		}

// 		i++
// 		for i < len && nums[i-1] == nums[i] {
// 			i++
// 		}
// 		os.Exit(0)
// 	}

// 	return result
// }
