package abstract

import (
	"gogonotebook/algorithm/abstract/data_structure/array"
	"gogonotebook/algorithm/abstract/data_structure/list"
	"gogonotebook/algorithm/abstract/double_pointer"
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
	t.Log(nums)
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
			t.Fatalf("case#%d, expected%+v, actual:%+v", id, tc.expected, tc.nums)
		}
	}
}
