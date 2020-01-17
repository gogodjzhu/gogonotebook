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
	max := zeroOnePack01(target, nums)
	return max == target
}

func zeroOnePack01(V int, W []int) int {
	F := make([][]int, len(W)+1)
	F[0] = make([]int, V+1)        // F[0, 0..V] <- 0
	for i := 1; i <= len(W); i++ { // i <- 1 to N
		Fi_1 := F[i-1]
		Fi := make([]int, V+1) // F[i,v]
		copy(Fi, Fi_1)
		for v := W[i-1]; v <= V; v++ { // v <- Wi to V
			leaveIt := Fi_1[v]                // F[i-1, v]
			takeIt := Fi_1[v-W[i-1]] + W[i-1] // F[i-1, v-Wi] + Wi
			if takeIt > leaveIt {
				Fi[v] = takeIt // max{F[i-1, v], F[i-1, v-Wi] + Wi}
			} else {
				Fi[v] = leaveIt
			}
		}
		F[i] = Fi // F[i, V]
	}
	return F[len(W)][V]
}
