package main

import (
	"fmt"
	. "golang/tree"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/10/15 下午4:16
 * @Desc:
 */

func main() {
	tree := TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 10}},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(maxDepth(&tree))
}

// 最大深度等于左子树和右子树中最大深度+1
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	d := max(maxDepth(root.Left), maxDepth(root.Right))
	return d + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type ColorTreeNode struct {
	*TreeNode
	RootColor int
}
