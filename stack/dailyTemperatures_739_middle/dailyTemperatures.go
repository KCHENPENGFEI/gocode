package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2024/1/13 下午4:14
 * @Desc: 给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，
其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。
*/

func main() {
	t := []int{30, 40, 50, 60}
	r := dailyTemperatures(t)
	fmt.Println(r)
}

func dailyTemperatures(temperatures []int) []int {
	// 单调栈，栈中元素存储下标index，并且温度是单调下降的
	var stack []int
	var result = make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		// 维护单调栈同时写结果
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			// 出现上升的温度则需要判断栈顶元素是否要出栈
			result[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		// 入栈
		stack = append(stack, i)
	}
	for _, index := range stack {
		result[index] = 0
	}
	return result
}
