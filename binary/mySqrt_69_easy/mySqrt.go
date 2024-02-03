package main

/**
 * @Author: chenpengfei
 * @Date: 2023/12/17 下午2:15
 * @Desc:给你一个非负整数 x ，计算并返回 x 的 算术平方根 。

由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
*/

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	// 二分范围可以使用x/2缩小
	maxNum := x / 2
	if maxNum < 1 {
		maxNum = 1
	}
	l, r := 1, maxNum
	for l < r {
		mid := l + (r-l+1)/2
		if mid*mid > x {
			r = mid - 1
		} else {
			l = mid
		}
	}
	return l
}
