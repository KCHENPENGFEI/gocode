package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/12/31 下午2:35
 * @Desc:
 */

func main() {
	n := 2
	p := [][]int{{1, 0}}
	r := canFinish(n, p)
	fmt.Println(r)
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 入度表
	var inDegrees = make([]int, numCourses)
	// 构建一个相邻边数组
	var adj = make([][]int, numCourses)
	// 生成入度表和边
	for _, p := range prerequisites {
		inDegrees[p[0]]++
		adj[p[1]] = append(adj[p[1]], p[0])
	}
	var queue []int
	for i, d := range inDegrees {
		if d == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		// 先学course这门课
		course := queue[0]
		queue = queue[1:]
		for _, canStudy := range adj[course] {
			// canStudy入度减1
			inDegrees[canStudy]--
			if inDegrees[canStudy] == 0 {
				queue = append(queue, canStudy)
			}
		}
		// 学完course这门课
		numCourses--
	}
	return numCourses == 0
}
