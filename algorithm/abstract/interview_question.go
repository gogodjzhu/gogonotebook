package abstract

/**
leetcode.面试题04.01.节点间通路
节点间通路。给定有向图，设计一个算法，找出两个节点之间是否存在一条路径。

eg.1.
input: n = 3, graph = [[0, 1], [0, 2], [1, 2], [1, 2]], start = 0, target = 2
output: true

注意: 其中input.n是节点数量,节点下标从0开始
*/
type RouteBetweenNodes interface {
	RouteBetweenNodesSolve(n int, graph [][]int, start int, target int) bool
}
