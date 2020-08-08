package abstract

import (
	"gogonotebook/algorithm/abstract/data_structure/graph"
	"gogonotebook/algorithm/abstract/data_structure/list"
	_map "gogonotebook/algorithm/abstract/data_structure/map"
	"gogonotebook/algorithm/abstract/double_pointer"
	"gogonotebook/algorithm/abstract/dynamic_programming"
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

func TestShortestCompletingWord(t *testing.T) {
	type testCase struct {
		licensePlate string
		words        []string
		expected     string
	}
	testCases := []testCase{
		{
			"so",
			[]string{"sad", "so", "soso"},
			"so",
		},
		{
			"1s3 PSt",
			[]string{"step", "steps", "stripe", "stepple"},
			"steps",
		},
		{
			"1s3 456",
			[]string{"looks", "pest", "stew", "show"},
			"pest",
		},
	}
	solver := _map.ShortestCompletingWordSolver{}
	for id, testCase := range testCases {
		actual := solver.ShortestCompletingWordSolve(testCase.licensePlate, testCase.words)
		if actual != testCase.expected {
			t.Fatalf("case#%d, expected:%s, actual:%s\n", id, testCase.expected, actual)
		}
	}
}

func TestFindTownJudge(t *testing.T) {
	type testCase struct {
		n        int
		trust    [][]int
		expected int
	}
	testCases := []testCase{
		{
			2,
			[][]int{{1, 2}},
			2,
		},
		{
			3,
			[][]int{{1, 3}, {2, 3}},
			3,
		},
		{
			3,
			[][]int{{1, 3}, {2, 3}, {3, 1}},
			-1,
		},
	}
	solver := graph.FindTheTownJudge{}
	for id, tc := range testCases {
		actual := solver.FindTheTownJudgeSolve(tc.n, tc.trust)
		if actual != tc.expected {
			t.Fatalf("case#%d, expeced:%v, actual:%v", id, tc.expected, actual)
		}
	}
}

func TestMaximumSubarray(t *testing.T) {
	type testCase struct {
		nums     []int
		expected int
	}
	testCases := []testCase{
		{
			[]int{1},
			1,
		},
		{
			[]int{3, 2, 1},
			6,
		},
		{
			[]int{1, -2, 3},
			3,
		},
		{
			[]int{-1, 3, -2, 5},
			6,
		}, {
			[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			6,
		},
	}
	solver := dynamic_programming.MaximumSubarraySolver{}
	for i, tc := range testCases {
		actual := solver.MaximumSubarraySolve(tc.nums)
		if actual != tc.expected {
			t.Fatalf("case#%d, expected:%d, actual:%d", i, tc.expected, actual)
		}
	}
}
