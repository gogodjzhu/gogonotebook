package common

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
前序遍历
*/
func (t *TreeNode) PreOrderTraversal() []int {
	var ret []int
	if t == nil {
		return ret
	}
	ret = append(ret, t.Val)
	if t.Left != nil {
		leftArr := t.Left.PreOrderTraversal()
		ret = append(ret, leftArr...)
	}
	if t.Right != nil {
		rightArr := t.Right.PreOrderTraversal()
		ret = append(ret, rightArr...)
	}
	return ret
}

func PreOrderTraversal(t *TreeNode) []int {
	var ret []int
	if t == nil {
		return ret
	}
	ret = append(ret, t.Val)
	if t.Left != nil {
		leftArr := PreOrderTraversal(t.Left)
		ret = append(ret, leftArr...)
	}
	if t.Right != nil {
		rightArr := PreOrderTraversal(t.Right)
		ret = append(ret, rightArr...)
	}
	return ret
}

/**
中序遍历
*/
func (t *TreeNode) InOrderTraversal() []int {
	var ret []int
	if t == nil {
		return ret
	}
	if t.Left != nil {
		leftArr := t.Left.InOrderTraversal()
		ret = append(ret, leftArr...)
	}
	ret = append(ret, t.Val)
	if t.Right != nil {
		rightArr := t.Right.InOrderTraversal()
		ret = append(ret, rightArr...)
	}
	return ret
}

/**
后续遍历
*/
func (t *TreeNode) PostOrderTraversal() []int {
	var ret []int
	if t == nil {
		return ret
	}
	if t.Left != nil {
		leftArr := t.Left.PostOrderTraversal()
		ret = append(ret, leftArr...)
	}
	if t.Right != nil {
		rightArr := t.Right.PostOrderTraversal()
		ret = append(ret, rightArr...)
	}
	ret = append(ret, t.Val)
	return ret
}
