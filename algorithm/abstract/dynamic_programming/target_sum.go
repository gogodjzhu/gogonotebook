package dynamic_programming

import . "gogonotebook/common"

type TargetSumSolver struct {
}

func (TargetSumSolver) TargetSumSolve(nums []int, t int) int {
	if Abs(SumInt(nums)) < t {
		return 0
	}
	maxNumsLen := 10
	F := make([]int, maxNumsLen*2+1)
	F[maxNumsLen] = 1
	for i := 1; i <= len(nums); i++ {
		Fi := make([]int, len(F))
		for j := 0; j <= maxNumsLen*2; j++ {
			if j-nums[i-1] >= 0 {
				Fi[j] += F[j-nums[i-1]]
			}
			if j+nums[i-1] < len(F) {
				Fi[j] += F[j+nums[i-1]]
			}
		}
		F = Fi
	}
	return F[maxNumsLen+t]
}
