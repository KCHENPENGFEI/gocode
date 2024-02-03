package main

import (
	"fmt"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/12/31 下午1:06
 * @Desc:
 */

func main() {
	g := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	//g := [][]int{{1, 2}}
	r := orangesRotting(g)
	fmt.Println(r)
}

type point struct {
	i int
	j int
}

func orangesRotting(grid [][]int) int {
	list, blank := addFirstNode(grid)
	// 添加初始节点
	if len(list) == 0 {
		if blank {
			return 0
		} else {
			return -1
		}
	}
	days := 0
	for len(list) > 0 {
		var newList []point
		for _, p := range list {
			next := changeAround(grid, p)
			newList = append(newList, next...)
		}
		list = newList
		if len(list) > 0 {
			days++
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return days

}

func changeAround(grid [][]int, p point) []point {
	var r []point
	next := point{i: p.i - 1, j: p.j}
	if next.i >= 0 && next.j >= 0 && next.i < len(grid) && next.j < len(grid[0]) && grid[next.i][next.j] == 1 {
		grid[next.i][next.j] = 2
		r = append(r, next)
	}
	next = point{i: p.i + 1, j: p.j}
	if next.i >= 0 && next.j >= 0 && next.i < len(grid) && next.j < len(grid[0]) && grid[next.i][next.j] == 1 {
		grid[next.i][next.j] = 2
		r = append(r, next)
	}
	next = point{i: p.i, j: p.j - 1}
	if next.i >= 0 && next.j >= 0 && next.i < len(grid) && next.j < len(grid[0]) && grid[next.i][next.j] == 1 {
		grid[next.i][next.j] = 2
		r = append(r, next)
	}
	next = point{i: p.i, j: p.j + 1}
	if next.i >= 0 && next.j >= 0 && next.i < len(grid) && next.j < len(grid[0]) && grid[next.i][next.j] == 1 {
		grid[next.i][next.j] = 2
		r = append(r, next)
	}
	return r
}

func addFirstNode(grid [][]int) ([]point, bool) {
	var list []point
	blank := true
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				list = append(list, point{i: i, j: j})
				blank = false
			} else if grid[i][j] == 1 {
				// 非空
				blank = false
			}
		}
	}
	return list, blank
}

func maxTwo(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDir(a, b, c, d int) int {
	m := maxTwo(a, b)
	m = maxTwo(m, c)
	m = maxTwo(m, d)
	return m
}
