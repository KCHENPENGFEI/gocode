package main

import . "golang/tree"

/**
 * @Author: chenpengfei
 * @Date: 2024/2/3 下午1:15
 * @Desc:
 */

func main() {
	tree := TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}},
	}
	flatten(&tree)
}

func flatten(root *TreeNode) {
	// 先进行一个先序遍历得到遍历数组
	slice := travel(root)
	var last *TreeNode
	for i := 0; i < len(slice); i++ {
		slice[i].Left = nil
		slice[i].Right = nil
		if last != nil {
			last.Right = slice[i]
		}
		last = slice[i]
	}
}

func travel(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	var r []*TreeNode
	r = append(r, root)
	r = append(r, travel(root.Left)...)
	r = append(r, travel(root.Right)...)
	return r
}
