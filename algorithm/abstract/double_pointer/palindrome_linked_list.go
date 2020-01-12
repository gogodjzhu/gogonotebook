package double_pointer

import . "gogonotebook/common"

type PalindromeLinkedListSolver struct {
}

/**
快慢双指针解法

关键在找到链表的中点,然后将中点以后的链表翻转, 最后以翻转后的链表(只有一半长)为标准跟
原链表比较即可

巧妙的地方在于以翻转后的半长链表为准去比较, 这可以兼容奇偶长度的情况
*/
func (p *PalindromeLinkedListSolver) PalindromeLinkedListSolve(node *ListNode) bool {
	if node == nil || node.Next == nil {
		return true
	}
	fast, low := node, node // fast节点仅仅用于协助寻找中点, 对于偶数长度fast最后为nil
	for fast != nil && fast.Next != nil {
		low = low.Next
		fast = fast.Next.Next
	}
	rl := reverse(low)
	for rl != nil {
		if rl.Val != node.Val {
			return false
		}
		rl = rl.Next
		node = node.Next
	}
	return true
}

func reverse(l *ListNode) *ListNode {
	if l == nil || l.Next == nil {
		return l
	}
	lNext := reverse(l.Next)
	l.Next.Next = l
	l.Next = nil
	return lNext
}
