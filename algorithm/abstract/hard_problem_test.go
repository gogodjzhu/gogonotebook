package abstract

import (
	"gogonotebook/algorithm/abstract/data_structure/heap"
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
		{
			[]int{5},
			5,
		},
		{
			[]int{4, 2, 1},
			4,
		},
		{
			[]int{1, 2, 4},
			4,
		},
		{
			[]int{1, 9, 20},
			20,
		},
		{
			[]int{20, 9, 1},
			20,
		},
		{
			[]int{1, 1, 6, 2},
			6,
		},
		{
			[]int{2, 1, 3, 2},
			4,
		},
		{
			[]int{2, 1, 5, 6, 2, 3},
			10,
		},
		{
			[]int{1, 2, 2},
			4,
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
