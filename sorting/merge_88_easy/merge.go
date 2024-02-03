package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/8/5 下午1:07
 * @Desc:
 */

func main() {
	nums1 := []int{1, 2, 3, 4, 10, 0, 0, 0}
	nums2 := []int{7, 8, 9}
	merge(nums1, 5, nums2, 3)
	fmt.Println(nums1)
}

//
// merge
//  @Description: 从后往前遍历，可以减少数据的移动次数
//  @param nums1
//  @param m
//  @param nums2
//  @param n
//
func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	i := m - 1
	j := n - 1
	index := m + n - 1
	for i >= 0 || j >= 0 {
		if j < 0 {
			return
		}
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[index] = nums1[i]
			index--
			i--
		} else {
			nums1[index] = nums2[j]
			index--
			j--
		}
	}
}
