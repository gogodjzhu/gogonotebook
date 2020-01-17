package dynamic_programming

type CoinChangeSolver2 struct {
}

/**
状态转移方程为：
dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]

dp[i][j]代表仅利用前i种钱币能组成和为j的情况数
dp[i-1][j]代表不使用第i中钱币能组成和为j的情况数
dp[i][j-coins[i-1]]代表使用前i种钱币能组成j-coins[i-1](即j-第i种金币的面额)的情况数

1,2,3
5

1,0,0,0,0,0
1,1,1,1,1,1
1,1,2,2,3,3
1,1,2,3,4,5

*/
func (ccs CoinChangeSolver2) CoinChange2Solve(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 1; i <= len(coins); i++ {
		dpi := make([]int, amount+1)
		copy(dpi, dp)
		for j := coins[i-1]; j <= amount; j++ {
			dpi[j] = dpi[j-coins[i-1]] + dp[j]
		}
		dp = dpi
	}
	return dp[amount]
}
