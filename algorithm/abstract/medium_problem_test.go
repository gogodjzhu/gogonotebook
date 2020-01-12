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
