package graph

import "math"

type NetworkDelayTimeSolver struct {
}

func (n NetworkDelayTimeSolver) NetworkDelayTimeSolve(times [][]int, N int, K int) int {
	linkTable := make(map[int]map[int]int) // {nodeId -> {nextId1->nextLen1, nextId2->nextLen2}}
	for _, arr := range times {
		if _, ok := linkTable[arr[0]-1]; !ok {
			linkTable[arr[0]-1] = make(map[int]int)
		}
		linkTable[arr[0]-1][arr[1]-1] = arr[2]
	}
	return n.dijkstraSolve(linkTable, N, K-1)
}

/**
{nodeId -> {nextId1->nextLen1, nextId2->nextLen2}}
*/
func (NetworkDelayTimeSolver) dijkstraSolve(graph map[int]map[int]int, num int, start int) int {
	var Dis = make([]int, num)   // Dis保存从start到任意节点的最短距离(当前最短)
	var seen = make([]bool, num) // seen保存Dis中敲定为最短距离的点(松弛结束)
	for i := 0; i < num; i++ {
		Dis[i] = math.MaxInt32
		seen[i] = false
	}
	Dis[start] = 0 // <0>表示start到当前节点[start],距离为<0>

	for {
		canNode := -1
		canDist := math.MaxInt32
		for i := 0; i < num; i++ {
			if !seen[i] && Dis[i] < canDist {
				canDist = Dis[i]
				canNode = i
			}
		}

		if canNode < 0 {
			break
		}
		seen[canNode] = true
		if nextArr, ok := graph[canNode]; ok {
			for nextNode, nextDis := range nextArr {
				if Dis[nextNode] > Dis[canNode]+nextDis {
					Dis[nextNode] = Dis[canNode] + nextDis
				}
			}
		}
	}

	// 遍历从start到所有点的最短距离,取最大者为全局到达时间
	max := -1
	for _, dis := range Dis {
		if dis == math.MaxInt32 {
			return -1
		} else {
			if max < dis {
				max = dis
			}
		}
	}
	return max
}
