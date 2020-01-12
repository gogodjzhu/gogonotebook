package abstract

import . "gogonotebook/common"

/**
// TODO 待完成
leetcode.q3.无重复字符的最长子串
eg.1.
input: s=abcab
return: 3 (最长字串为3)
*/
type LengthOfLongestSubstring interface {
	LengthOfLongestSubstringSolve(s string) int
}

/**
leetcode.q5.最长回文字串
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
leetcode.q15.三数之和
从给定数组中找出所有和为零的三数组合
*/
type ThreeNumberSum interface {
	ThreeNumberSumSolve(nums []int) [][]int
}

/**
leetcode.q16.最接近的三数之和
在指定数组nums找出和最接近target的三数和,假定存在唯一解
*/
type ThreeSumClosest interface {
	ThreeSumClosestSolve(nums []int, target int) int
}

/**
leetcode.q17.电话号码的字母组合

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
leetcode.q18.四数之和
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
leetcode.q19.删除链表的倒数第n个节点
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
leetcode.q22.括号生成
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
leetcode.q24.两两交换链表中的节点
在链表中交换相邻的两个节点. PS.要求必须实际地改变链表指针而不是单纯交换值.
eg.1.
input: 1->2->3->4
output: 2->1->4->3
*/
type SwapNodesInPairs interface {
	SwapNodesInPairsSolve(head *ListNode) *ListNode
}

/**
leetcode.q46.无重复数组的全排列
给定无重复数字组成的的数组，求所有全排列组合
*/
type Permutations interface {
	PermutationsSolve(nums []int) [][]int
}

/**
leetcode.q56.合并区间
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
leetcode.q61.旋转链表
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
leetcode.q62.不同路径
给定m*n的矩阵，每一步只能向左或者向下，求从左上角(0,0)走到右下角(m-1,n-1)有多少种路径选择.
input: m=3,n=2
output: 3
*/
type UniquePaths interface {
	UniquePathsSolve(m int, n int) int
}

/**
leetcode.q63.不同路径II
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
leetcode.q64.最小路径和
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
leetcode.q75.颜色分类
以红(0)白(1)蓝(2)的顺序原地排序颜色数组
eg.1.
input:  [2,0,2,1,1,0]
output: [0,0,1,1,2,2]
*/
type SortColors interface {
	SortColorsSolve(nums []int) []int
}

/**
leetcode.q91.解码方法
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
leetcode.q120.三角形最大(小)路径和
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
