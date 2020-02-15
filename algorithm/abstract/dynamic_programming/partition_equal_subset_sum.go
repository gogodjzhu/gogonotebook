package dynamic_programming

import . "gogonotebook/common"

type PartitionEqualSubsetSumSolver struct {
}

func (PartitionEqualSubsetSumSolver) PartitionEqualSubsetSumSolve(nums []int) bool {
	total := SumInt(nums)
	if total%2 != 0 {
		return false
	}
	target := total / 2
	dpSolver := PackProblemSolver{}
	max := dpSolver.ZeroOnePack01(target, nums)
	return max == target
}
