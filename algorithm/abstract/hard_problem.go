package abstract

import . "gogonotebook/common"

/**
leetcode.q23.合并k个有序链表
eg.1.
input:
[
  1->4->5,
  1->3->4,
  2->6
]
output:
1->1->2->3->4->4->5->6
*/
type MergeKSortedList interface {
	MergeKSortedListSolve(lists []*ListNode) *ListNode
}

/**
leetcode.q25.K个一组翻转链表
eg.1.
input:
1->2->3->4->5
k=2
output:
2->1->4->3->5
eg.2.
input:
1->2->3->4->5
k=3
output:
3->2->1->4->5
*/
type ReverseKGroup interface {
	ReverseKGroupSolve(head *ListNode, k int) *ListNode
}

/**
// TODO 待完成
leetcode.q44.通配符匹配
给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。
s只包含a-z小写字符
p值包含a-z小写字符,'?','*'
eg.1.
input: s="aa",p="a"
return: false
eg.2.
input: s="aabb",p="*a*b"
return: true
*/
type WildcardMatching interface {
	WildcardMatchingSolve(s string, p string) bool
}

/**
leetcode.q72.编辑距离
给定两个单词 word1 和 word2，计算出将 word1 转换成 word2 所使用的最少操作数。
eg.1.
input: word1="horse", word2="ros"
return: 3
解释:
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
*/
type EditDistance interface {
	EditDistanceSolve(word1, word2 string) int
}

/**
leetcode.q97.交错字符串
给定三个字符串 s1, s2, s3, 验证 s3 是否是由 s1 和 s2 交错组成的。
eg.1.
input: s1="aabcc", s2="dbbca", s3="aadbbcbcac"
output: true
eg.2.
input: s1="aabcc", s2="dbbca", s3="aadbbbaccc"
output: false
*/
type InterleavingString interface {
	InterleavingStringSolve(s1 string, s2 string, s3 string) bool
}

/**
// TODO 待完成
leetcode.q84.柱状图中的最大矩阵
在给出的柱状图中, 找出能框出来的最大连续矩形面积
eg.1.
input: [2,1,5,6,2,3]
output:10
*/
type LargestRectangleInHistogram interface {
	LargestRectangleInHistogramSolve(heights []int) int
}
