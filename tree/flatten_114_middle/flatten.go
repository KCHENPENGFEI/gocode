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

// flatten2
//
//	@Description: 二刷使用递归
//	@param root
func flatten2(root *TreeNode) {
	helper(root)
}

// helper
//
//	@Description: 先进行递归左右子树，然后将右子树替换成左子树
//
// 再遍历找到当前右子树的最右下角节点，将原来的右子树拼接上去，然后使root的左子树为空
//
//	@param root
//	@return *TreeNode
func helper(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	// 递归处理左子树
	left := helper(root.Left)
	helper(root.Right)

	tmp := root.Right
	// 右子树替换成左子树
	root.Right = left
	p := root
	for p.Right != nil {
		u
		p = p.Right
	}
	p.Right = tmp
	root.Left = nil
	return root
}
