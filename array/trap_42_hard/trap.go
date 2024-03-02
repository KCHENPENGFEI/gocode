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

// trap2
//
//	@Description: 二刷，掌握单调栈
//	@param height
//	@return int
func trap2(height []int) int {
	var stack []int
	result := 0
	for i := 0; i < len(height); i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			// 调整这个递减的单调栈
			// 先把栈顶元素pop出来
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 栈空了说明没法再生成积水区域了
			if len(stack) == 0 {
				break
			}
			// 获取当前栈顶元素
			currentTop := stack[len(stack)-1]
			// 计算区域面积公式
			result += (minTwo(height[i], height[currentTop]) - height[top]) * (i - currentTop - 1)
		}
		stack = append(stack, i)
	}
	return result
}

func minTwo(a, b int) int {
	if a < b {
		return a
	}
	return b
}
