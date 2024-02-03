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
