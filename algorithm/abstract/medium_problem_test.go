package abstract

import (
	"gogonotebook/algorithm/abstract/double_pointer"
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
