package backtracking

import (
	"fmt"
	. "gogonotebook/common"
)

type ThreeNumberSumSolve struct {
	res     [][]int
	hashSet map[string]bool
}

func (solver *ThreeNumberSumSolve) ThreeNumberSumSolve(nums []int) [][]int {
	solver.res = make([][]int, 0)
	solver.hashSet = make(map[string]bool)

	nums = MergeSort(nums)
	solver.backtracking([]int{}, nums)
	return solver.res
}

func (solver *ThreeNumberSumSolve) backtracking(tracks, options []int) {
	if len(tracks) == 3 && SumInt(tracks) == 0 {
		tracks = MergeSort(tracks)
		key := fmt.Sprintf("%+v", tracks)
		if _, ok := solver.hashSet[key]; !ok {
			solver.res = append(solver.res, tracks)
			solver.hashSet[key] = true
		}
		return
	}

	for i := 0; i < len(options); i++ {
		// 做选择
		if len(tracks) >= 3 {
			continue
		}
		if SumInt(tracks) > 0 && options[i] >= 0 {
			continue
		}
		tracks := append(tracks, options[i])
		// 下一层决策
		solver.backtracking(tracks,
			DeleteIntByIndex(options, i))
		// 撤销选择
		tracks = DeleteIntByIndex(tracks, len(tracks)-1)
	}
}
