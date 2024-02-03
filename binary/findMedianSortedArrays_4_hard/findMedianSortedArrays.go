package main

import "sort"

/**
 * @Author: chenpengfei
 * @Date: 2023/12/17 下午4:08
 * @Desc:
 */

func main() {

}

func findMedianSortedArraysUseSort(nums1 []int, nums2 []int) float64 {
	bigSlice := make([]int, len(nums1)+len(nums2))
	copy(bigSlice[0:len(nums1)], nums1)
	copy(bigSlice[len(nums1):], nums2)
	sort.Ints(bigSlice)
	l := len(bigSlice)
	if l%2 != 0 {
		return float64(bigSlice[l/2])
	} else {
		return (float64(bigSlice[l/2]) + float64(bigSlice[l/2-1])) / 2
	}
}
