package abstract

import (
	"gogonotebook/algorithm/abstract/data_structure/array"
	"gogonotebook/algorithm/abstract/data_structure/graph"
	"gogonotebook/algorithm/abstract/data_structure/list"
	"gogonotebook/algorithm/abstract/data_structure/tree"
	"gogonotebook/algorithm/abstract/double_pointer"
	"gogonotebook/algorithm/abstract/dynamic_programming"
	"gogonotebook/algorithm/abstract/recursion"
	. "gogonotebook/common"
	"testing"
)

func TestRemoveNthNodeFromEndOfList(t *testing.T) {
	type testCase struct {
		list     *ListNode
		n        int
		expected *ListNode
	}
	testCases := []testCase{
		{Arr2List([]int{1}), 1, nil},
		{Arr2List([]int{1, 2}), 1, Arr2List([]int{1})},
		{Arr2List([]int{1, 2}), 2, Arr2List([]int{2})},
		{Arr2List([]int{1, 2, 3}), 2, Arr2List([]int{1, 3})},
	}

	solver := double_pointer.RemoveNthNodeFromEndOfListSolver{}
	for id, tc := range testCases {
		actual := solver.RemoveNthNodeFromEndOfListSolve(tc.list, tc.n)
		if !IsIdentical(actual, tc.expected) {
			t.Fatal("case#", id, "|n:", tc.n, "|expected:", tc.expected.ToString(), "|actual:", actual.ToString())
		}
	}
}

func TestSwapNodesInPairs(t *testing.T) {
	type testCase struct {
		list     *ListNode
		expected *ListNode
	}
	testCases := []testCase{
		{Arr2List([]int{1}), Arr2List([]int{1})},
		{Arr2List([]int{1, 2}), Arr2List([]int{2, 1})},
		{Arr2List([]int{1, 2, 3}), Arr2List([]int{2, 1, 3})},
		{Arr2List([]int{1, 2, 3, 4}), Arr2List([]int{2, 1, 4, 3})},
		{Arr2List([]int{1, 2, 3, 4, 5}), Arr2List([]int{2, 1, 4, 3, 5})},
	}

	solver := recursion.SwapNodesInPairsSolver{}
	for id, tc := range testCases {
		actual := solver.SwapNodesInPairsSolve(tc.list)
		if !IsIdentical(actual, tc.expected) {
			t.Fatal("case#", id, "|expected:", tc.expected.ToString(), "|actual:", actual.ToString())
		}
	}
}

func TestRotateList(t *testing.T) {
	type testCase struct {
		list     *ListNode
		k        int
		expected *ListNode
	}
	testCases := []testCase{
		{Arr2List([]int{1}), 0, Arr2List([]int{1})},
		{Arr2List([]int{1}), 1, Arr2List([]int{1})},
		{Arr2List([]int{1}), 3, Arr2List([]int{1})},
		{Arr2List([]int{1, 2}), 0, Arr2List([]int{1, 2})},
		{Arr2List([]int{1, 2}), 1, Arr2List([]int{2, 1})},
		{Arr2List([]int{1, 2}), 3, Arr2List([]int{2, 1})},
		{Arr2List([]int{1, 2, 3}), 0, Arr2List([]int{1, 2, 3})},
		{Arr2List([]int{1, 2, 3}), 1, Arr2List([]int{3, 1, 2})},
		{Arr2List([]int{1, 2, 3}), 3, Arr2List([]int{1, 2, 3})},
	}
	solver := list.RotateListSolver{}
	for id, tc := range testCases {
		actual := solver.RotateListSolve(tc.list, tc.k)
		if !IsIdentical(actual, tc.expected) {
			t.Fatal("case#", id, "|expected:", tc.expected.ToString(), "|actual:", actual.ToString())
		}
	}
}

func TestSortColors(t *testing.T) {
	solver := array.SortColorsSolver{}
	nums := []int{2, 0, 2, 1, 1, 0}
	solver.SortColorsSolve(nums)
}

func TestMergeIntervals(t *testing.T) {
	type testCase struct {
		intervals [][]int
		expected  [][]int
	}
	testCases := []testCase{
		{
			[][]int{{2, 6}, {1, 3}, {8, 10}, {15, 18}},
			[][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			[][]int{{1, 4}, {4, 5}},
			[][]int{{1, 5}},
		},
		{
			[][]int{{1, 4}},
			[][]int{{1, 4}},
		},
		{
			[][]int{},
			[][]int{},
		},
	}
	solver := array.MergeIntervalsSolver{}
	for id, tc := range testCases {
		actual := solver.MergeIntervalsSolve(tc.intervals)
		for i, interval := range actual {
			if interval[0] != tc.expected[i][0] || interval[1] != tc.expected[i][1] {
				t.Fatalf("case#%d, expected:%+v, actual:%+v", id, tc.expected, actual)
			}
		}
	}
}

func TestRotateImage(t *testing.T) {
	type textCase struct {
		matrix   [][]int
		expected [][]int
	}
	textCases := []textCase{
		{
			[][]int{
				{1},
			},
			[][]int{
				{1},
			},
		},
		{
			[][]int{
				{1, 2},
				{4, 5},
			},
			[][]int{
				{4, 1},
				{5, 2},
			},
		},
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			[][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
	}
	solver := array.RotateImageSolver{}
	for id, tc := range textCases {
		solver.RotateImageSolve(tc.matrix)
		if !IdenticalMatrix(tc.matrix, tc.expected) {
			t.Fatalf("case#%d, expected%+v, actual:%+v", id, tc.expected, tc.matrix)
		}
	}
}

func TestNextPermutation(t *testing.T) {
	type testCase struct {
		nums     []int
		expected []int
	}
	testCases := []testCase{
		{
			[]int{},
			[]int{},
		},
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{1, 2, 3},
			[]int{1, 3, 2},
		},
		{
			[]int{3, 2, 1},
			[]int{1, 2, 3},
		},
		{
			[]int{3, 3, 2, 1},
			[]int{1, 2, 3, 3},
		},
		{
			[]int{3, 4, 2, 1},
			[]int{4, 1, 2, 3},
		},
		{
			[]int{6, 2, 5, 2, 9, 3, 2, 1},
			[]int{6, 2, 5, 3, 1, 2, 2, 9},
		},
	}
	solver := array.NextPermutationSolver{}
	for id, tc := range testCases {
		solver.NextPermutationSolve(tc.nums)
		if !IdenticalArray(tc.nums, tc.expected) {
			t.Fatalf("case#%d, expected:%+v, actual:%+v", id, tc.expected, tc.nums)
		}
	}
}

func TestPartitionEqualSubsetSum(t *testing.T) {
	type testCase struct {
		nums     []int
		expected bool
	}
	testCases := []testCase{
		{
			[]int{1, 2, 5},
			false,
		},
	}
	solver := dynamic_programming.PartitionEqualSubsetSumSolver{}
	for id, tc := range testCases {
		actual := solver.PartitionEqualSubsetSumSolve(tc.nums)
		if tc.expected != actual {
			t.Fatalf("case#%d, expected: %v, actual: %v", id, tc.expected, actual)
		}
	}
}

func TestTargetSum(t *testing.T) {
	type testCase struct {
		nums     []int
		t        int
		expected int
	}
	testCases := []testCase{
		{
			[]int{1, 1, 1},
			2,
			0,
		},
		{
			[]int{5},
			5,
			1,
		},
	}
	solver := dynamic_programming.TargetSumSolver{}
	for id, tc := range testCases {
		actual := solver.TargetSumSolve(tc.nums, tc.t)
		if actual != tc.expected {
			t.Fatalf("case#%d, expected: %v, actual: %v", id, tc.expected, actual)
		}
	}
}

func TestCourseSchedule(t *testing.T) {
	type testCase struct {
		n        int
		p        [][]int
		expected bool
	}
	testCases := []testCase{
		{
			1,
			[][]int{},
			true,
		},
		{
			2,
			[][]int{},
			true,
		},
		{
			2,
			[][]int{{1, 0}},
			true,
		},
		{
			2,
			[][]int{{1, 0}, {0, 1}},
			false,
		},
		{
			3,
			[][]int{{1, 0}, {1, 2}, {0, 1}},
			false,
		},
		{
			3,
			[][]int{{1, 0}, {1, 2}, {2, 0}},
			true,
		},
		{
			3,
			[][]int{{0, 2}, {1, 2}, {2, 0}},
			false,
		},
		{
			4,
			[][]int{{1, 0}, {1, 2}, {2, 0}, {1, 3}},
			true,
		},
		{
			4,
			[][]int{{1, 0}, {1, 2}, {2, 0}, {2, 1}},
			false,
		},
		{
			4,
			[][]int{{0, 1}, {0, 2}, {1, 3}, {3, 0}},
			false,
		},
	}

	solver := graph.CourseScheduleSolver{}
	for i, tc := range testCases {
		actual := solver.CourseScheduleSolve(tc.n, tc.p)
		if actual != tc.expected {
			t.Fatalf("case#%d, expected:%v, actual:%v", i, tc.expected, actual)
		}
	}
}

func TestLongestIncreasingSubsequence(t *testing.T) {
	type testCase struct {
		nums     []int
		expected int
	}
	testCases := []testCase{
		{
			[]int{},
			0,
		},
		{
			[]int{1},
			1,
		},
		{
			[]int{3, 2, 1},
			1,
		},
		{
			[]int{1, 2, 3},
			3,
		},
		{
			[]int{1, 3, 2, 5},
			3,
		}, {
			[]int{10, 9, 2, 5, 3, 7, 101, 18},
			4,
		},
	}
	solver := dynamic_programming.LongestIncreasingSubsequence{}
	for i, tc := range testCases {
		actual := solver.LongestIncreasingSubsequenceSolve(tc.nums)
		if actual != tc.expected {
			t.Fatalf("case#%d, expected:%d, actual:%d", i, tc.expected, actual)
		}
	}
}

func TestUniqueBinarySearchTree(t *testing.T) {
	type testCase struct {
		n        int
		expected int
	}
	testCases := []testCase{
		{
			0,
			1,
		},
		{
			1,
			1,
		},
		{
			2,
			2,
		},
		{
			3,
			5,
		},
		{
			4,
			14,
		},
	}
	solver := dynamic_programming.UniqueBinarySearchTree{}
	for i, tc := range testCases {
		actual := solver.UniqueBinarySearchTreeSolve(tc.n)
		if actual != tc.expected {
			t.Fatalf("case#%d, expected:%d, actual:%d", i, tc.expected, actual)
		}
	}
}

func TestNetworkDelayTimeSolve(t *testing.T) {
	type testCase struct {
		time     [][]int
		num      int
		start    int
		expected int
	}
	testCases := []testCase{
		{
			[][]int{{2, 1, 2}, {2, 3, 1}, {3, 4, 3}},
			4, 1, -1,
		},
		{
			[][]int{{2, 1, 2}, {2, 3, 1}, {3, 4, 3}},
			4, 2, 4,
		},
		{
			[][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}},
			4, 2, 2,
		},
		{
			[][]int{{1, 3, 1}, {2, 1, 1}, {2, 3, 1}, {3, 4, 1}},
			4, 2, 2,
		},
		{
			[][]int{{1, 3, 0}, {2, 1, 2}, {2, 3, 1}, {3, 4, 3}},
			4, 2, 4,
		},
		{
			[][]int{{4, 2, 76}, {1, 3, 79}, {3, 1, 81}, {4, 3, 30}, {2, 1, 47}, {1, 5, 61}, {1, 4, 99}, {3, 4, 68}, {3, 5, 46}, {4, 1, 6}, {5, 4, 7}, {5, 3, 44}, {4, 5, 19}, {2, 3, 13}, {3, 2, 18}, {1, 2, 0}, {5, 1, 25}, {2, 5, 58}, {2, 4, 77}, {5, 2, 74}},
			5,
			3,
			59,
		},
	}
	solver := graph.NetworkDelayTimeSolver{}
	for i, tc := range testCases {
		actual := solver.NetworkDelayTimeSolve(tc.time, tc.num, tc.start)
		if actual != tc.expected {
			t.Fatalf("case#%d, expected:%d, actual:%d", i, tc.expected, actual)
		}
	}
}

/**
有向无权图最短路径
*/
func TestFindConstantWeightShortestPath(t *testing.T) {
	type testConsLenCase struct {
		graph           map[int][]int
		num, start, end int
		expected        []int
	}
	cyclicConsLenGraph := map[int][]int{0: {1, 2}, 1: {3, 5, 4}, 2: {0, 1}, 3: {4, 1}, 4: {5}, 5: {2}}
	cyclicConsLenGraphCases := []testConsLenCase{
		{
			cyclicConsLenGraph, 6, 0, 5, []int{0, 1, 5},
		},
		{
			cyclicConsLenGraph, 6, 0, 1, []int{0, 1},
		},
		{
			cyclicConsLenGraph, 6, 1, 0, []int{1, 5, 2, 0},
		},
		{
			cyclicConsLenGraph, 6, 4, 1, []int{4, 5, 2, 1},
		},
	}
	solver := graph.FindShortestPathSolver{}
	for id, tc := range cyclicConsLenGraphCases {
		actual := solver.BFS(tc.graph, tc.num, tc.start, tc.end)
		if !IdenticalArray(actual, tc.expected) {
			t.Fatalf("case#%d, expected:%+v, actual%+v", id, tc.expected, actual)
		}
	}
	for id, tc := range cyclicConsLenGraphCases {
		actual := solver.DijkstraConstantWeight(tc.graph, tc.num, tc.start, tc.end)
		if !IdenticalArray(actual, tc.expected) {
			t.Fatalf("case#%d, expected:%+v, actual%+v", id, tc.expected, actual)
		}
	}
}

/**
有向有权图最短路径
*/
func TestFindDiffWeightShortestPath(t *testing.T) {
	// 带权重的图
	type testDiffLenCase struct {
		graph           map[int]map[int]int
		num, start, end int
		expected        []int
	}
	// 无环有权图
	unCyclicDiffLenGraph := map[int]map[int]int{0: {1: 10, 2: 1}, 2: {3: 1}}
	// 有环有权图
	cyclicDiffLenGraph := map[int]map[int]int{0: {1: 10, 2: 1}, 1: {3: 1, 4: 1, 5: 1}, 2: {1: 1, 0: 1}, 3: {1: 1, 4: 1}, 4: {5: 1}, 5: {2: 1}}
	cyclicDiffLenGraphCases := []testDiffLenCase{
		{
			unCyclicDiffLenGraph, 4, 2, 1, nil,
		},
		{
			cyclicDiffLenGraph, 6, 0, 5, []int{0, 2, 1, 5},
		}, {
			cyclicDiffLenGraph, 6, 0, 1, []int{0, 2, 1},
		}, {
			cyclicDiffLenGraph, 6, 1, 0, []int{1, 5, 2, 0},
		}, {
			map[int]map[int]int{
				0: {2: 79, 4: 61, 3: 99, 1: 0},
				1: {0: 47, 2: 13, 4: 58, 3: 77},
				2: {3: 68, 0: 81, 4: 46, 1: 18},
				3: {1: 76, 0: 6, 2: 30, 4: 19},
				4: {3: 7, 2: 44, 0: 25, 1: 74}},
			5, 2, 0, []int{2, 4, 3, 0},
		},
	}
	solver := graph.FindShortestPathSolver{}
	for id, tc := range cyclicDiffLenGraphCases {
		actual := solver.DijkstraDiffWeight(tc.graph, tc.num, tc.start, tc.end)
		if !IdenticalArray(actual, tc.expected) {
			t.Fatalf("case#%d, expected:%+v, actual%+v", id, tc.expected, actual)
		}
	}
}

func TestBinaryTreeInorderTraversal(t *testing.T) {
	type testCase struct {
		root     *TreeNode
		expected []int
	}
	ts := []testCase{
		{&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 5,
				},
			},
		},
			[]int{4, 1, 6, 3, 5, 2},
		},
	}

	solver := tree.BinaryTreeInorderTraversal{}
	for i, tc := range ts {
		if actual := solver.BinaryTreeInorderTraversalSolve(tc.root); !IdenticalArray(actual, tc.expected) {
			t.Fatalf("case#%d, expected:%d, actual:%d", i, tc.expected, actual)
		}
	}
}
