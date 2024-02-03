package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/12/16 下午3:54
 * @Desc:
 */

func main() {
	m := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	t := 24
	fmt.Println(searchMatrix(m, t))
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
	l := 0
	r := n - 1
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
