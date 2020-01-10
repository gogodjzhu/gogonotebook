package backtracking

import . "gogonotebook/common"

/**

 */
type ClimbingStairsSolver struct {
	res [][]int
}

func (css *ClimbingStairsSolver) ClimbingStairsSolve(total int) int {
	track := make([]int, 0)
	css.res = [][]int{}
	switch {
	case total < 1:
		return 0
	case total == 1:
		options := []int{1}
		css.backtracking(track, options, total)
	default:
		options := []int{1, 2}
		css.backtracking(track, options, total)
	}
	return len(css.res)
}

func (css *ClimbingStairsSolver) backtracking(track []int, options []int, n int) {
	if len(options) <= 0 {
		// 满足结束条件
		css.res = append(css.res, track)
		return
	}
	for _, opt := range options {
		// 执行选择，更新路径
		track := append(track, opt)
		left := n - SumInt(track)

		// 根据已做的选择，更新可选择列表
		switch {
		case left < 1:
			css.backtracking(track, []int{}, n)
		case left == 1:
			css.backtracking(track, []int{1}, n)
		default:
			css.backtracking(track, []int{1, 2}, n)
		}

		// 撤销选择, 恢复路径
		track = track[:len(track)-1]
	}
}
