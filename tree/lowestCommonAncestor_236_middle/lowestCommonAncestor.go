package main

import . "golang/tree"

/**
 * @Author: chenpengfei
 * @Date: 2024/2/3 下午3:10
 * @Desc: 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 在左树和右树分别搜索p，q节点
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 都不在左子树，那么最近节点就是右子树
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
