package abstract

import (
	"gogonotebook/algorithm/abstract/data_structure/graph"
	"testing"
)

func TestRouteBetweenNodes(t *testing.T) {
	type testCase struct {
		n             int
		graph         [][]int
		start, target int
		expected      bool
	}
	testCases := []testCase{
		{
			3,
			[][]int{{0, 1}, {0, 2}, {1, 2}, {1, 2}},
			0,
			2,
			true,
		},
		{
			5,
			[][]int{{0, 1}, {0, 2}, {0, 4}, {0, 4}, {0, 1}, {1, 3}, {1, 4}, {1, 3}, {2, 3}, {3, 4}},
			0,
			4,
			true,
		},
		{
			5,
			[][]int{{0, 1}, {0, 3}, {2, 3}, {1, 4}},
			0,
			2,
			false,
		},
	}
	solver := graph.RouteBetweenNodesSolver{}
	for id, tc := range testCases {
		actual := solver.RouteBetweenNodesSolve(tc.n, tc.graph, tc.start, tc.target)
		if actual != tc.expected {
			t.Fatalf("case#%d, expected:%v, actual:%v", id, tc.expected, actual)
		}
	}
}
