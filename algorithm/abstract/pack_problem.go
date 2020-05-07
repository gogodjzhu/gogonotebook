package abstract

type PackProblem interface {

	/* 01背包问题 */

	// 纯粹寻找最大容量
	// 已知背包最大容量max, 有物品重weights, 每件物品数量均只有一个, 求最多可放下多重的物品
	ZeroOnePack01(max int, weights []int) int
	// 寻找最大价值
	// 已知背包最大容量max, 有物品重weights, 价值values, 每件物品数量均只有一个, 求最多可放下多少价值的物品
	ZeroOnePack02(max int, weights []int, values []int) int
	// 寻找最大价值, 且必须放满
	// 已知背包最大容量max, 有物品重weights, 价值values, 每件物品数量均只有一个, 求放满时最多可放下多少价值的物品
	ZeroOnePack03(max int, weights []int, values []int) int

	/* 完全背包问题 */

	// 寻找最大价值
	// 已知背包最大容量max, 有一批物品重weights, 价值values, 每件物品数量无限, 求最多可放下多少价值的物品
	CompletePack01(max int, weights []int, values []int) int
	// 寻找最大价值, 且必须放满
	// 已知背包最大容量max, 有一批物品重weights, 价值values, 每件物品数量无限, 求放满时最多可放下多少价值的物品. 如果无法放满返回-1
	CompletePack02(max int, weights []int, values []int) int
	// 寻找方案数, 必须放满
	// 已知背包最大容量max, 有一批物品重weights, 价值values, 每件物品数量无限, 求放满背包有多少种方案
	CompletePack03(max int, weights []int) int

	/* 多重背包问题 */

	// 寻找最大值
	MultiplePack01(max int, weights []int, counts []int, values []int) int
	// 寻找最大值,且必须放满
	MultiplePack02(max int, weights []int, counts []int, values []int) int
	// 寻找方案数,且必须放满
	MultiplePack03(max int, weights []int, counts []int, values []int) int
}
