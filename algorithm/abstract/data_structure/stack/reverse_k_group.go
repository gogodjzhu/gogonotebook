package stack

import . "gogonotebook/common"

type ReverseKGroupSolver struct {
}

func (r ReverseKGroupSolver) ReverseKGroupSolve(head *ListNode, k int) *ListNode {
	stack := Stack{}
	for head != nil {
		stack.Add(head)
		head = head.Next
	}
	skipUnGroup := stack.Left() % k
	groupNum := stack.Left() / k

	var ret *ListNode
	for i := 0; i < skipUnGroup; i++ {
		cur := stack.Poll().(*ListNode)
		if ret != nil {
			cur.Next = ret
		}
		ret = cur
	}
	for i := 0; i < groupNum; i++ {
		first := stack.Poll().(*ListNode)
		tmp := first
		for j := 0; j < k-1; j++ {
			cur := stack.Poll().(*ListNode)
			tmp.Next = cur
			tmp = cur
		}
		tmp.Next = ret
		ret = first
	}
	return ret
}
