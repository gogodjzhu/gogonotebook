package abstract

import (
	"gogonotebook/algorithm/abstract/data_structure/heap"
	"gogonotebook/algorithm/abstract/data_structure/stack"
	"gogonotebook/algorithm/abstract/dynamic_programming"
	"gogonotebook/algorithm/abstract/merge"
	"gogonotebook/algorithm/abstract/recursion"
	. "gogonotebook/common"
	"testing"
)

func TestInterleavingString(t *testing.T) {
	type testCase struct {
		s1       string
		s2       string
		s3       string
		expected bool
	}
	testCases := []testCase{
		{s1: "a", s2: "a", s3: "aa",
			expected: true},
		{s1: "ab", s2: "cd", s3: "acbd",
			expected: true},
		{s1: "ab", s2: "cd", s3: "abcd",
			expected: true},
		{s1: "", s2: "cd", s3: "cd",
			expected: true},
		{s1: "c", s2: "ca", s3: "cac",
			expected: true},
		{s1: "db", s2: "b", s3: "cbb",
			expected: false},
	}
	dynamicSolver := dynamic_programming.InterleavingStringSolver{}
	for _, testCase := range testCases {
		actual := dynamicSolver.InterleavingStringSolve(testCase.s1, testCase.s2, testCase.s3)
		if actual != testCase.expected {
			t.Fatal("case:[", testCase.s1, testCase.s2, testCase.s3, "]", ", expected:",
				testCase.expected, ", actual:", actual)
		}
	}
}

func TestMergeKSortedList(t *testing.T) {
	type testCase struct {
		lists    []*ListNode
		expected *ListNode
	}
	testCases := []testCase{
		{
			lists: []*ListNode{
				Arr2List([]int{1, 3, 5}),
				Arr2List([]int{2, 4, 6}),
				Arr2List([]int{6, 7}),
			},
			expected: Arr2List([]int{1, 2, 3, 4, 5, 6, 6, 7})},
		{
			lists: []*ListNode{
				Arr2List([]int{1}),
				Arr2List([]int{2, 4, 6}),
				Arr2List([]int{}),
			},
			expected: Arr2List([]int{1, 2, 4, 6})},
	}
	dsSolver := heap.MergeKSortedListSolver{}
	for i, testCase := range testCases {
		actual := dsSolver.MergeKSortedListSolve(testCase.lists)
		if !IsIdentical(actual, testCase.expected) {
			t.Fatal("case#", i, "case:", testCase.lists,
				"expected:", testCase.expected.ToString(), "actual:", actual.ToString())
		}
	}
	mergeSolver := merge.MergeKSortedListSolver{}
	for i, testCase := range testCases {
		actual := mergeSolver.MergeKSortedListSolve(testCase.lists)
		if !IsIdentical(actual, testCase.expected) {
			t.Fatal("case#", i, "case:", testCase.lists,
				"expected:", testCase.expected.ToString(), "actual:", actual.ToString())
		}
	}
}

func TestLargestRectangleInHistogram(t *testing.T) {
	type testCase struct {
		nums     []int
		expected int
	}
	testCases := []testCase{
		//{
		//	[]int{2, 1, 3, 2},
		//	4,
		//},
		//{
		//	[]int{2, 1, 5, 6, 2, 3},
		//	10,
		//},
		//{
		//	[]int{1, 2, 2},
		//	4,
		//},
		//{
		//	[]int{0,2,0},
		//	2,
		//},
		{
			[]int{5, 4, 1, 2},
			8,
		},
	}
	dynamicSolver := dynamic_programming.LargestRectangleInHistogramSolver{}
	for id, tc := range testCases {
		actual := dynamicSolver.LargestRectangleInHistogramSolve(tc.nums)
		if actual != tc.expected {
			t.Fatalf("case#%d, expected:%d, actual:%d", id, tc.expected, actual)
		}
	}
}

/**
[
  ["1","0","1","0","0"],
  ["1","0","1","1","1"],
  ["1","1","1","1","1"],
  ["1","0","0","1","0"]
]
*/
func TestMaximalRectangle(t *testing.T) {
	type testCase struct {
		nums     [][]byte
		expected int
	}
	testCases := []testCase{
		{
			[][]byte{
				{'1', '1', '1'},
				{'1', '1', '0'},
				{'1', '1', '0'},
			},
			6,
		},
		{
			[][]byte{
				{'1', '1', '1'},
				{'1', '0', '1'},
				{'1', '1', '1'},
			},
			3,
		},
		{
			[][]byte{
				{'1', '1', '1', '1', '1', '1', '1'},
				{'1', '1', '0', '0', '1', '0', '1'},
				{'1', '1', '0', '0', '0', '1', '0'},
			},
			7,
		},
	}
	dynamicSolver := dynamic_programming.MaximalRectangleSolver{}
	for id, tc := range testCases {
		actual := dynamicSolver.MaximalRectangleSolve(tc.nums)
		if actual != tc.expected {
			t.Fatalf("case#%d, expected:%d, actual:%d", id, tc.expected, actual)
		}
	}
}

func TestReverseKGroup(t *testing.T) {
	type testCase struct {
		list     *ListNode
		k        int
		expected *ListNode
	}
	testCases := []testCase{
		{Arr2List([]int{}), 1, Arr2List([]int{})},
		{Arr2List([]int{1}), 1, Arr2List([]int{1})},
		{Arr2List([]int{1, 2}), 1, Arr2List([]int{1, 2})},
		{Arr2List([]int{1, 2}), 2, Arr2List([]int{2, 1})},
		{Arr2List([]int{1, 2, 3}), 2, Arr2List([]int{2, 1, 3})},
		{Arr2List([]int{1, 2, 3}), 3, Arr2List([]int{3, 2, 1})},
		{Arr2List([]int{1, 2, 3, 4, 5}), 3, Arr2List([]int{3, 2, 1, 4, 5})},
	}
	solver := recursion.ReverseKGroupSolver{}
	for id, tc := range testCases {
		actual := solver.ReverseKGroupSolve(tc.list, tc.k)
		if !IsIdentical(actual, tc.expected) {
			t.Fatal("case#", id, "|expected:", tc.expected.ToString(), "|actual:", actual.ToString())
		}
	}
}

func TestTrappingRainWater(t *testing.T) {
	type testCase struct {
		arr      []int
		expected int
	}

	testCases := []testCase{
		{
			[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			6,
		},
		{
			[]int{0, 0, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			5,
		},
		{
			[]int{3, 0, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			9 + 6,
		},
	}
	solver := stack.TrappingRainWaterSolver{}
	for id, tc := range testCases {
		actual := solver.TrappingRainWaterSolve(tc.arr)
		if actual != tc.expected {
			t.Fatalf("case#%d, expected:%d, actual:%d", id, tc.expected, actual)
		}
	}
}

func TestEditDistance(t *testing.T) {
	type testCase struct {
		word1, word2 string
		expected     int
	}
	testCases := []testCase{
		{"nice", "nike", 1},
		{"nice", "ice", 1},
		{"nice", "nidle", 2},
	}
	solver := dynamic_programming.EditDistanceSolver{}
	for id, tc := range testCases {
		actual := solver.EditDistanceSolve(tc.word1, tc.word2)
		if actual != tc.expected {
			t.Fatalf("case#%d, expected:%d, actual:%d", id, tc.expected, actual)
		}
	}
}
