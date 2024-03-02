package main

import . "golang/tree"

/**
 * @Author: chenpengfei
 * @Date: 2024/2/3 下午1:00
 * @Desc: 二叉树的右视图：给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，
	返回从右侧所能看到的节点值。
*/

func rightSideView(root *TreeNode) []int {
	// 层序遍历二叉树，将每一层最后一个节点放入结果集中
	var s []*TreeNode
	var result []int
	if root != nil {
		s = append(s, root)
	}

	for len(s) > 0 {
		var tmp []*TreeNode
		for i, node := range s {
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
			// 每层最后一个节点
			if i == len(s)-1 {
				result = append(result, node.Val)
			}
		}
		s = tmp
	}
	return result
}

// rightSideView2
//
//	@Description: 二刷，每层先遍历右子树，然后每层第一个元素就是右视图
//	@param root
//	@return []int
func rightSideView2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var queue = []*TreeNode{root}
	var result []int
	for len(queue) > 0 {
		var level []*TreeNode
		for i := 0; i < len(queue); i++ {
			node := queue[i]
			if i == 0 {
				result = append(result, node.Val)
			}
			if node.Right != nil {
				level = append(level, node.Right)
			}
			if node.Left != nil {
				level = append(level, node.Left)
			}
		}
		queue = level
	}
	return result
}
