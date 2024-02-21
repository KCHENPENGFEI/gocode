package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2024/2/17 下午1:49
 * @Desc: 239. 滑动窗口最大值
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回 滑动窗口中的最大值 。
*/

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	r := maxSlidingWindow(nums, k)
	fmt.Println(r)
}

func maxSlidingWindow1(nums []int, k int) []int {
	q := []int{}
	push := func(i int) {
		for len(q) > 0 && nums[i] > nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}

// maxSlidingWindow
//
//	@Description: 构造一个单调栈maxStack存储降序的nums的下标
//	@param nums
//	@param k
//	@return []int
func maxSlidingWindow(nums []int, k int) []int {
	var maxStack []int
	// 单调栈push函数
	push := func(i int) {
		for len(maxStack) > 0 && nums[i] > nums[maxStack[len(maxStack)-1]] {
			maxStack = maxStack[:len(maxStack)-1]
		}
		maxStack = append(maxStack, i)
	}
	for i := 0; i < k; i++ {
		// 初始化单调栈
		push(i)
	}
	n := len(nums)
	res := make([]int, n-k+1)
	res[0] = nums[maxStack[0]]
	for i := k; i < n; i++ {
		push(i)
		// 说明此时头部元素不在滑动窗口中了
		if maxStack[0] <= i-k {
			maxStack = maxStack[1:]
		}
		// 头部元素就是最大值的下标
		res[i-k+1] = nums[maxStack[0]]
	}
	return res
}
