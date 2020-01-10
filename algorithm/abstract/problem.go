package abstract

/**
给定一个固定容量为max的背包，从一堆物品中寻找能装下且总价值最高的组合的总价
*/
type ZeroOneKnapsack interface {
	// wArr物品重量列表
	// vArr物品价格列表
	ZeroOneKnapsackSolve(wArr, vArr []int, max int) int
}
