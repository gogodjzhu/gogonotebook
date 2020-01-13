package list

import . "gogonotebook/common"

type OddEvenLinkedListSolver struct {
}

func (o OddEvenLinkedListSolver) OddEvenLinkedListSolve(head *ListNode) *ListNode {
	oddDummy := &ListNode{}
	odd := oddDummy
	evenDummy := &ListNode{}
	even := evenDummy
	for idx := 1; head != nil; idx++ {
		if idx%2 == 1 {
			odd.Next = head
			odd = head
		} else {
			even.Next = head
			even = head
		}
		head = head.Next
	}
	odd.Next = evenDummy.Next
	even.Next = nil
	return oddDummy.Next
}
