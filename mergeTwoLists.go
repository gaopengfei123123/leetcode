package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l2 == nil && l1 == nil {
		return nil
	}
	if l1 == nil && l2 != nil {
		return l2
	}
	if l2 == nil && l1 != nil {
		return l1
	}

	if l1.Val < l2.Val {
		tmp := l1.Next
		l1.Next = mergeTwoLists(tmp, l2)
		return l1
	}

	tmp := l2.Next
	l2.Next = mergeTwoLists(tmp, l1)
	return l2
}
