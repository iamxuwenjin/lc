// Package code
// @Description: 141. Linked List Cycle 环形列表
package code

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	f := head.Next
	l := head
	for f != l {
		if f == nil || f.Next == nil {
			return false
		}
		l = l.Next
		f = f.Next.Next
	}
	return true
}
