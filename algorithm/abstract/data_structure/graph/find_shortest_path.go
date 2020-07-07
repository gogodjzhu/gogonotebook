package graph

import (
	. "gogonotebook/common"
	"math"
)

/**
                 Negative-Edge-Weight     Positive-Edge-Weight    Cyclic      Runtime
DFS                     NO                         NO               NO        O(n+e)
BFS                     NO                         NO               YES       O(n+e) or O(g^d)
Dijkstra                NO                         YES              YES       O(e+n*log(n))
BellmanFord             YES                        YES              YES       O(n*e)

n = number of nodes, e = number of edges, g = largest number of adjacent nodes for any node, d = length of the shortest path
*/
type FindShortestPathSolver struct {
}

/**
核心思路是从start开始,递归在栈中推入后继节点,直到后继节点为end,记下此时栈的高度即为当前路径的距离.
再依次出栈
*/
func (f FindShortestPathSolver) DFS(graph map[int][]int, num, start, end int) []int {
	// TODO
	return nil
}

/**
本质是把从start开始,递归遍历所有的后继节点进行+1步进, 并通过visited数组保存所有的路径. 直到:
1. 抵达end节点, 由于所有的路径都是+1步进的，因此第一条抵达end节点的路径即是最短路径(之一)
2. 所有的路径上都无未访问的后继节点, 即不存在start->end的路径
*/
func (f FindShortestPathSolver) BFS(graph map[int][]int, num, start, end int) []int {
	visited := [][]int{{start}}
	for {
		var nextGraph [][]int
		for _, arr := range visited {
			if arr[len(arr)-1] == end {
				return arr
			}
			if nextArr, ok := graph[arr[len(arr)-1]]; !ok {
				continue
			} else {
				for _, next := range nextArr {
					nextGraph = append(nextGraph, append(arr, next))
				}
			}
		}
		if len(nextGraph) == 0 {
			return nil
		}
		visited = nextGraph
	}
}

/**
Dijkstra算法的核心是: 假设当前节点cur的最近后继节点是shortest, 则shortest的最短前驱为cur. 遍历所有的路径,
将每个节点的最短前驱保存到数组shortestPrev之后, 即可以倒推出从start到任意节点end的最短路径
PS.对于本例,graph是无权图,所以cur的最近后继节点是所有的后继节点
*/
func (f FindShortestPathSolver) Dijkstra(graph map[int][]int, num, start, end int) []int {
	var shortestPrev = make([]int, num)
	for i := 0; i < num; i++ {
		shortestPrev[i] = math.MinInt32
	}
	shortestPrev[start] = -1 // -1表示无最短前驱

	nextVisit := []int{start}
	for len(nextVisit) > 0 {
		cur := nextVisit[0]
		nextVisit = nextVisit[1:]
		for _, next := range graph[cur] {
			// 由于本例graph是无权图,因此所有的next只要未访问过,最短前驱均为当前节点cur. 如果graph是有权的,则需要比较所有的next,
			// 只有最短的那个next才确定其最短前驱为当前节点cur
			if shortestPrev[next] == math.MinInt32 {
				shortestPrev[next] = cur
				nextVisit = append(nextVisit, next)
			}
		}
	}

	// 根据end节点进行回溯,得到最短路径
	path := []int{end}
	for shortestPrev[end] != -1 {
		path = append(path, shortestPrev[end])
		end = shortestPrev[end]
	}
	return ReverseArr(path)
}

func (f FindShortestPathSolver) BellmanFord(graph map[int][]int, num, start, end int) []int {
	// TODO
	panic("implement me")
}
