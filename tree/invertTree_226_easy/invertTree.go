package main

import (
	"fmt"
	. "golang/tree"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/10/15 下午4:31
 * @Desc:
 */

func main() {
	tree := TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 10}},
		Right: &TreeNode{Val: 3},
	}
	r := invertTree(&tree)
	fmt.Println(r)
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left, root.Right = right, left
	return root
}
