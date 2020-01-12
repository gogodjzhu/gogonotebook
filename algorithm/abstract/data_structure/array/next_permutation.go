package array

import . "gogonotebook/common"

type NextPermutationSolver struct {
}

// 1,3,7,2,5
//
func (n NextPermutationSolver) NextPermutationSolve(nums []int) {
	if len(nums) <= 1 {
		return
	}
	lastInc := -1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			lastInc = i - 1
		}
	}
	if lastInc == -1 {
		HeapSort(nums)
		return
	}
	minGreat := lastInc + 1
	for i := minGreat; i < len(nums); i++ {
		if nums[i] > nums[lastInc] && nums[i] < nums[minGreat] {
			minGreat = i
		}
	}
	Swap(nums, lastInc, minGreat)
	HeapSort(nums[lastInc+1:])
}
