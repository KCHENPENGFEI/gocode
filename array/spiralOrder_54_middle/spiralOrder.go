package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/12/10 下午4:02
 * @Desc:
 */

func main() {
	m := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	//m := [][]int{{1}, {2}, {3}}
	r := spiralOrder(m)
	fmt.Println(r)
}

type dir int

const (
	right dir = iota + 1
	down
	left
	up
)

type point struct {
	i int
	j int
}

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	visited := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}

	var result []int
	curPoint := point{i: 0, j: -1}
	curDir := right
	for {
		if needChangeDir(visited, curPoint, curDir, m, n) {
			curDir = changeDir(curDir)
			curPoint = nextPoint(curPoint, curDir)
			if curPoint.i < 0 || curPoint.j < 0 || curPoint.i >= m || curPoint.j >= n || visited[curPoint.i][curPoint.j] == 1 {
				break
			}
		} else {
			curPoint = nextPoint(curPoint, curDir)
		}
		result = append(result, matrix[curPoint.i][curPoint.j])
		visited[curPoint.i][curPoint.j] = 1
	}
	return result
}

func changeDir(curDir dir) dir {
	return curDir%4 + 1
}

func needChangeDir(visited [][]int, curPoint point, curDir dir, m, n int) bool {
	switch curDir {
	case right:
		nextPoint := point{i: curPoint.i, j: curPoint.j + 1}
		return nextPoint.j >= n || visited[nextPoint.i][nextPoint.j] == 1
	case down:
		nextPoint := point{i: curPoint.i + 1, j: curPoint.j}
		return nextPoint.i >= m || visited[nextPoint.i][nextPoint.j] == 1
	case left:
		nextPoint := point{i: curPoint.i, j: curPoint.j - 1}
		return nextPoint.j < 0 || visited[nextPoint.i][nextPoint.j] == 1
	case up:
		nextPoint := point{i: curPoint.i - 1, j: curPoint.j}
		return nextPoint.i < 0 || visited[nextPoint.i][nextPoint.j] == 1
	}
	// 不可达
	return true
}

func nextPoint(curPoint point, curDir dir) point {
	var next point
	switch curDir {
	case right:
		next = point{i: curPoint.i, j: curPoint.j + 1}
	case down:
		next = point{i: curPoint.i + 1, j: curPoint.j}
	case left:
		next = point{i: curPoint.i, j: curPoint.j - 1}
	case up:
		next = point{i: curPoint.i - 1, j: curPoint.j}
	}
	return next
}
