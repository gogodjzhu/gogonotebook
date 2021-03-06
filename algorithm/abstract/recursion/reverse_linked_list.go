package recursion

import . "gogonotebook/common"

type ReverseLinkedListSolver struct {
}

func (r *ReverseLinkedListSolver) ReverseLinkedListSolve(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}
	sub := r.ReverseLinkedListSolve(node.Next)
	node.Next.Next = node
	node.Next = nil // 这一步非常重要, 去除循环引用
	return sub
}
