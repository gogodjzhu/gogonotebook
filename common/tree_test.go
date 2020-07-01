package common

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 5,
			},
		},
	}
	fmt.Println(tree.PreOrderTraversal())
	fmt.Println(PreOrderTraversal(tree))
}
