package main

import (
	"fmt"
	. "golang/tree"
	"math"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/10/29 下午1:27
 * @Desc: 验证二叉搜索树
 */

func main() {
	tree := TreeNode{
		Val: 2147483647,
		//Left:  &TreeNode{Val: 2},
		//Right: &TreeNode{Val: 2},
	}
	fmt.Println(isValidBST(&tree))
}

// isValidBST
//
//	@Description: 每次给定一个范围(left, right)，root.Val必须在范围之中，递归时一直更新left和right的值
//	@param root
//	@return bool
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, left, right int) bool {
	if root == nil {
		return true
	}
	if root.Val <= left || root.Val >= right {
		return false
	}
	return helper(root.Left, left, min(right, root.Val)) && helper(root.Right, max(root.Val, left), right)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
