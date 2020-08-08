package tree

import . "gogonotebook/common"

type BinaryTreeInorderTraversalSolver struct {
}

func (b BinaryTreeInorderTraversalSolver) BinaryTreeInorderTraversalSolve(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return stackSolve(root)
}

/**
递归解法
*/
func recursionSolve(root *TreeNode) []int {
	arr := make([]int, 0)
	if root.Left != nil {
		arr = append(arr, recursionSolve(root.Left)...)
	}
	arr = append(arr, root.Val)
	if root.Right != nil {
		arr = append(arr, recursionSolve(root.Right)...)
	}
	return arr
}

/**
堆栈解法
*/
func stackSolve(root *TreeNode) []int {
	ret := make([]int, 0)
	st := NewStack()
	cur := root
	for cur != nil || st.Left() != 0 {
		for cur != nil {
			st.Add(cur)
			cur = cur.Left
		}
		node := st.Poll().(*TreeNode)
		ret = append(ret, node.Val)
		cur = node.Right
	}
	return ret
}
