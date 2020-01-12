package abstract

import (
	"gogonotebook/algorithm/abstract/data_structure/list"
	"gogonotebook/algorithm/abstract/double_pointer"
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
	solver := list.MergeTwoSortedListSolver{}
	for i, testCase := range testCases {
		actual := solver.MergeTwoSortedListsSolve(testCase.l1, testCase.l2)
		if !IsIdentical(actual, testCase.expected) {
			t.Fatal("case#", i, "expected", testCase.expected.ToString(), "actual", actual.ToString())
		}
	}
}

func TestPalindromeLinkedList(t *testing.T) {
	type testCase struct {
		l  *ListNode
		is bool
	}
	testCases := []testCase{
		{
			Arr2List([]int{1}),
			true,
		},
		{
			Arr2List([]int{1, 2}),
			false,
		},
		{
			Arr2List([]int{1, 1, 2, 1}),
			false,
		},
		{
			Arr2List([]int{1, 2, 3, 2, 1}),
			true,
		},
		{
			Arr2List([]int{1, 2, 3, 4, 2, 1}),
			false,
		},
	}
	solver := double_pointer.PalindromeLinkedListSolver{}
	for i, testCase := range testCases {
		actual := solver.PalindromeLinkedListSolve(testCase.l)
		if actual != testCase.is {
			t.Fatal("case#", i)
		}
	}
}
