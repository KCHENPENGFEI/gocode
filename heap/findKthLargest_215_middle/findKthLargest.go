package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2024/1/27 下午2:25
 * @Desc:
 */

func main() {
	n := []int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}
	k := 20
	r := findKthLargest(n, k)
	fmt.Println(r)
}

func findKthLargest(nums []int, k int) int {
	// 先构造一个大顶堆
	buildMaxHeap(nums)
	for i := 0; i < k-1; i++ {
		// 删除堆顶元素，将最后一个元素放在堆顶后重建堆
		nums[0], nums[len(nums)-1] = nums[len(nums)-1], nums[0]
		nums = nums[:len(nums)-1]
		adjust(nums, 0)
	}
	return nums[0]
}

func buildMaxHeap(nums []int) {
	for i := len(nums) / 2; i >= 0; i-- {
		adjust(nums, i)
	}
}

func adjust(num []int, parent int) {
	leftChild, rightChild := 2*parent+1, 2*parent+2
	idx := parent
	if leftChild < len(num) && num[leftChild] > num[idx] {
		// 调整位置
		idx = leftChild
	}
	if rightChild < len(num) && num[rightChild] > num[idx] {
		idx = rightChild
	}
	// 如果位置有调整
	if idx != parent {
		num[parent], num[idx] = num[idx], num[parent]
		adjust(num, idx)
	}
}

// findKthLargest2
//
//	@Description: 二刷
//	@param nums
//	@param k
//	@return int
func findKthLargest2(nums []int, k int) int {
	// 构建一个大根堆
	buildHeap(nums)
	for i := 0; i < k-1; i++ {
		// 交换顶部和尾部的数据，然后删除尾部数据
		nums[0], nums[len(nums)-1] = nums[len(nums)-1], nums[0]
		nums = nums[0 : len(nums)-1]
		adjust2(nums, 0)
	}
	return nums[0]
}

func buildHeap(nums []int) {
	// 一定要从底部往上开始调整
	// 大顶堆构建过程就是先从底部往上构建，然后针对每个节点都需要从顶部往下调整
	for i := len(nums) / 2; i >= 0; i-- {
		adjust2(nums, i)
	}
}

func adjust2(nums []int, rootIndex int) {
	// 将根节点下标为rootIndex的二叉树调整为大根堆
	left, right := 2*rootIndex+1, 2*(rootIndex+1)
	idx := rootIndex
	if left < len(nums) && nums[left] > nums[rootIndex] {
		idx = left
	}
	if right < len(nums) && nums[right] > nums[idx] {
		idx = right
	}
	if idx != rootIndex {
		nums[idx], nums[rootIndex] = nums[rootIndex], nums[idx]
		adjust2(nums, idx)
	}
}
