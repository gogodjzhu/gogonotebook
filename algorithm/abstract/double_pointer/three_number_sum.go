package double_pointer

import (
	. "gogonotebook/common"
)

/**
双指针解法，核心思路是在数组排序之后，大的数永远在小的数右边，双指针实现不重复选数
*/
type ThreeNumberSumSolve struct {
}

func (ThreeNumberSumSolve) ThreeNumberSumSolve(nums []int) [][]int {
	nums = MergeSort(nums)

	resultArr := make([][]int, 0)

	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		s := i + 1
		e := len(nums) - 1
		for e > s {
			for s+1 < e && nums[s] == nums[s+1] {
				s += 1
			}
			for e-1 > s && nums[e] == nums[e-1] {
				e -= 1
			}
			sum := nums[i] + nums[s] + nums[e]
			switch {
			case sum < 0:
				s += 1
			case sum > 0:
				e -= 1
			case sum == 0:
				arr := []int{nums[i], nums[s], nums[e]}
				resultArr = append(resultArr, arr)
				s += 1
				e -= 1
			}
		}
	}
	return resultArr
}
