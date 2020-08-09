package dynamic_programming

import (
	. "gogonotebook/common"
)

type EditDistanceSolver struct {
}

/**
dp[i][j] 代表 word1 到 i 位置转换成 word2 到 j 位置需要最少步数

所以，
当 word1[i] == word2[j]，dp[i][j] = dp[i-1][j-1]；
当 word1[i] != word2[j]，dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1

其中，dp[i-1][j-1] 表示替换操作，dp[i-1][j] 表示删除操作，dp[i][j-1] 表示插入操作。
*/
func (e EditDistanceSolver) EditDistanceSolve(word1, word2 string) int {
	dp := make([][]int, len(word2)+1)
	for i := 0; i <= len(word2); i++ {
		dp[i] = make([]int, len(word1)+1)
		dp[i][0] = i
	}
	for i := 0; i <= len(word1); i++ {
		dp[0][i] = i
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			change, add, del := 0, 0, 0
			if word1[i-1] == word2[j-1] {
				change = dp[j-1][i-1]
			} else {
				change = dp[j-1][i-1] + 1
			}
			add = dp[j][i-1] + 1
			del = dp[j-1][i] + 1
			dp[j][i] = Min(change, add, del)
		}
	}
	return dp[len(word2)][len(word1)]
}
