package heap

import . "gogonotebook/common"

type MergeKSortedListSolver struct {
}

func (m MergeKSortedListSolver) MergeKSortedListSolve(lists []*ListNode) *ListNode {
	h := NewHeap()
	count := 0
	for _, list := range lists {
		for list != nil {
			h.Put(list.Val)
			count += 1
			list = list.Next
		}
	}
	return h.ToList()
}
