package common

type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}

// [[from1, to1][from1, to2]]
// 转化为:
// [from1][to1,to2]
func Arr2GraphList(n int, arr [][]int) [][]int {
	lists := make([][]int, n)
	for i := 0; i < n; i++ {
		lists[i] = make([]int, 0)
	}
	for _, a := range arr {
		from := a[0]
		to := a[1]
		lists[from] = append(lists[from], to)
	}
	return lists
}

func Clone(g *GraphNode, visited map[*GraphNode]*GraphNode) *GraphNode {
	if g == nil {
		return nil
	}
	if len(g.Neighbors) == 0 {
		return &GraphNode{
			Val: g.Val,
		}
	}
	n := &GraphNode{
		Val:       g.Val,
		Neighbors: make([]*GraphNode, len(g.Neighbors)),
	}
	visited[g] = n
	for i := 0; i < len(g.Neighbors); i++ {
		t := g.Neighbors[i]
		if v, ok := visited[t]; ok {
			n.Neighbors[i] = v
		} else {
			n.Neighbors[i] = Clone(t, visited)
		}
	}
	return n
}
