package main

import (
	"fmt"
	. "golang/tree"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/10/21 下午3:22
 * @Desc: 543. 二叉树的直径
给你一棵二叉树的根节点，返回该树的 直径 。

二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。

两节点之间路径的 长度 由它们之间边数表示。
*/

func main() {
	tree := TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(diameterOfBinaryTree(&tree))
}

// diameterOfBinaryTree
//
//	@Description: 经过根节点的直径为左子树树高+右子树树高
//	@param root
//	@return int
func diameterOfBinaryTree(root *TreeNode) int {
	var ans int
	help(root, &ans)
	return ans
}

func help(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	l := help(root.Left, ans)
	r := help(root.Right, ans)
	// 计算经过当前root的直径
	*ans = maxTwo(*ans, l+r)
	// 返回树高
	return maxTwo(l, r) + 1
}

func maxTwo(a, b int) int {
	if a > b {
		return a
	}
	return b
}
