package main

import (
	"fmt"
	. "golang/tree"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/10/21 下午3:35
 * @Desc:
 */

func main() {
	tree := TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}
	fmt.Println(levelOrder(&tree))
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	var s []*TreeNode
	// 先把根节点写入s中
	if root != nil {
		s = append(s, root)
	}
	// 进入循环
	for len(s) > 0 {
		// 拿出所有该层的节点
		var tmp []*TreeNode
		var r []int
		for _, node := range s {
			r = append(r, node.Val)
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		result = append(result, r)
		s = tmp
	}
	return result
}

// levelOrder2
//
//	@Description: 二刷
//	@param root
//	@return [][]int
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		var tmp []*TreeNode
		var level []int
		for _, node := range queue {
			level = append(level, node.Val)
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		result = append(result, level)
		queue = tmp
	}
	return result
}
