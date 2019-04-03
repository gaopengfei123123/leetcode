package leetcode

import "fmt"

// 未做完
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var leftArr, rightArr []int
	if nums1[0] < nums2[0] {
		leftArr = nums1
		rightArr = nums2
	} else {
		leftArr = nums2
		rightArr = nums1
	}

	leftLen := len(leftArr)
	rightLen := len(rightArr)
	maxLen := leftLen + rightLen

	lHead := leftArr[0]
	lFoot := leftArr[leftLen-1]
	rHead := rightArr[0]
	rFoot := rightArr[rightLen-1]

	middle := 0

	// 分离
	if lFoot > rHead && lFoot > rFoot {
		remainder := maxLen % 2
		// 偶数
		if remainder == 0 {
			middle = int(maxLen / 2)
			fmt.Printf("middle: %d; maxLen: %d  \n", middle, maxLen)

			var seed1, seed2 int

			if leftLen >= middle {
				seed1 = leftArr[middle-1]
			} else {
				seed1 = rightArr[middle-leftLen]
			}

			if leftLen >= middle+1 {
				seed2 = leftArr[middle]
			} else {
				seed2 = rightArr[middle-leftLen+1]
			}

			fmt.Printf("leftArr: %v , rightArr: %v \n", leftArr, rightArr)
			fmt.Printf("middle: %d; maxLen: %d  seed1: %d  seed2: %d \n", middle, maxLen, seed1, seed2)
			// 奇数
		} else {

		}

	}

	// 交错
	if lFoot <= rHead && lFoot >= rFoot {
		for i := 0; i < rightLen; i++ {
			if rightArr[i] >= lFoot {
				maxLen -= i
			}
		}
	}

	// 包含
	if lHead < rHead && lFoot > rFoot {

	}

	fmt.Println("easy", middle)

	return 0
}
