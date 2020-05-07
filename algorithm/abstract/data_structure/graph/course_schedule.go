package graph

import . "gogonotebook/common"

type CourseScheduleSolver struct {
}

// 3, [[1,0][1,2],[2,0]]
// [[], [0,2], [0]]
func (c CourseScheduleSolver) CourseScheduleSolve(n int, p [][]int) bool {
	deptGraph := make(map[int][]int, 0)
	for i := 0; i < n; i++ {
		deptGraph[i] = []int{}
	}
	for _, arr := range p {
		deptGraph[arr[0]] = append(deptGraph[arr[0]], arr[1])
	}
	var tracks []int
	visited := make(map[int]bool)
	for i := 0; i < n; i++ {
		if len(deptGraph[i]) == 0 {
			tracks = append(tracks, i)
			visited[i] = true
		}
	}
	for len(tracks) > 0 {
		n -= 1
		idx := tracks[0]
		tracks = tracks[1:]
		for i := range deptGraph {
			arr := DeleteItem(deptGraph[i], idx)
			_, ok := visited[i]
			if len(arr) == 0 && !ok {
				tracks = append(tracks, i)
				visited[i] = true
			}
			deptGraph[i] = arr
		}
	}

	return n == 0
}
