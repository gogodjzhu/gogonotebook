package dynamic_programming

type ClimbingStairsSolver struct {
}

func (ClimbingStairsSolver) ClimbingStairsSolve(total int) int {
	if total < 1 {
		return 0
	}
	if total == 1 {
		return 1
	}

	// 到达i+1级台阶的路径数量
	stair2TypeOfWay := make([]int, total)
	stair2TypeOfWay[0] = 1 // 上到0+1级台阶的方法共有一种: [1]
	stair2TypeOfWay[1] = 2 // 上到1+1级台阶的方法共有两种: [1,1]/[2]

	for i := 2; i < total; i++ {
		// 到达i级台阶的方法共有 [从i-1级台阶走一级] + [从i-2级台阶走两级]
		stair2TypeOfWay[i] = stair2TypeOfWay[i-1] + stair2TypeOfWay[i-2]
	}
	return stair2TypeOfWay[total-1]
}
