package main

import (
	"fmt"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/9/3 下午4:37
 * @Desc:
 */

func main() {
	height := []int{4, 2, 0, 3, 2, 5}
	fmt.Println(trap(height))
}

func trap(height []int) int {
	sum := 0
	for i := 1; i < len(height)-1; i++ {
		leftMaxHeight, rightMaxHeight := 0, 0
		for j := 0; j < i; j++ {
			if height[j] > leftMaxHeight {
				leftMaxHeight = height[j]
			}
		}
		for j := i + 1; j < len(height); j++ {
			if height[j] > rightMaxHeight {
				rightMaxHeight = height[j]
			}
		}
		m := min(leftMaxHeight, rightMaxHeight)
		if height[i] < m {
			sum += m - height[i]
		}
	}
	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
