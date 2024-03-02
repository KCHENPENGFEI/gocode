package main

import (
	"fmt"
	. "golang/tree"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/10/15 下午4:39
 * @Desc:
 */

func main() {
	tree := TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 10}},
		Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 10}},
	}
	r := isSymmetric(&tree)
	fmt.Println(r)
}

func isSymmetric(root *TreeNode) bool {
	return check(root, root)

}

// 这里p、q是两棵树，同时遍历两棵树，所以最后需要两个check递归
func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil || q == nil) || (p.Val != q.Val) {
		return false
	}

	return check(p.Left, q.Right) && check(p.Right, q.Left)
}

// 迭代方法
func isSymmetric1(root *TreeNode) bool {
	u, v := root, root
	q := []*TreeNode{}
	q = append(q, u)
	q = append(q, v)
	for len(q) > 0 {
		u, v = q[0], q[1]
		q = q[2:]
		if u == nil && v == nil {
			continue
		}
		if u == nil || v == nil {
			return false
		}
		if u.Val != v.Val {
			return false
		}
		q = append(q, u.Left)
		q = append(q, v.Right)

		q = append(q, u.Right)
		q = append(q, v.Left)
	}
	return true
}

// isSymmetric2
//
//	@Description: 二刷
//	@param root
//	@return bool
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return symmetricCheck(root.Left, root.Right)
}

func symmetricCheck(p, q *TreeNode) bool {
	// 首先要对比两个根节点是否相同
	if p == nil && q == nil {
		return true
	} else if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	} else {
		if p.Val != q.Val {
			return false
		}
		// 根节点相同，再对比左右子树
		return symmetricCheck(p.Left, q.Right) && symmetricCheck(p.Right, q.Left)
	}

}
