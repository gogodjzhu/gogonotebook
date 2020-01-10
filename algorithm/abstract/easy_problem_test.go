package abstract

import (
	"fmt"
	"gogonotebook/algorithm/abstract/data_structure/heap"
	"gogonotebook/algorithm/abstract/recursion"
	. "gogonotebook/common"
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {
	type testCase struct {
		l1       *ListNode
		l2       *ListNode
		expected *ListNode
	}
	testCases := []testCase{
		{
			l1:       Arr2List([]int{1, 2, 3}),
			l2:       Arr2List([]int{3, 6}),
			expected: Arr2List([]int{1, 2, 3, 3, 6}),
		},
		{
			l1:       Arr2List([]int{}),
			l2:       Arr2List([]int{3, 6}),
			expected: Arr2List([]int{3, 6}),
		},
		{
			l1:       Arr2List([]int{}),
			l2:       Arr2List([]int{}),
			expected: Arr2List([]int{}),
		},
	}
	solver := heap.MergeTwoSortedListSolver{}
	for i, testCase := range testCases {
		actual := solver.MergeTwoSortedListsSolve(testCase.l1, testCase.l2)
		if !IsIdentical(actual, testCase.expected) {
			t.Fatal("case#", i, "expected", testCase.expected.ToString(), "actual", actual.ToString())
		}
	}
}

func TestReverseLinkedList(t *testing.T) {
	recursionSolver := recursion.ReverseLinkedList{}
	n := recursionSolver.ReverseLinkedListSolve(Arr2List([]int{}))
	fmt.Println(n.ToString())
}
