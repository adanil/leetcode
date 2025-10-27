package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	curr := result

	debt := 0
	for {
		var val1, val2 int

		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		s := val1 + val2 + debt
		debt = 0
		if s >= 10 {
			debt = 1
		}

		curr.Val = s % 10

		if l1 == nil && l2 == nil {
			break
		}

		curr.Next = &ListNode{}
		curr = curr.Next

	}

	if debt > 0 {
		curr.Next = &ListNode{Val: debt}
	}
	return result
}

func main() {

}
