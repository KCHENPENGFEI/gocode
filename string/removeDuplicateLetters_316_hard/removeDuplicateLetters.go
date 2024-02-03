package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/8/19 下午4:20
 * @Desc:
 */

func main() {
	s := "cbacdcbc"
	fmt.Println(removeDuplicateLetters(s))
}

func removeDuplicateLetters(s string) string {
	// 首先统计字符个数
	cntMap := make(map[rune]int)
	for _, c := range s {
		if _, ok := cntMap[c]; !ok {
			cntMap[c] = 1
		} else {
			cntMap[c]++
		}
	}

	isInStack := make(map[rune]bool)
	var stack = make([]rune, 0)
	for _, c := range s {
		if in, ok := isInStack[c]; !ok || !in {
			// 该字符不在stack中
			for len(stack) > 0 && c < stack[len(stack)-1] {
				top := stack[len(stack)-1]
				if cntMap[top] == 0 {
					break
				}
				// 删除栈顶元素
				stack = stack[:len(stack)-1]
				isInStack[top] = false
			}
			// 元素入栈
			stack = append(stack, c)
			isInStack[c] = true
		}
		cntMap[c]--
	}

	return string(stack)
}
