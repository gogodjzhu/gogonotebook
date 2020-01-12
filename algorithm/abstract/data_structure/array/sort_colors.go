package array

import . "gogonotebook/common"

type SortColorsSolver struct {
}

func (SortColorsSolver) SortColorsSolve(nums []int) []int {
	return MergeSort(nums)
}
