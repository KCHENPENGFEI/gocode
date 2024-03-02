package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/12/10 下午4:02
 * @Desc: 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
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

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	var result []int
	visited := make(map[[2]int]struct{}, m*n)
	curPoint := [2]int{0, 0}
	curDir := right
	for len(visited) < m*n {
		result = append(result, matrix[curPoint[0]][curPoint[1]])
		visited[curPoint] = struct{}{}
		// 使用_curlPoint临时保存下一个点
		_curPoint := nextPoint(curDir, curPoint)
		if _, ok := visited[_curPoint]; ok || !validPoint(_curPoint, m, n) {
			// 判断是否要修改方向
			curDir = changeDir(curDir)
			curPoint = nextPoint(curDir, curPoint)
		} else {
			curPoint = _curPoint
		}
	}
	return result
}

func changeDir(curDir dir) dir {
	return curDir%4 + 1
}

func validPoint(curPoint [2]int, m, n int) bool {
	return curPoint[0] >= 0 && curPoint[1] >= 0 && curPoint[0] < m && curPoint[1] < n
}

func nextPoint(curDir dir, curPoint [2]int) [2]int {
	var next [2]int
	switch curDir {
	case right:
		return [2]int{curPoint[0], curPoint[1] + 1}
	case down:
		return [2]int{curPoint[0] + 1, curPoint[1]}
	case left:
		return [2]int{curPoint[0], curPoint[1] - 1}
	case up:
		return [2]int{curPoint[0] - 1, curPoint[1]}
	default:
		return next
	}
}
