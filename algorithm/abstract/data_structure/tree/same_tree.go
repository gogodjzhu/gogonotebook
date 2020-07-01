package tree

import . "gogonotebook/common"

type IsSameTreeSolver struct {
}

/**
注意树相等 != 树的遍历结果相等
*/
func (i IsSameTreeSolver) IsSameTreeSolve(p *TreeNode, q *TreeNode) bool {
	if q == nil && p == nil {
		return true
	}
	if (p != nil && q == nil) || (p == nil && q != nil) || p.Val != q.Val {
		return false
	}
	if (p.Left == nil && q.Left != nil) ||
		(p.Left != nil && q.Left == nil) ||
		!i.IsSameTreeSolve(p.Left, q.Left) {
		return false
	}
	if (p.Right == nil && q.Right != nil) ||
		(p.Right != nil && q.Right == nil) ||
		!i.IsSameTreeSolve(p.Right, q.Right) {
		return false
	}
	return true
}
