package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	oddList := &ListNode{
		Val: head.Val,
	}

	oddCurr := oddList

	var evenList, evenCurr *ListNode

	idx := 1
	for {
		if head.Next == nil {
			break
		}

		if idx%2 == 0 {
			oddCurr.Next = &ListNode{
				Val:  head.Next.Val,
				Next: nil,
			}
			oddCurr = oddCurr.Next
		} else {
			if evenList == nil {
				evenList = &ListNode{
					Val:  head.Next.Val,
					Next: nil,
				}
				evenCurr = evenList
			} else {
				evenCurr.Next = &ListNode{
					Val:  head.Next.Val,
					Next: nil,
				}

				evenCurr = evenCurr.Next
			}
		}

		head = head.Next
		idx++
	}

	oddCurr.Next = evenList

	return oddList
}
