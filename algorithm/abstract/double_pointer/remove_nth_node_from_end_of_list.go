package double_pointer

import . "gogonotebook/common"

type RemoveNthNodeFromEndOfListSolver struct {
}

/**
本题实现要求的难度不大，关键是如何优化性能 和 让代码更优雅，本解使用了两个小技巧值得学习：

Tips.1. 使用哑节点来简化头部节点为空的一些边界判断

Tips.2. 使用双指针来实现O(n)时间复杂度, 具体地：
1. 使用一个临时节点tmpNode跟dummy节点一致，移动n
2. 再将tmpNode节点移动到链表尾部,此时需要移动的次数是(len(list) - n), 这也刚好是从头
   部移动到倒数第n个节点的需要经过的移动次数. 所以我们在移动tmpNode的时候同步移动另一
   个从头开始移动跟的moveNode节点. 那么在tmpNode到尾的时候，moveNode刚好在倒数n的位置

eg.

有:
list: n1->n2->n3->n4
n = 2

令:
n1->n2->n3->n4
|
|tmpNode
|moveNode

先让tmpMode移动2得到:
n1->n2->n3->n4
|       |
|       |tmpNode
|moveNode

然后让tmpNode继续移动到尾部, 同时moveNode一起移动:
n1->n2->n3->n4->n5
        |       |
        |       |tmpNode
        |moveNode
那么此时moveNode刚好落在倒数第n个节点上,基于同样的逻辑, 需要移除第n个节点, 只需要稍
作改动即可

*/
func (r RemoveNthNodeFromEndOfListSolver) RemoveNthNodeFromEndOfListSolve(head *ListNode, n int) *ListNode {

	dummy := &ListNode{}
	dummy.Next = head

	tmpNode := dummy
	moveNode := dummy

	for i := 0; i < n+1; i++ { // 由于使用了dummy节点, n+1
		tmpNode = tmpNode.Next
	}

	for tmpNode != nil {
		tmpNode = tmpNode.Next
		moveNode = moveNode.Next
	}

	moveNode.Next = moveNode.Next.Next

	return dummy.Next
}
