package dynamic_programming

import (
	"math"
)

type ZeroOneKnapsackSolver struct {
}

func (zs ZeroOneKnapsackSolver) ZeroOneKnapsackSolve(wArr, vArr []int, max int) int {
	if len(wArr) != len(vArr) {
		panic("wArr doesn't equal to vArr")
	}
	if len(wArr) < 1 {
		return 0
	}
	cache := make([][]int, len(wArr))

	// 初始化第一件物品
	maxValueWithFirst := make([]int, max+1)
	for w := 0; w <= max; w++ {
		// 遍历所有可能的重量
		if w >= wArr[0] {
			// 因为只有第一件物品可以供选择,达到重量要求时, 放入第一件物品就是最佳的选择
			maxValueWithFirst[w] = vArr[0]
		}
		// else default to zero
	}
	cache[0] = maxValueWithFirst
	maxWeight := math.MinInt32

	// 注意起始下标是1, 跳过了物品i=0
	for i := 1; i < len(wArr); i++ {
		// 遍历所有物品加入选择的情况
		for w := 0; w <= max; w++ {
			if w < wArr[i] {
				cache[i] = append(cache[i], cache[i-1][w])
				continue // 放不下物品i
			}
			withoutThis := cache[i-1][w] // 不放物品i的最优解

			// 由于物品i的重量是wArr[i], 减去物品i之后的重量剩余w-wArr[i],
			// 而除了物品i之外, 在w-wArr[i]重量下剩余物品的组成的最优解为cache[i-1][w-wArr[i]]
			withThis := cache[i-1][w-wArr[i]] + vArr[i] // 放下物品i的最优解

			if withThis > withoutThis {
				cache[i] = append(cache[i], withThis)
			} else {
				cache[i] = append(cache[i], withoutThis)
			}
			if maxWeight < cache[i][w] {
				maxWeight = cache[i][w]
			}
		}
	}
	return maxWeight
}
