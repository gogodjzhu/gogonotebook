package backtracking

import (
	. "gogonotebook/common"
)

type ZeroOneKnapsackSolver struct {
	maxValue int
	wArr     []int
	vArr     []int
}

/**
在背包问题中，决策树的裁剪效果跟背包大小max有很大关系
极端情况背包足以装下所有物品，那么复杂度将变为o(n!)
*/
func (zs *ZeroOneKnapsackSolver) ZeroOneKnapsackSolve(wArr, vArr []int, max int) int {
	options := make([]int, 0)
	for i, w := range wArr {
		if w <= max {
			options = append(options, i)
		}
	}
	zs.wArr = wArr
	zs.vArr = vArr
	zs.backtracking([]int{}, options, max)
	return zs.maxValue
}

// tracks是被选择的物品的下标，其对应的重量 和 价格保存在 wArr/vArr
// options是当前决策允许的选择
// max最大允许的重量
func (zs *ZeroOneKnapsackSolver) backtracking(tracks []int, options []int, max int) {
	if len(options) <= 0 {
		// 满足结束条件
		sumValue := 0
		for _, index := range tracks {
			sumValue += zs.vArr[index]
		}
		if sumValue > zs.maxValue {
			zs.maxValue = sumValue
		}
		return
	}

	// 注意golang range的语法，
	// 第一个i是自增下标[0,1,...]，第二个idx是数组的值，在这里options保存的是wArr的下标
	for i, idx := range options {
		/*计算新的选择列表*/
		tracks = append(tracks, idx)
		trackWeightSum := 0
		for _, trackIdx := range tracks {
			trackWeightSum += zs.wArr[trackIdx]
		}
		var nextOptions []int // 新的选择列表
		for _, option := range DeleteIntByIndex(options, i) {
			if trackWeightSum+zs.wArr[option] <= max {
				nextOptions = append(nextOptions, option)
			}
		}

		// 进入下一轮决策
		zs.backtracking(tracks, nextOptions, max)
		// 撤销选择
		tracks = DeleteIntByIndex(tracks, len(tracks)-1)
	}
}
