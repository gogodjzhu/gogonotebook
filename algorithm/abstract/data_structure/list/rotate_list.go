package list

import . "gogonotebook/common"

type RotateListSolver struct {
}

func (RotateListSolver) RotateListSolve(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}
	r := head
	total := 1
	for r.Next != nil {
		r = r.Next
		total += 1
	}
	r.Next = head // 链表成环

	idx := total - k%total
	h := head
	for i := 0; i < idx-1; i++ {
		h = h.Next
	}
	ret := h.Next
	h.Next = nil
	return ret
}
