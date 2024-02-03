package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/12/31 下午3:27
 * @Desc:
 */

func main() {
	n := 6
	e := [][]int{{0, 1}, {0, 2}, {0, 3}, {3, 4}, {4, 5}}
	//e := [][]int{{1, 0}, {1, 2}, {1, 3}}
	r := findMinHeightTrees(n, e)
	fmt.Println(r)
}

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	// 生成adj信息
	var adj = make([][]int, n)
	var degree = make([]int, n)
	var res []int
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
		degree[edge[0]]++
		degree[edge[1]]++
	}
	var queue []int
	for i, d := range degree {
		if d == 1 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		res = []int{}
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			// 度为1的节点要写到结果集
			res = append(res, cur)
			for _, next := range adj[cur] {
				degree[next]--
				if degree[next] == 1 {
					queue = append(queue, next)
				}
			}
		}
	}
	return res
}

func getHeight(root int, adj [][]int) int {
	queue := []int{root}
	visited := map[int]struct{}{}
	height := 1
	var nextQueue []int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if len(queue) == 0 {
			height++
		}
		nextNodes := adj[node]
		visited[node] = struct{}{}
		for _, next := range nextNodes {
			if _, ok := visited[next]; ok {
				continue
			}
			nextQueue = append(nextQueue, next)
		}
		if len(queue) == 0 {
			queue = nextQueue
		}

	}
	return height
}
