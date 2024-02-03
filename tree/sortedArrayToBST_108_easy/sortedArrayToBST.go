package main

import (
	. "golang/tree"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/10/29 下午1:23
 * @Desc:
 */

func main() {

}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) >> 1
	return &TreeNode{Val: nums[mid], Left: sortedArrayToBST(nums[:mid]), Right: sortedArrayToBST(nums[mid+1:])}
}
