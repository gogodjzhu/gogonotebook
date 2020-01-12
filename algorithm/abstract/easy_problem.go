package abstract

import . "gogonotebook/common"

/**
leetcode.q21.合并两个有序链表
eg.1.
input:  l1=1->2->4, l2=1->3->4
output: 1->1->2->3->4->4
*/
type MergeTwoSortedLists interface {
	MergeTwoSortedListsSolve(l1, l2 *ListNode) *ListNode
}

/**
leetcode.q28.实现strStr()
从给定的字符串s中寻找字串t,返回t在s中的起始位置，不存在则返回-1
eg.1.
input:s="hello",t="ll"
return:2
*/
type FindSubstring interface {
	FindSubstringSolver(s, t string) int
}

/**
leetcode.q53.求最大子序和
给定一个数字数组，求一个和最大的连续子序列
eg.1.
input: nums=[-2,1,-3,4,-1,2,1,-5,4]
output: 6(对应子序列为:[4,-1,2,1])
*/
type MaximumSubarray interface {
	MaximumSubarraySolve(nums []int) int
}

/**
leetcode.q70.爬梯子
每次只能上一个或者两个台阶，求上到n层台阶上有几种方法
eg.1.
input: total=3
return: [[1,1,1][1,2][2,1]]
*/
type ClimbingStairs interface {
	ClimbingStairsSolve(total int) int
}

/**
leetcode.q121.买卖股票的最佳时机
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
求最多只完成一笔交易(一次买入+一次卖出 or 不交易)能达成的最大收益
eg.1.
input: prices=[7,1,5,3,6,4]
output: 5 在第2天(价格为1)买入, 第5天(价格为6)卖出， 5=6-1
eg.2.
input: prices=[7,6,5,4]
output: 0 不交易
*/
type BestTimeToBuyAndSellStock interface {
	BestTimeToBuyAndSellStockSolve(prices []int) int
}

/**
leetcode.q206.反转链表
*/
type ReverseLinkedList interface {
	ReverseLinkedListSolve(node *ListNode) *ListNode
}

/**
leetcode.q234.回文链表
判断链表是否满足回文条件
eg.1.
input:1->2->1
output:true
input:1->2->3
output:false
*/
type PalindromeLinkedList interface {
	PalindromeLinkedListSolve(node *ListNode) bool
}
