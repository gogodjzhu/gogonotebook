package graph

import . "gogonotebook/common"

type RouteBetweenNodesSolver struct {
}

func (r RouteBetweenNodesSolver) RouteBetweenNodesSolve(n int, graph [][]int, start int, target int) bool {
	graphList := Arr2GraphList(n, graph)
	visited := map[int]bool{start: true}
	tracks := []int{start}
	for len(tracks) > 0 {
		cur := tracks[0]
		tracks = tracks[1:]
		if cur == target {
			return true
		}
		for _, next := range graphList[cur] {
			if _, ok := visited[next]; ok {
				continue
			}
			visited[next] = true
			tracks = append(tracks, next)
		}
	}
	return false
}
