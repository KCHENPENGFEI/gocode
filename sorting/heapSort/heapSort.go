package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2024/2/28 上午1:07
 * @Desc:
 */

func main() {
	nums := []int{4, 5, 2, 1, 8, 6, 9, 12, 56, 34, 14, 25}
	heapSort(nums)
	fmt.Println(nums)
}

func heapSort(nums []int) {
	// 构造一个大顶堆，然后每次循环都把顶部元素和尾部元素对调，
	// 继续重新构造nums[:len(nums)-1]的大顶堆
	n := len(nums)
	for i := 0; i < n; i++ {
		buildHeap(nums[:n-i])
		nums[0], nums[n-i-1] = nums[n-i-1], nums[0]
	}
}

func buildHeap(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		adjustLoop(nums, i)
	}
}

func adjustLoop(nums []int, rootIndex int) {
	n := len(nums)
	son := 2*rootIndex + 1
	for son < n {
		if son+1 < n && nums[son+1] > nums[son] {
			son++
		}
		if nums[son] > nums[rootIndex] {
			nums[son], nums[rootIndex] = nums[rootIndex], nums[son]
			son = 2*son + 1
		} else {
			break
		}
	}
}
