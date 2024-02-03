package main

import . "golang/tree"

/**
 * @Author: chenpengfei
 * @Date: 2024/2/3 下午2:05
 * @Desc:给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。
判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。
如果存在，返回 true ；否则，返回 false 。

*/

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// 递归出口
	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		return true
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
