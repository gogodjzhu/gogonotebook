package dynamic_programming

import (
	. "gogonotebook/common"
	"math"
)

type PackProblemSolver struct {
}

func (PackProblemSolver) ZeroOnePack01(V int, W []int) int {
	return zeroOnePack01Beta(V, W)
}

// 最简单的01背包问题，寻求背包能装下的最大重量
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

// 上一中解法在二维数组中保存了所有的中间结果, 而实际上我们并不需要这些中间结果, 可以优化为一维数组解决
func zeroOnePack01Beta(V int, W []int) int {
	F := make([]int, V+1)          // F[0..V] <- 0
	for i := 1; i <= len(W); i++ { // i <- 1 to N
		Fi := make([]int, V+1) // F[i,v]
		copy(Fi, F)
		for v := W[i-1]; v <= V; v++ { // v <- Wi to V
			leaveIt := F[v]                // F[v]
			takeIt := F[v-W[i-1]] + W[i-1] // F[v-Wi] + Wi
			if takeIt > leaveIt {
				Fi[v] = takeIt // max{F[v], F[v-Wi] + Wi}
			} else {
				Fi[v] = leaveIt
			}
		}
		F = Fi // F[V] update
	}
	return F[V]
}

// 典型的01背包问题
func (PackProblemSolver) ZeroOnePack02(V int, W []int, P []int) int {
	F := make([][]int, len(W)+1)
	F[0] = make([]int, V+1)        // F[0, 0..V] <- 0
	for i := 1; i <= len(W); i++ { // i <- 1 to N
		Fi_1 := F[i-1]
		Fi := make([]int, V+1) // F[i,v]
		copy(Fi, Fi_1)
		for v := W[i-1]; v <= V; v++ { // v <- Wi to V
			leaveIt := Fi_1[v]                // F[i-1, v]
			takeIt := Fi_1[v-W[i-1]] + P[i-1] // F[i-1, v-Wi] + Wi
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

// 满包问题，注意初始值
func (PackProblemSolver) ZeroOnePack03(V int, W []int, P []int) int {
	F := make([]int, V+1)
	// init
	F[0] = 0
	for i := 1; i <= V; i++ {
		F[i] = math.MinInt32
	}
	for i := 1; i <= len(W); i++ { // i <- 1 to N
		Fi := make([]int, V+1)
		copy(Fi, F)
		for v := W[i-1]; v <= V; v++ { // v <- Wi to V
			takeIt := F[v-W[i-1]] + P[i-1]                        // F[i-1,v-Wi] + Pi
			leaveIt := F[v]                                       // F[i-1, V]
			if F[v-W[i-1]] != math.MinInt32 && takeIt > leaveIt { // if F[v-W[v-1]] == Min then F[v-W[W-1]] it's unreachable
				Fi[v] = takeIt
			} else {
				Fi[v] = leaveIt
			}
		}
		F = Fi
	}
	if F[V] == math.MinInt32 {
		return 0
	} else {
		return F[V]
	}
}

/**
完全背包问题-寻找最大价值
已知背包最大容量max, 有一批物品重weights, 价值values, 每件物品数量无限, 求最多可放下多少价值的物品
*/
func (PackProblemSolver) CompletePack01(V int, W []int, P []int) int {
	F := make([][]int, len(W)+1)
	F[0] = make([]int, V+1)
	for i := 1; i <= len(W); i++ {
		Fi_1 := F[i-1]
		Fi := make([]int, V+1)
		copy(Fi, Fi_1)
		for v := W[i-1]; v <= V; v++ {
			Fi[v] = Max(F[i-1][v], Fi[v-W[i-1]]+P[i-1])
		}
		F[i] = Fi
	}
	return F[len(W)][V]
}

// 完全背包问题-寻找最大价值, 且必须放满
// 已知背包最大容量max, 有一批物品重weights, 价值values, 每件物品数量无限, 求放满时最多可放下多少价值的物品
// 无法放满返回-1
// leetcode.322
func (PackProblemSolver) CompletePack02(V int, W []int, P []int) int {
	F := make([]int, V+1)
	for i := 1; i < len(F); i++ {
		F[i] = -1
	}
	for i := 1; i <= len(W); i++ {
		for j := W[i-1]; j <= V; j++ {
			if F[j-W[i-1]] != -1 {
				F[j] = Max(F[j], F[j-W[i-1]]+P[i-1])
			}
		}
	}
	return F[V]
}

// 完全背包问题-寻找最大价值, 且必须放满
// 已知背包最大容量V, 有一批物品重W, 每件物品数量无限, 求放满背包有多少种方案
// leetcode.518
func (PackProblemSolver) CompletePack03(V int, W []int) int {
	F := make([]int, V+1)
	F[0] = 1
	for i := 1; i <= len(W); i++ {
		for j := W[i-1]; j <= V; j++ {
			F[j] = F[j-W[i-1]] + F[j]
		}
	}
	return F[V]
}

// 寻找最大值
func (PackProblemSolver) MultiplePack01(max int, weights []int, counts []int, values []int) int {
	F := make([]int, max+1)
	for i := 1; i <= len(weights); i++ {
		Fi := make([]int, max+1)
		copy(Fi, F)
		for v := weights[i-1]; v <= max; v++ {
			for c := 0; c <= counts[i-1] && v-weights[i-1]*c >= 0; c++ {
				Fi[v] = Max(Fi[v], F[v-weights[i-1]*c]+values[i-1]*c)
			}
		}
		F = Fi
	}
	return F[len(F)-1]
}

func (PackProblemSolver) MultiplePack02(max int, weights []int, counts []int, values []int) int {
	// TODO
	panic("implement me")
}

func (PackProblemSolver) MultiplePack03(max int, weights []int, counts []int, values []int) int {
	// TODO
	panic("implement me")
}
