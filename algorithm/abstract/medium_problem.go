package abstract

import . "gogonotebook/common"

/**
// TODO 待完成
leetcode.3.无重复字符的最长子串
eg.1.
input: s=abcab
return: 3 (最长字串为3)
*/
type LengthOfLongestSubstring interface {
	LengthOfLongestSubstringSolve(s string) int
}

/**
leetcode.5.最长回文字串
eg.1.
input: s="hello"
return: "ll"
eg.2.
input: s="wow"
return: "wow"
*/
type LongestPalindromicSubstring interface {
	LongestPalindromicSubstringSolve(s string) string
}

/**
leetcode.15.三数之和
从给定数组中找出所有和为零的三数组合
*/
type ThreeNumberSum interface {
	ThreeNumberSumSolve(nums []int) [][]int
}

/**
leetcode.16.最接近的三数之和
在指定数组nums找出和最接近target的三数和,假定存在唯一解
*/
type ThreeSumClosest interface {
	ThreeSumClosestSolve(nums []int, target int) int
}

/**
leetcode.17.电话号码的字母组合

|  1  |  2  |  3  |
|_____|_abc_|_def_|
|  4  |  5  |  6  |
|_ghi_|_jkl_|_mno_|
|  7  |  8  |  9  |
|_pqrs|_tuv_|wxyz_|

在九宫格键盘上，输入数字，返回所有可能的字符串组合
eg.1.
input:digits="23"
output:["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]
*/
type LetterCombinationsOfAPhoneNumber interface {
	LetterCombinationsOfAPhoneNumberSolve(digits string) []string
}

/**
leetcode.18.四数之和
从给定的数组中找出和与target相等的四数组合，组合不能重复
eg.1.
input: nums=[1, 0, -1, 0, -2, 2], target=0
output:
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/
type FourNumberSum interface {
	FourNumberSumSolve(nums []int, target int) [][]int
}

/**
leetcode.19.删除链表的倒数第n个节点
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
eg.1.
input: head=1->2->3->4->5, n = 2.
output: 1->2->3->5
*/
type RemoveNthNodeFromEndOfList interface {
	RemoveNthNodeFromEndOfListSolve(head *ListNode, n int) *ListNode
}

/**
// TODO 待完成
leetcode.22.括号生成
指定括号的对数，生成所有合法的组合
eg.1.
input: n=3
output:
[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/
type GenerateParentheses interface {
	GenerateParenthesesSolve(n int) []string
}

/**
leetcode.24.两两交换链表中的节点
在链表中交换相邻的两个节点. PS.要求必须实际地改变链表指针而不是单纯交换值.
eg.1.
input: 1->2->3->4
output: 2->1->4->3
*/
type SwapNodesInPairs interface {
	SwapNodesInPairsSolve(head *ListNode) *ListNode
}

/**
leetcode.31.下一个排列
将给定的数组替换为字典排序的下一个
eg.1.
input: 1,2,3
output: 1,3,2
eg.2.
input: 3,2,1
output: 1,2,3
eg.3.
input: 1,1,5
output: 1,5,1
*/
type NextPermutation interface {
	NextPermutationSolve(nums []int)
}

/**
leetcode.46.无重复数组的全排列
给定无重复数字组成的的数组，求所有全排列组合
*/
type Permutations interface {
	PermutationsSolve(nums []int) [][]int
}

/**
leetcode.48.旋转图像
准确地说应该是原地旋转n*n二维数组
eg.1.
input:
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
]
output:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
*/
type RotateImage interface {
	RotateImageSolve(matrix [][]int)
}

/**
leetcode.49.字母异位词分组
给定一组字符串数组, 将字母异位词组合为一组. 字母异位指字母相同, 但排列不同的字符串
eg.1.
input: ["eat", "tea", "tan", "ate", "nat", "bat"],
output:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
*/
type GroupAnagrams interface {
	GroupAnagramsSolve(strs []string) [][]string
}

/**
leetcode.56.合并区间
合并给出的区间集合
eg.1.
input: [[1,3],[2,6],[8,10],[15,18]]
output: [[1,6],[8,10],[15,18]]
eg.2
input: [[1,4],[4,5]]
output: [[1,5]]
*/
type MergeIntervals interface {
	MergeIntervalsSolve(intervals [][]int) [][]int
}

/**
leetcode.61.旋转链表
将链表每个节点向右移动k个位置,k非负
eg.1.
input: l=[1->2->3->4], k=2
output: [3->4->1->2]
eg.2.
input: l=[1->2->3->4], k=5
output: [4->1->2->3]
*/
type RotateList interface {
	RotateListSolve(l *ListNode, k int) *ListNode
}

/**
leetcode.62.不同路径
给定m*n的矩阵，每一步只能向左或者向下，求从左上角(0,0)走到右下角(m-1,n-1)有多少种路径选择.
input: m=3,n=2
output: 3
*/
type UniquePaths interface {
	UniquePathsSolve(m int, n int) int
}

/**
leetcode.63.不同路径II
给定矩阵,中间包含若干障碍(用数字1表示)，每一步只能向左或者向下，求从左上角(0,0)走到右下角(m-1,n-1)有多少种路径选择.
input:
[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
output: 2
*/
type UniquePathsII interface {
	UniquePathsIISolve(obstacleGrid [][]int) int
}

/**
leetcode.64.最小路径和
给定矩阵，每一个座标有一个数字，求从左上角到右下角经过数字和最小的路径
每一步只能往右或者往下走一步
eg.1.
input:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
output: 7 [1→3→1→1→1]
*/
type MinimumPathSum interface {
	MinimumPathSumSolve(grid [][]int) int
}

/**
leetcode.75.颜色分类
以红(0)白(1)蓝(2)的顺序原地排序颜色数组
eg.1.
input:  [2,0,2,1,1,0]
output: [0,0,1,1,2,2]
*/
type SortColors interface {
	SortColorsSolve(nums []int) []int
}

/**
leetcode.91.解码方法
给定只包含数字的字符串，使用以下规则将数字解码为英文字符，求有几种解法
'A' -> 1
'B' -> 2
...
'Z' -> 26
eg.1.
input: "12"
output: 2 可解码为:"AB"->[1,2],也可以解码为:"L"->[12]
eg.2.
input: "226"
output: 2 可解码为:"BZ"->[2,26],也可以解码为:"VF"->[22,6], 也可以解码为:"BBF"->[2,2,6]
*/
type DecodeWays interface {
	DecodeWaysSolve(s string) int
}

/**
leetcode.120.三角形最大(小)路径和
在数字三角形中寻找一条从顶部到底部的路径使得路径上所经过的数字之和最大。
路径上每一步都只能往左下 或 右下走。
eg.1.
input: data=
    7             [[7]
   3 8             [3 8]
  8 1 0  =保存为=>  [8 1 0]
 2 7 4 4           [2 7 4 4]
4 5 2 6 5          [4 5 2 6 5]]
return: 30 (路径为:7->3->8->7->5)
*/

type Triangle interface {
	TriangleSolve(data [][]int) int
}

/**
leetcode.133.克隆图
深拷贝图
*/
type CloneGraph interface {
	CloneGraphSolve(n *GraphNode) *GraphNode
}

/**
leetcode.207.课程表
必须选修 n 门课程，记为 0 到 n-1 。
在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]
给定课程总量n以及它们的先决条件p，请你判断是否可能完成所有课程的学习？

eg.1.
input: n=2, p=[[1,0]]
output: true

eg.2.
input: n=2, p=[[1,0],[0,1]]
output: false
*/
type CourseSchedule interface {
	CourseScheduleSolve(n int, p [][]int) bool
}

/**
leetcode.300.最长上升子序列
从给定序列中找出最长的上升子序列的长度
eg.1.
input: [10,9,2,5,3,7,101,18,1,2,3,4,5]
output: 4
最长上升子序列为[2,3,7,101], 他的长度为4
*/
type LongestIncreasingSubsequence interface {
	LongestIncreasingSubsequenceSolve(nums []int) int
}

/**
leetcode.322.零钱兑换
给定不同面额的硬币coins, 和 总金额amount, 求能凑成总金额的方案中最少需要金币的数量. 如果没法组成amount, 返回-1
eg.1.
input: coins=[1, 2, 5], amount=11
output: 3
eg.2.
input: coins=[2], amount=3
output: -1
*/
type CoinChange interface {
	CoinChangeSolve(coins []int, amount int) int
}

/**
leetcode.328.奇偶链表
给定一个单链表，把所有的奇数节点和偶数节点分别排在一起(奇在前, 偶在后)
注意这里的奇偶值得是下标
eg.1.
input:  1->2->3->4->5->NULL
output: 1->3->5->2->4->NULL
*/
type OddEvenLinkedList interface {
	OddEvenLinkedListSolve(head *ListNode) *ListNode
}

/**
leetcode.416.分割等子集
给定一个数组，判断是否可以将其分为两个和相等的子集
eg.1.
input: [1,5,11,5]
output: true ([1,5,5] 和 [11])
*/
type PartitionEqualSubsetSum interface {
	PartitionEqualSubsetSumSolve(nums []int) bool
}

/**
leetcode.494.目标和
给定个整型数组nums 和 目标t, 可以在数组任意元素上加+/-, 求最终的和等于t的方案数
注意:

数组非空，且长度不会超过20。
初始的数组的和不会超过1000。
保证返回的最终结果能被32位整数存下

eg.1.
input: nums: [1, 1, 1, 1, 1], S: 3
output: 5

explain:
-1+1+1+1+1 = 3
+1-1+1+1+1 = 3
+1+1-1+1+1 = 3
+1+1+1-1+1 = 3
+1+1+1+1-1 = 3
*/
type TargetSum interface {
	TargetSumSolve(nums []int, t int) int
}

/**
leetcode.518.零钱兑换II
给定不同面额的硬币coins, 和 总金额amount, 求能凑成总金额的方案数量
eg.1.
input: coins=[1, 2, 5], amount=5
output: 4
eg.2.
input: coins=[2], amount=3
output: 0
*/
type CoinChange2 interface {
	CoinChange2Solve(coins []int, amount int) int
}

/**
leetcode.743.网络延迟时间
有 N 个网络节点，标记为 1 到 N。
给定一个列表 times，表示信号经过有向边的传递时间。 times[i] = (u, v, w)，其中 u 是源节点，v 是目标节点， w 是一个信号从源节点传递到目标节点的时间。
求从给定节点K发出一个信号，所有节点都收到这个信号的最短时间。若无法让所有节点都收到，则返回-1

eg.1.
input: times = [[2,1,1],[2,3,1],[3,4,1]], N = 4, K = 2
画成图即:
[2]--1-->[1]
   \
    1-->[3]--1-->[4]
output: 2
*/
type NetworkDelayTime interface {
	NetworkDelayTimeSolve(times [][]int, N int, K int) int
}
