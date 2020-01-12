package recursion

import . "gogonotebook/common"

type ReverseKGroupSolver struct {
}

func (r ReverseKGroupSolver) ReverseKGroupSolve(head *ListNode, k int) *ListNode {
	nextGroup := head
	groupNum := 0
	for groupNum = 0; groupNum < k && nextGroup != nil; groupNum++ {
		nextGroup = nextGroup.Next
	}

	if groupNum < k { // 未满组
		return head
	} else {
		nextGroup = r.ReverseKGroupSolve(nextGroup, k)
		tmp := reverseK(head, k)
		head.Next = nextGroup // 头指向nextGroup
		return tmp
	}
}

func reverseK(node *ListNode, k int) *ListNode {
	if k > 1 {
		n := reverseK(node.Next, k-1)
		node.Next.Next = node
		return n
	} else {
		return node
	}
}
