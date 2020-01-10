package backtracking

import . "gogonotebook/common"

type PermutationsSolver struct {
	result [][]int
}

func (bs *PermutationsSolver) PermutationsSolve(nums []int) [][]int {
	bs.backTracking(make([]int, 0), nums)
	return bs.result
}

func (bs *PermutationsSolver) backTracking(tracks []int, options []int) {
	if len(options) == 0 {
		// 满足结束条件
		bs.result = append(bs.result, tracks)
		return
	}

	for i := 0; i < len(options); i++ {
		// 做选择
		tracks = append(tracks, options[i])

		// 更新(减小)选择列表，并进入下一层决策
		bs.backTracking(tracks, DeleteIntByIndex(options, i))
		// 撤销当前选择
		tracks = DeleteIntByIndex(tracks, len(tracks)-1)
		// tracks = tracks[:len(tracks) - 1] // TODO 这种写法当nums=[1,2,3,4]时会出错，why?
	}
}
