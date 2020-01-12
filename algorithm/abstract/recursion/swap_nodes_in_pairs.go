package recursion

import . "gogonotebook/common"

type SwapNodesInPairsSolver struct {
}

func (s *SwapNodesInPairsSolver) SwapNodesInPairsSolve(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ret := head.Next
	head.Next = s.SwapNodesInPairsSolve(head.Next.Next)
	ret.Next = head
	return ret
}
