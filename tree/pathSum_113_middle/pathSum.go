package main

import (
	"fmt"
	. "golang/tree"
)

/**
 * @Author: chenpengfei
 * @Date: 2024/2/3 下午2:22
 * @Desc: 给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

叶子节点 是指没有子节点的节点。
*/

func main() {
	tree := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}},
		Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}},
	}
	t := 22
	r := pathSum(tree, t)
	fmt.Println(r)
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	helper(root, nil, &result, targetSum)
	return result
}

func helper(root *TreeNode, path []int, paths *[][]int, targetSum int) {
	if root == nil {
		return
	}
	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		// 一定要拷贝一份数据
		var copyPath = make([]int, len(path))
		copy(copyPath, path)
		copyPath = append(copyPath, root.Val)
		// paths必须要是指针类型才能append
		*paths = append(*paths, copyPath)
	}
	path = append(path, root.Val)
	helper(root.Left, path, paths, targetSum-root.Val)
	helper(root.Right, path, paths, targetSum-root.Val)
}
