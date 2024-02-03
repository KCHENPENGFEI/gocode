package main

import (
	"fmt"
	. "golang/tree"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/10/29 下午2:00
 * @Desc:
 */

func main() {
	tree := TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 6},
	}
	fmt.Println(kthSmallest(&tree, 2))
}

func kthSmallest(root *TreeNode, k int) int {
	var arr []int
	order(root, &arr)
	return arr[k-1]
}

func order(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	order(root.Left, arr)
	*arr = append(*arr, root.Val)
	order(root.Right, arr)
}
