package heap

import (
	. "gogonotebook/common"
)

type MergeTwoSortedListSolver struct {
}

func (m *MergeTwoSortedListSolver) MergeTwoSortedListsSolve(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	l1Node := l1
	l2Node := l2
	var curList *ListNode
	var resultList *ListNode // 在go语言中，如何正确地理解指针和对象是实现链表算法的关联
	for {
		switch {
		case l1Node != nil && l2Node != nil:
			var next *ListNode
			if l1Node.Val < l2Node.Val {
				next = l1Node
				l1Node = l1Node.Next
			} else {
				next = l2Node
				l2Node = l2Node.Next
			}
			if curList == nil {
				curList = next
				resultList = curList
			} else {
				curList.Next = next
				curList = next
			}
		case l1Node != nil:
			curList.Next = l1Node
			goto RET
		case l2Node != nil:
			curList.Next = l2Node
			goto RET
		}
	}
RET:
	return resultList
}
