package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/12/16 下午3:54
 * @Desc:
 */

func main() {
	//m := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	//t := 1
	m := [][]int{{-5}}
	t := -5
	fmt.Println(searchMatrix2(m, t))
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	maxLine := findMaxLine(matrix, target, m)
	if maxLine == 0 && matrix[0][0] > target {
		return false
	}
	for i := 0; i <= maxLine; i++ {
		if findInLine(matrix[i], target, n) {
			return true
		}
	}
	return false
}

func findMaxLine(matrix [][]int, target int, m int) int {
	l := 0
	r := m - 1
	for l < r {
		mid := l + (r-l+1)/2
		num := matrix[mid][0]
		if target == num {
			return mid
		}
		if target > num {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}

func findInLine(line []int, target int, n int) bool {
	l, r := 0, n // mid偏左，因此r需要取到n
	for l < r {
		mid := l + (r-l)/2
		num := line[mid]
		if num == target {
			return true
		}
		if num < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return false
}

func searchMatrix2(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	if matrix[0][0] > target || matrix[m-1][n-1] < target {
		return false
	}
	up := findUpRow(matrix, target)
	bottom := findBottomRow(matrix, target)
	if up < 0 || up >= m || bottom < 0 || bottom >= m {
		return false
	}
	for i := up; i <= bottom; i++ {
		if findInLine(matrix[i], target, n) {
			return true
		}
	}
	return false
}

func findUpRow(matrix [][]int, target int) int {
	m, n := len(matrix), len(matrix[0])
	// mid偏左，因此r取m不会越界
	l, r := 0, m
	for l < r {
		mid := l + (r-l)/2
		if matrix[mid][n-1] < target {
			// mid偏左，因此变化l
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func findBottomRow(matrix [][]int, target int) int {
	m, _ := len(matrix), len(matrix[0])
	// mid偏右，因此l取-1不会越界，但是r不能m，只能取m-1
	l, r := -1, m-1
	for l < r {
		mid := l + (r-l+1)/2
		if matrix[mid][0] > target {
			// mid偏右，因此变化r
			r = mid - 1
		} else {
			l = mid
		}
	}
	return r
}
