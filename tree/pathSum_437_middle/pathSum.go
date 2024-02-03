package main

import . "golang/tree"

/**
 * @Author: chenpengfei
 * @Date: 2024/2/3 下午2:44
 * @Desc: 给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。

路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
*/

func pathSum(root *TreeNode, targetSum int) int {
	count := 0
	helper(root, nil, targetSum, &count)
	return count
}

func helper(root *TreeNode, path []int, targetSum int, count *int) {
	if root == nil {
		return
	}
	sum := root.Val
	// 避免path为空进入不了循环的情况
	if sum == targetSum {
		*count++
	}
	// 从当前节点向根节点遍历
	for i := len(path) - 1; i >= 0; i-- {
		sum += path[i]
		if sum == targetSum {
			*count++
		}
	}
	path = append(path, root.Val)
	helper(root.Left, path, targetSum, count)
	helper(root.Right, path, targetSum, count)
}
