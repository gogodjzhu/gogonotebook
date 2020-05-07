package graph

import . "gogonotebook/common"

type FindTheTownJudge struct {
}

func (f FindTheTownJudge) FindTheTownJudgeSolve(n int, trust [][]int) int {
	for _, a := range trust {
		a[0]--
		a[1]--
	}
	trustGraphList := Arr2GraphList(n, trust)
	for _, arr := range trust {
		Swap(arr, 0, 1)
	}
	trustedGraphList := Arr2GraphList(n, trust)

	var trustNoneArr []int

	for i, trusts := range trustGraphList {
		if len(trusts) == 0 {
			trustNoneArr = append(trustNoneArr, i)
		}
	}
	if len(trustNoneArr) == 0 {
		return -1
	}
	for i, trusted := range trustedGraphList {
		if len(trusted)+1 == n {
			return i + 1
		}
	}
	return -1
}
