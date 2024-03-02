package main

import (
	"fmt"
	. "golang/tree"
)

/**
 * @Author: chenpengfei
 * @Date: 2024/2/3 下午2:44
 * @Desc: 给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。

路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
*/

func main() {
	root := &TreeNode{Val: 10, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: -2}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 1}}},
		Right: &TreeNode{Val: -3, Right: &TreeNode{Val: 11}}}
	t := 8
	r := pathSum2(root, t)
	fmt.Println(r)
}

// pathSum
//
//	@Description: 一刷做法，维护一个从根节点到当前节点的路径切片，每次递归时候从切片中找和为targetSum的数据
//	@param root
//	@param targetSum
//	@return int
func pathSum(root *TreeNode, targetSum int) int {
	count := 0
	helper(root, nil, targetSum, &count)
	return count
}

func helper(root *TreeNode, path []int, targetSum int, count *int) {
	if root == nil {
		return
	}
	sum := root.Val
	// 避免path为空进入不了循环的情况
	if sum == targetSum {
		*count++
	}
	// 从当前节点向根节点遍历
	for i := len(path) - 1; i >= 0; i-- {
		sum += path[i]
		if sum == targetSum {
			*count++
		}
	}
	// 因为path不是指针数组，因此递归函数中只是拷贝，不会进行回退操作
	path = append(path, root.Val)
	helper(root.Left, path, targetSum, count)
	helper(root.Right, path, targetSum, count)
}

// pathSum2
//
//	@Description: 二刷，结合560题求解连续子数组方法，使用一个map来存储连续和和数量的映射
//
// 要注意的是每次回退的时候需要回退sum和sumCount的值
//
//	@param root
//	@param targetSum
//	@return int
func pathSum2(root *TreeNode, targetSum int) int {
	sumCount := make(map[int]int)
	sumCount[0] = 1
	var ans, sum int
	helper2(root, targetSum, sumCount, &ans, &sum)
	return ans
}

func helper2(root *TreeNode, targetSum int, sumCount map[int]int, ans *int, sum *int) {
	if root == nil {
		return
	}
	// 当前节点+sum
	*sum += root.Val
	if cnt, ok := sumCount[*sum-targetSum]; ok {
		*ans += cnt
	}
	sumCount[*sum]++
	// 处理左子树
	if root.Left != nil {
		helper2(root.Left, targetSum, sumCount, ans, sum)
		// 注意要处理回退的情况
		sumCount[*sum]--
		*sum -= root.Left.Val
	}
	if root.Right != nil {
		helper2(root.Right, targetSum, sumCount, ans, sum)
		sumCount[*sum]--
		*sum -= root.Right.Val
	}
}
