package stack

import . "gogonotebook/common"

type ReverseLinkedListSolver struct {
}

func (r *ReverseLinkedListSolver) ReverseLinkedListSolve(node *ListNode) *ListNode {
	if node == nil {
		return node
	}
	stack := &Stack{}
	for node != nil {
		stack.Add(node)
		node = node.Next
	}
	head := stack.Poll().(*ListNode)
	idx := head
	for stack.Left() > 0 {
		idx.Next = stack.Poll().(*ListNode)
		idx = idx.Next
	}
	idx.Next = nil
	return head
}
