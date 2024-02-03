package main

import . "golang/tree"

/**
 * @Author: chenpengfei
 * @Date: 2024/2/3 下午1:58
 * @Desc: 给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历，
postorder 是同一棵树的后序遍历，请你构造并返回这颗 二叉树 。


*/

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	// 后序遍历数组最后个元素是根节点
	root := &TreeNode{
		Val: postorder[len(postorder)-1],
	}
	for i := 0; i < len(inorder); i++ {
		// 在中序数组中找到元素
		if root.Val == inorder[i] {
			leftInorder := inorder[:i]
			rightInorder := inorder[i+1:]
			rightPostorder := postorder[len(postorder)-1-len(rightInorder) : len(postorder)-1]
			leftPostorder := postorder[:len(postorder)-1-len(rightInorder)]
			root.Left = buildTree(leftInorder, leftPostorder)
			root.Right = buildTree(rightInorder, rightPostorder)
			break
		}
	}
	return root
}
