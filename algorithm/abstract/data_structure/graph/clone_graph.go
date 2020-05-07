package graph

import . "gogonotebook/common"

type CloneGraphSolver struct {
}

func (c CloneGraphSolver) CloneGraphSolve(n *GraphNode) *GraphNode {
	visited := make(map[*GraphNode]*GraphNode)
	return Clone(n, visited)
}
