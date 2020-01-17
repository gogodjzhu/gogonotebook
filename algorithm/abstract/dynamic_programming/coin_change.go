package dynamic_programming

type CoinChangeSolver struct {
}

/**
状态转移方程:
f(n) = min{f(n-coins[i])}+1

基于自底向上的思路实现动态规划解法
*/
func (ccs CoinChangeSolver) CoinChangeSolve(coins []int, amount int) int {
	// 备忘录, mem[n]保存实现总额为n的最少零钱数, 若无法实现记为-1
	mem := []int{0}
	for cur := 1; cur <= amount; cur++ {
		// 先假设无法实现amount的总额组合
		mem = append(mem, -1)
		// 遍历所有的coins寻找最小值, 注意不能单纯使用最大面额的满足条件的coin作为最优解(否则会陷入贪心算法陷阱)
		for _, coin := range coins {
			// 通过备忘录寻找计算结果缓存
			if cur-coin >= 0 && cur-coin < len(mem) && mem[cur-coin] != -1 {
				// 如果当前是第一个实现cur总额的组合, 或者 新组合比已有组合耗费的金币少
				if mem[cur] == -1 || mem[cur] > mem[cur-coin]+1 {
					// 更新实现cur总额的最优结果
					mem[cur] = mem[cur-coin] + 1
				}
			}
			// 在备忘录录中找不到合法的中间结果, 跳过即可, 因为前面已经假设无法完成(-1)
		}
	}
	return mem[amount]
}
