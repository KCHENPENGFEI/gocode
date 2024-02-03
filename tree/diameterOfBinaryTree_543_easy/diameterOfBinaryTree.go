package main

import (
	"fmt"
	. "golang/tree"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/10/21 下午3:22
 * @Desc:
 */

func main() {
	tree := TreeNode{
		Val: 1,
		//Left: &TreeNode{Val: 2},
		//Right: &TreeNode{Val: 3},
	}
	fmt.Println(diameterOfBinaryTree(&tree))
}

var ans int = 0

func diameterOfBinaryTree(root *TreeNode) int {
	depth(root)
	return ans
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := depth(root.Left)
	r := depth(root.Right)
	ans = max(ans, l+r)
	return max(l, r) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
