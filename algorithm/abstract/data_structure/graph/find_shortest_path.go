package graph

import "math"

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

func (f FindShortestPathSolver) DFS(graph map[int][]int, num, start, end int) []int {
	// TODO
	return nil
}

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

func (f FindShortestPathSolver) Dijkstra(graph map[int][]int, num, start, end int) []int {
	// TODO
	var shortestPrev = make([]int, num)
	shortestPrev[start] = -1 // -1表示无最短前驱

	nextVisit := []int{start}
	for len(nextVisit) > 0 {
		cur := nextVisit[0]
		nextVisit = nextVisit[1:]
		shortest := math.MaxInt32
		for _, next := range graph[cur] {
			if next < shortest {
				shortest = next
			}
		}
		if shortest == math.MaxInt32 {
			shortestPrev[cur] = -1
			continue
		}
		shortestPrev[shortest] = cur
		for _, next := range graph[shortest] {
			nextVisit = append(nextVisit, next)
		}
	}
	var path []int
	for shortestPrev[end] != -1 {
		path = append(path, shortestPrev[end])
		end = shortestPrev[end]
	}
	return path
}

func (f FindShortestPathSolver) BellmanFord(graph map[int][]int, num, start, end int) []int {
	// TODO
	panic("implement me")
}
