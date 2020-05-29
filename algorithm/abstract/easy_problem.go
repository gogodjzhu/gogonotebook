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

/**
leetcode.q242.有效的字母异位词
判断给定的两个字符串是否为字母异位词, 所谓字母异位词指的是包含相同的字符但位置不同的词
eg.1.
input: s="anagram",t="nagaram"
output:true
eg.2.
input: s="rat",t="car"
output:false
*/
type ValidAnagram interface {
	ValidAnagramSolve(s, t string) bool
}

/**
leetcode.q509.斐波那契数
给定函数:
F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
求F(N)
eg.1.
input: N=3
F(3) = F(2) + F(1) = (F(1) + F(0)) + F(1) = 2
*/
type FibonacciNumber interface {
	FibonacciNumberSolve(n int) int
}

/**
leetcode.q748.最短完整词
如果单词列表（words）中的一个单词包含牌照（licensePlate）中所有的字母，那么我们称之为完整词。在所有完整词中，最短的单词我们称之为最短完整词
如给定licensePlate="sa", 则words=["sad", "see", "sand"]中"sad"是最短完整词
当在words拥有相同长度的最短完整词时，取words中位置靠前的作为结果

eg.1.
input: licensePlate = "1s3 PSt", words = ["step", "steps", "stripe", "stepple"]
output："steps"
注意, 我们仅关心字母，数字、空格或者其余额特殊字符不考虑
*/
type ShortestCompletingWord interface {
	ShortestCompletingWordSolve(licensePlate string, words []string) string
}

/**
leetcode.q997. 找到小镇的法官
本质是在DAG中寻找一个节点，它满足下面的条件:
1. 其余所有节点均直接指向此节点
2. 此节点未指向任何别的节点
*/
type FindTheTownJudge interface {
	FindTheTownJudgeSolve(n int, trust [][]int) int
}

/**
gogodjzhu.notebook.1. 各种方法寻找最短路径的算法，及相关性能比较
*/
type FindShortestPath interface {
	DFS(graph map[int][]int, num, start, end int) []int
	BFS(graph map[int][]int, num, start, end int) []int
	Dijkstra(graph map[int][]int, num, start, end int) []int
	BellmanFord(graph map[int][]int, num, start, end int) []int
}
