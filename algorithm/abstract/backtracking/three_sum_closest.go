package backtracking

import (
	. "gogonotebook/common"
	"math"
)

type ThreeSumClosestSolver struct {
	closest    int
	closestDif float64
}

func (solver *ThreeSumClosestSolver) ThreeSumClosestSolve(nums []int, target int) int {
	solver.closest = SumInt(nums[0:3])
	solver.closestDif = math.Abs(float64(solver.closest - target))
	solver.backtracking([]int{}, nums, target)
	return solver.closest
}

func (solver *ThreeSumClosestSolver) backtracking(tracks, options []int, target int) {
	if len(options) <= 0 {
		sum := SumInt(tracks)
		newDis := math.Abs(float64(sum - target))
		if newDis < solver.closestDif {
			solver.closest = sum
			solver.closestDif = newDis
		}
		return
	}

	for i := 0; i < len(options); i++ {
		// 做选择
		switch len(tracks) {
		case 0:
			fallthrough
		case 1:
			// 进入下一层决策
			solver.backtracking(append(tracks, options[i]),
				DeleteIntByIndex(options, i), target)
		case 2:
			// 超过三数, 决策结束
			solver.backtracking(append(tracks, options[i]),
				[]int{}, target)
		default:
		}
		// 撤销选择
		//tracks = DeleteIntByIndex(tracks, len(tracks) - 1)
	}
}
