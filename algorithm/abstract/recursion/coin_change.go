package recursion

import "math"

type CoinChangeSolver struct {
}

/**
状态转移方程:
f(n) = min{f(n-coins[i])} + 1

使用自顶向下的方式实现递归解法
*/
func (ccs CoinChangeSolver) CoinChangeSolve(coins []int, amount int) int {
	// 朴素递归法重复计算多, 不可用
	// return pureRecursion(coins, amount)
	for i := 0; i <= amount; i++ {
		mem = append(mem, -2)
	}
	min := withMemRecursion(coins, amount)
	return min
}

func pureRecursion(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	min := math.MaxInt32
	for _, coin := range coins {
		if amount-coin == 0 {
			return 1
		}
		ans := pureRecursion(coins, amount-coin)
		if ans != -1 && ans < min {
			min = ans + 1
		}
	}
	if min == math.MaxInt32 {
		return -1
	}
	return min
}

var mem []int

// 2,3
// 6
// 0,
// mem 备忘录, mem[i] = 总额为i需要的最小零钱数，-1 表示无法实现
func withMemRecursion(coins []int, amount int) int {
	// TODO, need to fix
	if amount == 0 {
		return 0
	}
	if mem[amount] != -2 {
		return mem[amount]
	}
	min := math.MaxInt32
	for _, coin := range coins {
		if amount-coin < 0 {
			continue
		}
		ans := withMemRecursion(coins, amount-coin)
		if ans != -1 && ans+1 < min {
			min = ans + 1
		}
	}
	if min == math.MaxInt32 {
		mem[amount] = -1
	} else {
		mem[amount] = min
	}
	return mem[amount]
}
