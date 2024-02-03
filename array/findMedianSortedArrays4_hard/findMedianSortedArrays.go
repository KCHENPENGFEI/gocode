package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2022/9/11 2:57 PM
 * @Desc: 给定两个大小分别为 m 和 n 的正序（从小到大）数组nums1 和nums2。请你找出并返回这两个正序数组的中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。

*/

func main() {
	nums1 := []int{1}
	nums2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// num1和num2不会同时为空
	len1 := len(nums1)
	len2 := len(nums2)
	if (len1+len2)%2 == 1 {
		return float64(findTopKNum(nums1, nums2, (len1+len2)/2+1))
	} else {
		return (float64(findTopKNum(nums1, nums2, (len1+len2)/2)) +
			float64(findTopKNum(nums1, nums2, (len1+len2)/2+1))) / 2
	}
}

//
// findTopKNum
//  @Description: 将找中位数问题转换成找topN问题
// 如果nums1长度为m，nums2长度为n，则中位数为topK小的数，K = (m+n)/2，切向上取整。那么则可以比较nums1的第K/2和nums2
// 的第K/2的数，删除较小数所在数组的钱K/2个数字，然后重新寻找两个数组中第top(K-K/2)的数即可。
// 算法需要注意边界情况
//  @param nums1
//  @param nums2
//  @param k
//  @return int
//
func findTopKNum(nums1, nums2 []int, k int) int {
	len1 := len(nums1)
	len2 := len(nums2)
	i, j := 0, 0
	for {
		iBak := i
		jBak := j
		// 特殊情况：nums1或者nums2为空
		if i >= len1 {
			return nums2[j+k-1]
		}
		if j >= len2 {
			return nums1[i+k-1]
		}
		// 循环出口：k == 1
		if k == 1 {
			return min(nums1[i], nums2[j])
		}

		halfK := k / 2 // 此时halfK必然>=1
		var newI, newJ int
		// 边界情况，NewI取值到数字最后一个数
		if i+halfK-1 >= len1 {
			newI = len1 - 1
		} else {
			newI = i + halfK - 1
		}

		// 边界情况，NewJ取值到数字最后一个数
		if j+halfK-1 >= len2 {
			newJ = len2 - 1
		} else {
			newJ = j + halfK - 1
		}

		// moves表示实际上删除的数字数量
		moves := 0
		if nums1[newI] < nums2[newJ] {
			i = newI + 1
			moves = i - iBak
		} else {
			j = newJ + 1
			moves = j - jBak
		}
		k = k - moves
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
