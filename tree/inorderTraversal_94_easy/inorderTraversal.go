package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
	. "golang/tree"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/10/15 下午3:06
 * @Desc:
 */

func main() {
	tree := TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(inorderTraversalLoop(&tree))
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	result = append(result, inorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)
	return result
}

type ColorTreeNode struct {
	*TreeNode
	RootColor int
}

// 通用格式三种遍历的循环写法
func inorderTraversalLoop(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	white, gray := 0, 1
	s := stack.New()
	s.Push(&ColorTreeNode{RootColor: white, TreeNode: root})
	for s.Len() > 0 {
		top := s.Pop().(*ColorTreeNode)
		if top.TreeNode == nil {
			continue
		}
		if top.RootColor == white {
			s.Push(&ColorTreeNode{RootColor: white, TreeNode: top.Right})
			s.Push(&ColorTreeNode{RootColor: gray, TreeNode: top.TreeNode})
			s.Push(&ColorTreeNode{RootColor: white, TreeNode: top.Left})
		} else {
			result = append(result, top.Val)
		}
	}
	return result
}
