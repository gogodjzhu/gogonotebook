package common

import "testing"

func TestPriorityQueue(t *testing.T) {
	type testCase struct {
		pushItems map[int]interface{}
		popItems  []interface{}
	}
	testCases := []testCase{
		{
			pushItems: map[int]interface{}{
				2: 2,
				1: 1,
				3: 3,
			},
			popItems: []interface{}{
				3, 2, 1,
			},
		},
		{
			pushItems: map[int]interface{}{
				2: "2B",
				1: "1A",
				3: "3C",
			},
			popItems: []interface{}{
				"3C", "2B", "1A",
			},
		},
	}
	for idx, tc := range testCases {
		pq := NewPriorityQueue()
		for p, v := range tc.pushItems {
			pq.Push(p, v)
		}
		for _, v := range tc.popItems {
			if v != pq.Pop() {
				t.Fatalf("case#%d", idx)
			}
		}
	}
}
