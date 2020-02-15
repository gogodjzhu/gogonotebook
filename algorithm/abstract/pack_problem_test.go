package abstract

import (
	"gogonotebook/algorithm/abstract/dynamic_programming"
	"testing"
)

func TestZeroOnePack01(t *testing.T) {
	type testCase struct {
		V        int   // 最大容量
		W        []int // 物品重量
		expected int
	}
	testCases := []testCase{
		{
			3,
			[]int{1, 2, 3},
			3,
		},
		{
			10,
			[]int{1, 2, 3, 4},
			10,
		},
		{
			0,
			[]int{1, 2, 3},
			0,
		},
		{
			2,
			[]int{1, 2, 3},
			2,
		},
	}
	solver := dynamic_programming.PackProblemSolver{}
	for id, tc := range testCases {
		actual := solver.ZeroOnePack01(tc.V, tc.W)
		if actual != tc.expected {
			t.Fatalf("case#%d, expected:%d, actual:%d", id, tc.expected, actual)
		}
	}
}

func TestZeroOnePack02(t *testing.T) {
	type testCase struct {
		V        int   // 最大容量
		W        []int // 物品重量
		P        []int // 物品价格
		expected int
	}
	testCases := []testCase{
		{
			3,
			[]int{1, 2, 3},
			[]int{2, 4, 9},
			9,
		},
		{
			10,
			[]int{1, 2, 3, 4},
			[]int{1, 3, 3, 4},
			11,
		},
		{
			0,
			[]int{1, 2, 3},
			[]int{1, 2, 3},
			0,
		},
		{
			2,
			[]int{1, 2, 3},
			[]int{3, 2, 3},
			3,
		},
	}
	solver := dynamic_programming.PackProblemSolver{}
	for id, tc := range testCases {
		actual := solver.ZeroOnePack02(tc.V, tc.W, tc.P)
		if actual != tc.expected {
			t.Fatalf("case#%d, expected:%d, actual:%d", id, tc.expected, actual)
		}
	}
}

func TestZeroOnePack03(t *testing.T) {
	type testCase struct {
		V        int   // 最大容量
		W        []int // 物品重量
		P        []int // 物品价格
		expected int
	}
	testCases := []testCase{
		{
			3,
			[]int{1, 2, 3},
			[]int{2, 4, 9},
			9,
		},
		{
			5,
			[]int{2, 4},
			[]int{3, 4},
			0,
		},
		{
			0,
			[]int{1, 2, 3},
			[]int{1, 2, 3},
			0,
		},
		{
			7,
			[]int{1, 2, 6},
			[]int{3, 2, 3},
			6,
		},
		{
			7,
			[]int{1, 2, 6},
			[]int{3, 2, 3},
			6,
		},
	}
	solver := dynamic_programming.PackProblemSolver{}
	for id, tc := range testCases {
		actual := solver.ZeroOnePack03(tc.V, tc.W, tc.P)
		if actual != tc.expected {
			t.Fatalf("case#%d, expected:%d, actual:%d", id, tc.expected, actual)
		}
	}
}

func TestCompletePack01(t *testing.T) {
	type testCase struct {
		V        int   // 最大容量
		W        []int // 物品重量
		P        []int // 物品价格
		expected int
	}
	testCases := []testCase{
		{
			5,
			[]int{2, 4},
			[]int{2, 4},
			4,
		},
		{
			6,
			[]int{2, 4},
			[]int{2, 4},
			6,
		},
		{
			6,
			[]int{2, 4},
			[]int{5, 6},
			15,
		},
	}
	solver := dynamic_programming.PackProblemSolver{}
	for id, tc := range testCases {
		actual := solver.CompletePack01(tc.V, tc.W, tc.P)
		if actual != tc.expected {
			t.Fatalf("case#%d, expected:%d, actual:%d", id, tc.expected, actual)
		}
	}
}

func TestCompletePack02(t *testing.T) {
	type testCase struct {
		V        int   // 最大容量
		W        []int // 物品重量
		P        []int // 物品价格
		expected int
	}
	testCases := []testCase{
		{
			5,
			[]int{2, 4},
			[]int{2, 4},
			-1,
		},
		{
			6,
			[]int{2, 4},
			[]int{2, 4},
			6,
		},
		{
			7,
			[]int{2, 4},
			[]int{2, 4},
			-1,
		},
		{
			6,
			[]int{2, 4, 5},
			[]int{2, 4, 10},
			6,
		},
	}
	solver := dynamic_programming.PackProblemSolver{}
	for id, tc := range testCases {
		actual := solver.CompletePack02(tc.V, tc.W, tc.P)
		if actual != tc.expected {
			t.Fatalf("case#%d, expected:%d, actual:%d", id, tc.expected, actual)
		}
	}
}

func TestCompletePack03(t *testing.T) {
	type testCase struct {
		V        int   // 最大容量
		W        []int // 物品重量
		expected int
	}
	testCases := []testCase{
		{
			2,
			[]int{3},
			0,
		},
		{
			2,
			[]int{2},
			1,
		},
		{
			5,
			[]int{2, 4},
			0,
		},
		{
			6,
			[]int{2, 4},
			2,
		},
		{
			7,
			[]int{2, 4},
			0,
		},
		{
			6,
			[]int{2, 4, 5},
			2,
		},
	}
	solver := dynamic_programming.PackProblemSolver{}
	for id, tc := range testCases {
		actual := solver.CompletePack03(tc.V, tc.W)
		if actual != tc.expected {
			t.Fatalf("case#%d, expected:%d, actual:%d", id, tc.expected, actual)
		}
	}
}
