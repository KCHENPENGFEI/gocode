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
	depth(root, &ans)
	return ans
}

func depth(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	// 计算左右子树的树高
	l := depth(root.Left, ans)
	r := depth(root.Right, ans)
	// 经过当前根节点的直径为两边子树的树高之和
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
