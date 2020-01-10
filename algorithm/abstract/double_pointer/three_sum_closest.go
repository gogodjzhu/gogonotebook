package double_pointer

import (
	"gogonotebook/common"
	"math"
)

type ThreeSumClosestSolver struct {
}

func (t ThreeSumClosestSolver) ThreeSumClosestSolve(nums []int, target int) int {
	nums = common.SortInt(nums)
	closest := nums[0] + nums[1] + nums[2]
	closestDif := math.Abs(float64(closest - target))

	for i := 0; i < len(nums); i++ {
		s := i + 1
		e := len(nums) - 1

		for e > s {
			sum := nums[i] + nums[s] + nums[e]
			abs := math.Abs(float64(sum - target))
			if abs < closestDif {
				closest = sum
				closestDif = abs
			}
			if sum < target {
				s += 1
			} else {
				e -= 1
			}
		}
	}
	return closest
}
