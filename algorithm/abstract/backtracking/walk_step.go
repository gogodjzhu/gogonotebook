package backtracking

import . "gogonotebook/common"

/**

 */
type WalkStepSolver struct {
	res [][]int
}

func (bs *WalkStepSolver) SolveWalkStep(total int) int {
	track := make([]int, 0)
	bs.res = [][]int{}
	switch {
	case total < 1:
		return 0
	case total == 1:
		options := []int{1}
		bs.backtracking(track, options, total)
	default:
		options := []int{1, 2}
		bs.backtracking(track, options, total)
	}
	return len(bs.res)
}

func (bs *WalkStepSolver) backtracking(track []int, options []int, n int) {
	if len(options) <= 0 {
		// 满足结束条件
		bs.res = append(bs.res, track)
		return
	}
	for _, opt := range options {
		// 执行选择，更新路径
		track := append(track, opt)
		left := n - SumInt(track)

		// 根据已做的选择，更新可选择列表
		switch {
		case left < 1:
			bs.backtracking(track, []int{}, n)
		case left == 1:
			bs.backtracking(track, []int{1}, n)
		default:
			bs.backtracking(track, []int{1, 2}, n)
		}

		// 撤销选择, 恢复路径
		track = track[:len(track)-1]
	}
}
