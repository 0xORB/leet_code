package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) MiddleNode() *ListNode {
	middle, end := head, head
	for end != nil && end.Next != nil {
		middle = middle.Next
		end = end.Next.Next
	}
	return middle
}
