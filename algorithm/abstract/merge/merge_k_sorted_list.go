package merge

import (
	"gogonotebook/algorithm/abstract/data_structure/heap"
	. "gogonotebook/common"
)

type MergeKSortedListSolver struct {
	heap.MergeTwoSortedListSolver
}

func (m *MergeKSortedListSolver) MergeKSortedListSolve(lists []*ListNode) *ListNode {
	switch len(lists) {
	case 1:
		return lists[0]
	case 2:
		return m.MergeTwoSortedListsSolve(lists[0], lists[1])
	default:
		var listArr1, listArr2 []*ListNode
		for i := 0; i < len(lists); i++ {
			if i%2 == 0 {
				listArr1 = append(listArr1, lists[i])
			} else {
				listArr2 = append(listArr2, lists[i])
			}
		}
		list1 := m.MergeKSortedListSolve(listArr1)
		list2 := m.MergeKSortedListSolve(listArr2)
		return m.MergeTwoSortedListsSolve(list1, list2)
	}
}
