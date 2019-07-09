package leetcode

import "fmt"

// ListNode 链表结构
type ListNode struct {
	Val  int
	Next *ListNode
}

var level int

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return head
	}
	preHeader := &ListNode{Next: head}
	left, right := preHeader, head

	for i := 0; right != nil && i < n; i++ {
		right = right.Next
	}
	for right != nil {
		left, right = left.Next, right.Next
	}
	left.Next = left.Next.Next
	return preHeader.Next
}

func array2SingleList(nums []int) *ListNode {
	var next *ListNode
	for i := 0; i < len(nums); i++ {
		current := &ListNode{}
		current.Val = nums[i]
		current.Next = next
		next = current
	}

	// echoList(next)
	return next
}

func echoList(head *ListNode) {
	fmt.Printf("val: %d, level: %d \n", head.Val, level)
	level++
	if head.Next != nil {
		echoList(head.Next)
	} else {
		level = 0
	}
}
