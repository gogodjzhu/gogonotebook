package dynamic_programming

import (
	. "gogonotebook/common"
)

type LongestIncreasingSubsequence struct {
}

/**
使用dp求解,动态转移方程为:
dp[i] = max(dp[j]) + 1, 其中j∈[0,i) 且 nums[j] < nums[i]
*/
func (l LongestIncreasingSubsequence) LongestIncreasingSubsequenceSolve(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	for idx := 0; idx < len(nums); idx++ {
		var greatest = 0
		for i := 0; i < idx; i++ {
			if nums[i] < nums[idx] && dp[i] > greatest {
				greatest = dp[i]
			}
		}
		if greatest > 0 {
			dp[idx] = greatest + 1
		} else {
			dp[idx] = 1
		}
	}
	return MaxInt(dp)
}
