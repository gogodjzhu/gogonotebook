package dynamic_programming

type BestTimeToBuyAndSellStockSolver struct {
}

/**
状态转移方程:
第i天卖出的最优利润 = 第i天当天利润 > 0
				 ? 第i-1天卖出的最优利润 + 第i天当天利润
				 : 第i-1天卖出的最优利润 > -第i天当天利润
									  ? 第i-1天卖出的最优利润 + 第i天当天利润
									  : 0
*/
func (BestTimeToBuyAndSellStockSolver) BestTimeToBuyAndSellStockSolve(prices []int) int {
	yesterdayProfit := 0
	best := 0

	for i := 1; i < len(prices); i++ {
		todayProfit := prices[i] - prices[i-1]
		if todayProfit > 0 {
			todayProfit += yesterdayProfit
		} else {
			if yesterdayProfit > -todayProfit {
				todayProfit += yesterdayProfit
			} else {
				todayProfit = 0
			}
		}
		if todayProfit > best {
			best = todayProfit
		}
		yesterdayProfit = todayProfit
	}
	return best
}
