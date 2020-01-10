package abstract

import (
	"gogonotebook/algorithm/abstract/data_structure/heap"
	"gogonotebook/algorithm/abstract/dynamic_programming"
	"gogonotebook/algorithm/abstract/merge"
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

// TODO 待完成
func TestWildcardMatching(t *testing.T) {
	type testCase struct {
		s        string
		p        string
		expected bool
	}
	testCases := []testCase{
		// {s:"",p:"",expected:true},
		// {s:"",p:"*",expected:true},
		// {s:"",p:"?",expected:false},
		// {s:"aa",p:"a",expected:false},
		// {s:"aa",p:"*",expected:true},
		// {s:"aa",p:"aa",expected:true},
		// {s:"aa",p:"a?",expected:true},
		// {s:"aaa",p:"a?a",expected:true},
		// {s:"aaa",p:"a*a",expected:true},
		// {s:"aaaa",p:"a*a",expected:true},
		// {s:"aabb",p:"*a**b",expected:true},
		// {s:"aaccccbb",p:"*a?*b",expected:true},
		// {s:"acdcb",p:"*a*c?b",expected:false},
		{s: "leetcode", p: "*e*t?d*", expected: false},
	}
	dynamicSolver := dynamic_programming.WildcardMatchingSolver{}
	for _, testCase := range testCases {
		actual := dynamicSolver.WildcardMatching(testCase.s, testCase.p)
		if actual != testCase.expected {
			t.Fatal("case:[", testCase.s, testCase.p, "]", ", expected:",
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
	dynamicSolver := dynamic_programming.LargestRectangleInHistogramSolver{}
	if dynamicSolver.LargestRectangleInHistogramSolve([]int{5}) != 5 {
		t.Fatal()
	}
	if dynamicSolver.LargestRectangleInHistogramSolve([]int{4, 2, 1}) != 4 {
		t.Fatal()
	}
	if dynamicSolver.LargestRectangleInHistogramSolve([]int{1, 2, 4}) != 4 {
		t.Fatal()
	}
	if dynamicSolver.LargestRectangleInHistogramSolve([]int{1, 9, 20}) != 20 {
		t.Fatal()
	}
	if dynamicSolver.LargestRectangleInHistogramSolve([]int{20, 9, 1}) != 20 {
		t.Fatal()
	}
	if dynamicSolver.LargestRectangleInHistogramSolve([]int{1, 1, 6, 2}) != 6 {
		t.Fatal()
	}
	if dynamicSolver.LargestRectangleInHistogramSolve([]int{2, 1, 3, 2}) != 4 {
		t.Fatal()
	}
	if dynamicSolver.LargestRectangleInHistogramSolve([]int{2, 1, 5, 6, 2, 3}) != 10 {
		t.Fatal()
	}
	if dynamicSolver.LargestRectangleInHistogramSolve([]int{1, 2, 2}) != 4 {
		t.Fatal()
	}
}
