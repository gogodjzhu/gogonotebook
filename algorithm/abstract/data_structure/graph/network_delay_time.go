package graph

import "math"

type NetworkDelayTimeSolver struct {
}

func (n NetworkDelayTimeSolver) NetworkDelayTimeSolve(times [][]int, N int, K int) int {
	linkTable := make(map[int][][]int) // nodeId -> [nextId, nextLen]
	for _, arr := range times {
		if _, ok := linkTable[arr[0]]; !ok {
			linkTable[arr[0]] = [][]int{}
		}
		linkTable[arr[0]] = append(linkTable[arr[0]], []int{arr[1], arr[2]})
	}
	return n.dfsSolve(linkTable, N, K)
}

func (NetworkDelayTimeSolver) dfsSolve(graph map[int][][]int, num int, start int) int {
	var shortestLen []int
	for i := 0; i < num; i++ {
		shortestLen = append(shortestLen, math.MaxInt32)
	}
	// TODO
	return 0
}
