package main

import . "golang/tree"

/**
 * @Author: chenpengfei
 * @Date: 2024/2/3 下午1:40
 * @Desc: 给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历，
inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// 前序数组第一个元素是根节点
	root := &TreeNode{
		Val: preorder[0],
	}
	for i := 0; i < len(inorder); i++ {
		// 在中序数组中找到元素
		if root.Val == inorder[i] {
			leftInorder := inorder[:i]
			rightInorder := inorder[i+1:]
			leftPreorder := preorder[1 : 1+len(leftInorder)]
			rightPreorder := preorder[1+len(leftInorder):]
			root.Left = buildTree(leftPreorder, leftInorder)
			root.Right = buildTree(rightPreorder, rightInorder)
			break
		}
	}
	return root
}
