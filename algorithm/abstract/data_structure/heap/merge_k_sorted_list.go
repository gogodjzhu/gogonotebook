package heap

import . "gogonotebook/common"

type MergeKSortedListSolver struct {
}

func (m MergeKSortedListSolver) MergeKSortedListSolve(lists []*ListNode) *ListNode {
	var arr []int
	for _, list := range lists {
		for list != nil {
			arr = append(arr, list.Val)
			list = list.Next
		}
	}
	HeapSort(arr)
	return Arr2List(arr)
}
