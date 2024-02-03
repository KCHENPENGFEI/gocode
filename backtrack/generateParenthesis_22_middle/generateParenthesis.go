package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/12/30 下午2:44
 * @Desc:数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

 */

func main() {
	r := generateParenthesis(3)
	fmt.Println(r)
}

func generateParenthesis(n int) []string {
	var (
		res   []string
		track []byte
	)
	dfs(&res, n, track, n, n)
	return res
}

func dfs(result *[]string, n int, track []byte, left, right int) {
	if len(track) == 2*n {
		*result = append(*result, string(track))
		return
	}
	if left > right {
		// 说明左括号不够直接剪枝
		return
	}
	if left > 0 {
		track = append(track, '(')
		dfs(result, n, track, left-1, right)
		track = track[:len(track)-1]
	}
	if right > 0 {
		track = append(track, ')')
		dfs(result, n, track, left, right-1)
		track = track[:len(track)-1]
	}
}
