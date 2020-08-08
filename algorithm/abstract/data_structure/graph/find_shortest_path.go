package graph

import (
	. "gogonotebook/common"
	"math"
)

/**
                 Negative-Edge-Weight     Positive-Edge-Weight    Cyclic      Runtime
DFS                     NO                         NO               NO        O(n+e)
BFS                     NO                         NO               YES       O(n+e) or O(g^d)
DijkstraConstantWeight  NO                         YES              YES       O(e+n*log(n))
BellmanFord             YES                        YES              YES       O(n*e)

n = number of nodes, e = number of edges, g = largest number of adjacent nodes for any node, d = length of the shortest path
*/
type FindShortestPathSolver struct {
}

/**
核心思路是从start开始,递归在栈中推入后继节点,直到后继节点为end,记下此时栈的高度即为当前路径的距离.
再依次出栈
PS.注意DFS只能解决无环的场景
*/
func (f FindShortestPathSolver) DFS(graph map[int][]int, num, start, end int) []int {
	// 由于DFS不能解决有环的图, 亦即只能解决树的问题, 且只有针对二叉树比较好解, 若为多叉又更麻烦
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
			last := arr[len(arr)-1]
			if last == end {
				return arr
			}
			nextArr := graph[last]
			if nextArr != nil {
				for _, next := range nextArr {
					if Contains(arr, next) {
						continue
					} else {
						nextGraph = append(nextGraph, append(arr, next))
					}
				}
			}
		}
		if len(nextGraph) == 0 {
			return []int{}
		}
		visited = nextGraph
	}
}

/**
对比此方法与 DijkstraDiffWeight 方法的区别以深入理解
*/
func (f FindShortestPathSolver) DijkstraConstantWeight(graph map[int][]int, num, start, end int) []int {
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
			// PS. 有点类似预BFS的概念, 不同在于BFS未保存松弛状态shortestPrev, 导致其只能计算单源单目的最短路径,
			// 而这里可以利用shortestPrev计算源到任意节点的最短路径
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

/**
经典Dijkstra算法在单源最短路径问题的应用
核心在对'松弛(relaxation)'概念的理解, 通过在每一轮操作中取出距离start最近的节点,以之为步进源点松弛所有邻接点,
每轮松弛之后都可能导致部分节点跟start的距离变近. 而因为被松弛的节点是以步进源点为前驱的, 所以每一轮的步进源点都
可以被标记为最靠近start节点的状态, 进而每轮松弛都可以确定一个节点的最短距离, 使得算法收敛
*/
func (f FindShortestPathSolver) DijkstraDiffWeight(graph map[int]map[int]int, num, start, end int) []int {
	var Dis = make([]int, num)  // 保存从[start]到[每个节点]的最短距离, max表示不可达
	var Prev = make([]int, num) // 保存从[start]到[每个节点]的最短前驱节点, -1表示无前驱
	var visited = make([]bool, num)
	for i := 0; i < len(Dis); i++ {
		Dis[i] = math.MaxInt32
		Prev[i] = -1
	}
	Dis[start] = 0   // 0表示[start]到[当前节点start],距离为0
	Prev[start] = -1 // -1表示[当前节点start]无前驱(本身就是起点)
	for {            // 每次循环为一轮松弛
		visitedIdx := -1
		visitedLen := math.MaxInt32
		for i := 0; i < num; i++ {
			// 遍历所有未visited的节点, 获取Dis[i]最短距离的作为当前visited节点(即此轮松弛的起点)
			if !visited[i] && visitedLen > Dis[i] {
				visitedIdx = i
				visitedLen = Dis[i]
			}
		}
		// 无当前visited节点, 即所有未visited的节点(如果有的话)长度均为maxInt32,
		// 即start无法到达任何剩余的(!visited[i])节点, 无需继续松弛, 退出循环
		if visitedIdx == -1 {
			break
		}
		visited[visitedIdx] = true
		// visited 节点此时是所有!visited节点中最靠近start的一个节点, 那么从此节点出发到达其余节点可能松弛邻接节点的距离,
		// 我们逐个判断并更新Dis, 这就是松弛本身
		if nextMap, ok := graph[visitedIdx]; ok {
			for nextIdx, nextLen := range nextMap {
				if Dis[nextIdx] > Dis[visitedIdx]+nextLen {
					// 更新最短距离
					Dis[nextIdx] = Dis[visitedIdx] + nextLen
					// 更新最短前驱
					Prev[nextIdx] = visitedIdx
				}
			}
		}
	}
	// 根据end节点进行回溯,得到最短路径
	stepPath := []int{end}
	for Prev[end] != -1 {
		stepPath = append(stepPath, Prev[end])
		end = Prev[end]
	}
	if stepPath[len(stepPath)-1] != start {
		return nil
	}
	return ReverseArr(stepPath)
}

func (f FindShortestPathSolver) BellmanFord(graph map[int][]int, num, start, end int) []int {
	// TODO
	panic("implement me")
}
