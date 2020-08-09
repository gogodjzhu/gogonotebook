package dynamic_programming

import (
	. "gogonotebook/common"
	"math"
)

type MaximumProductSubarraySolver struct {
}

func (m MaximumProductSubarraySolver) MaximumProductSubarraySolve(nums []int) int {
	max, iMax, iMin := math.MinInt32, 1, 1
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			tmp := iMax
			iMax = iMin
			iMin = tmp
		}
		iMax = Max(iMax*nums[i], nums[i])
		iMin = Min(iMin*nums[i], nums[i])

		max = Max(max, iMax)
	}
	return max
}
