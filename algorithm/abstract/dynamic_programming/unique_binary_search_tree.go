package dynamic_programming

type UniqueBinarySearchTree struct {
}

/**
令DP[n]=共n个数,可以构成的BST树的种类
又令dp[i]=以第i个数为根,总共可以构成BST树的种类
则
DP[n]=dp[0]+dp[1]+...+dp[n] (1)

又因为以第i个数为根, 则其左侧(小于第i个数的)有i-1个数, 右侧(大于第i个数的)有n-i个数. 则有:
dp[i]=DP[i-1]*DP[n-i] (2)[i>=1]

结合(1)(2)两式, 可以得出:
DP[n]=
DP[0]*DP[n-1] +
DP[1]*DP[n-2] +
DP[2]*DP[n-3] +
... +
DP[n-i]*DP[0]

这是一个卡特兰序列

(注意构成BST树的种类跟数的具体组合无关,只跟数组的长度有关)
*/
func (u UniqueBinarySearchTree) UniqueBinarySearchTreeSolve(n int) int {
	if n <= 0 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		dpi := 0
		for j := 0; j < i; j++ {
			dpi += dp[j] * dp[i-j-1]
		}
		dp[i] = dpi
	}
	return dp[n]
}
