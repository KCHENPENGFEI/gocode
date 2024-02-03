package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2024/1/13 下午4:53
 * @Desc:
 */

func main() {
	h := []int{5, 4, 1, 2}
	r := largestRectangleArea(h)
	fmt.Println(r)
}

// 构造一个单调非递减栈
func largestRectangleArea(heights []int) int {
	ans := 0
	var stack []int
	for i := 0; i < len(heights); i++ {
		height := heights[i]
		for len(stack) > 0 && height < heights[stack[len(stack)-1]] {
			// 先pop出来
			topIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			var maxWidth int
			if len(stack) == 0 {
				maxWidth = i
			} else {
				maxWidth = i - stack[len(stack)-1] - 1
			}
			ans = maxTwo(ans, maxWidth*heights[topIndex])
		}
		stack = append(stack, i)
	}
	// 遍历stack直到stack为空
	for i := 0; i < len(stack); i++ {
		dis := stack[len(stack)-1] - stack[i] + 1
		ans = maxTwo(ans, dis*heights[stack[i]])
	}

	return ans
}

func maxTwo(a, b int) int {
	if a > b {
		return a
	}
	return b
}
